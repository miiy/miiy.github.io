#!/usr/bin/env python3
"""Check every link in _data/links.yml.

Reports, per entry:
  OK       2xx, final URL unchanged
  MOVED    redirects to another URL (domain/path changed)
  DEAD     4xx / 5xx / timeout / DNS failure
  TITLE    page title differs from the name stored in links.yml

Usage:
  python3 scripts/check_links.py                # check all, print report
  python3 scripts/check_links.py --json r.json  # also dump raw results
  python3 scripts/check_links.py --apply        # rewrite names that drifted
  python3 scripts/check_links.py --only 50      # check the first 50 entries
  HTTPS_PROXY=http://127.0.0.1:7890 python3 scripts/check_links.py

Exit code: 0 when nothing is dead, 1 otherwise (handy for CI / cron).
"""
from __future__ import annotations

import argparse
import concurrent.futures
import html
import json
import re
import socket
import ssl
import sys
import urllib.error
import urllib.request
from pathlib import Path

import yaml

ROOT = Path(__file__).resolve().parents[2]   # _data/scripts/ -> repo root
LINKS = ROOT / "_data" / "links.yml"
UA = "Mozilla/5.0 (compatible; link-checker/1.0)"
TIMEOUT = 10
TITLE_RE = re.compile(r"<title[^>]*>(.*?)</title>", re.I | re.S)
JUNK_TITLE = re.compile(
    r"^(just a moment|attention required|access denied|403 forbidden|404|page not found|not found|error)\b",
    re.I,
)
# suffix that sites add to their own <title>, stripped before comparing
SITE_SUFFIX = re.compile(r"\s+[|–—-]\s+[^|–—-]{1,40}$")


def fetch(url: str) -> dict:
    """Return {status, final_url, title, error} for one URL."""
    req = urllib.request.Request(url, headers={"User-Agent": UA, "Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8"})
    ctx = ssl.create_default_context()
    ctx.check_hostname = False
    ctx.verify_mode = ssl.CERT_NONE
    try:
        with urllib.request.urlopen(req, timeout=TIMEOUT, context=ctx) as resp:
            body = resp.read(200_000)
            charset = resp.headers.get_content_charset() or "utf-8"
            try:
                text = body.decode(charset, errors="ignore")
            except LookupError:
                text = body.decode("utf-8", errors="ignore")
            m = TITLE_RE.search(text)
            title = re.sub(r"\s+", " ", html.unescape(m.group(1))).strip() if m else ""
            return {"status": resp.status, "final_url": resp.geturl(), "title": title, "error": ""}
    except urllib.error.HTTPError as e:
        return {"status": e.code, "final_url": url, "title": "", "error": f"HTTP {e.code}"}
    except urllib.error.URLError as e:
        return {"status": 0, "final_url": url, "title": "", "error": str(e.reason)}
    except (socket.timeout, TimeoutError):
        return {"status": 0, "final_url": url, "title": "", "error": "timeout"}
    except Exception as e:  # noqa: BLE001 - report anything else as a failure
        return {"status": 0, "final_url": url, "title": "", "error": type(e).__name__}


def clean_title(title: str) -> str:
    t = SITE_SUFFIX.sub("", title).strip()
    return t


def same_title(a: str, b: str) -> bool:
    norm = lambda s: re.sub(r"[\s\-_·|:：/]+", "", s).lower()
    return norm(a) == norm(b) or norm(a) in norm(b) or norm(b) in norm(a)


def is_github(url: str) -> bool:
    return "github.com" in url


def main() -> int:
    ap = argparse.ArgumentParser()
    ap.add_argument("--json", metavar="FILE", help="dump raw results to FILE")
    ap.add_argument("--apply", action="store_true", help="rewrite names where the page title clearly changed")
    ap.add_argument("--only", type=int, help="only check the first N entries")
    ap.add_argument("--jobs", type=int, default=8)
    args = ap.parse_args()

    doc = yaml.safe_load(LINKS.read_text(encoding="utf-8"))
    entries = [(g["title"], i) for g in doc["groups"] for i in g["items"]]
    if args.only:
        entries = entries[: args.only]

    results = []
    with concurrent.futures.ThreadPoolExecutor(max_workers=args.jobs) as pool:
        futures = {pool.submit(fetch, item["url"]): (group, item) for group, item in entries}
        for fut in concurrent.futures.as_completed(futures):
            group, item = futures[fut]
            results.append({"group": group, "name": item["name"], "url": item["url"], **fut.result()})

    results.sort(key=lambda r: (r["group"], r["name"]))
    dead, moved, retitled, ok = [], [], [], []
    for r in results:
        if not r["status"] or r["status"] >= 400:
            dead.append(r)
        elif r["final_url"].rstrip("/") not in (r["url"].rstrip("/"), r["url"].rstrip("/") + "/"):
            moved.append(r)
        elif r["title"] and not JUNK_TITLE.search(r["title"]) and not is_github(r["url"]):
            new = clean_title(r["title"])
            if new and not same_title(r["name"], new):
                retitled.append({**r, "clean_title": new})
            else:
                ok.append(r)
        else:
            ok.append(r)

    print(f"\n共 {len(results)} 条：正常 {len(ok)} · 标题变化 {len(retitled)} · 跳转 {len(moved)} · 失效 {len(dead)}\n")

    if dead:
        print("== 失效 ==")
        for r in dead:
            print(f"  [{r['group']}] {r['name'][:28]:30s} {r['url'][:64]:66s} {r['error'] or r['status']}")
        print()
    if moved:
        print("== 跳转到新地址 ==")
        for r in moved:
            print(f"  [{r['group']}] {r['name'][:24]:26s} {r['url']}\n      → {r['final_url']}")
        print()
    if retitled:
        print("== 标题与页面不一致 ==")
        for r in retitled:
            print(f"  [{r['group']}] 现名: {r['name']}\n      页面: {r['clean_title'][:90]}\n      {r['url']}")
        print()

    if args.json:
        Path(args.json).write_text(json.dumps(results, ensure_ascii=False, indent=2), encoding="utf-8")
        print(f"原始结果已写入 {args.json}")

    if args.apply and retitled:
        by_url = {r["url"]: r["clean_title"] for r in retitled}
        changed = 0
        for g in doc["groups"]:
            for i in g["items"]:
                if i["url"] in by_url and len(by_url[i["url"]]) <= 40:
                    i["name"] = by_url[i["url"]]
                    changed += 1
        LINKS.write_text(
            yaml.dump(doc, allow_unicode=True, sort_keys=False, width=9999, default_flow_style=False),
            encoding="utf-8",
        )
        print(f"已更新 {changed} 个名称（超过 40 字符的自动跳过）")

    return 1 if dead else 0


if __name__ == "__main__":
    sys.exit(main())

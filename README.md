# Blog

## Build

Gems are installed into the image (`/usr/local/bundle`) at build time, so `docker run` never has to install anything.

```bash
docker build -t miiy/jekyll .
```

## Update bundle

Dependencies are baked into the image, so updating is two steps: rewrite `Gemfile.lock`, then rebuild the image.

```bash
docker run --rm -it -v "$PWD":/app miiy/jekyll sh -c 'bundle update --all'
docker build -t miiy/jekyll .
```

## run

```bash
docker run --rm -it -v "$PWD":/app -p 4000:4000 miiy/jekyll
```

## sh

```bash
docker run --rm -it -v "$PWD":/app -p 4000:4000 miiy/jekyll sh
```

## Jekyll

Create a new site

```bash
jekyll new ./
jekyll new blog
jekyll serve -H 0.0.0.0 -P 4000
jekyll build
```

Plugins on GitHub Pages

https://pages.github.com/versions

## Bundle - Ruby Dependency Management

```bash
bundle install
bundle show minima
```

## Rouge

Generate Rouge highlight syntax

```bash
rougify help style
rougify style github > github.css
rougify style github > github.css
rougify style colorful > colorful.css
```

## Tech stack

- [Jekyll](https://jekyllrb.com/) — static site generator
  - [jekyll/jekyll](https://github.com/jekyll/jekyll)
  - [jekyll-watch](https://github.com/jekyll/jekyll-watch) — rebuild on file change
  - [jekyll-sass-converter](https://github.com/jekyll/jekyll-sass-converter) — SCSS compilation
  - [Jekyll Themes](http://jekyllthemes.org/)
- [minima](https://github.com/jekyll/minima) — theme (pinned by commit in `Gemfile`)
- [Liquid](https://shopify.github.io/liquid/) — template language
- [Rouge](https://github.com/rouge-ruby/rouge) — syntax highlighting
- [Simple-Jekyll-Search](https://github.com/christian-fei/Simple-Jekyll-Search) — client-side search
- [GitHub Pages](https://docs.github.com/en/free-pro-team@latest/github/working-with-github-pages) — deployment

## Check links

`_data/links.yml` powers the `/links/` page. To find dead links, moved URLs and
titles that drifted from the live pages:

```bash
python3 _data/scripts/check_links.py                 # report only
python3 _data/scripts/check_links.py --json out.json # dump raw results
python3 _data/scripts/check_links.py --apply         # rewrite names from page titles
HTTPS_PROXY=http://127.0.0.1:7890 python3 _data/scripts/check_links.py
```

Exits with status 1 when a link is dead, so it can run in CI or a cron job.

# Blog

## Build

```bash
docker build -t test/jekyll .
```

## run

```bash
docker run --rm -it -v "$PWD":/app -p 4000:4000 test/jekyll
```

## sh

```bash
docker run --rm -it -v "$PWD":/app -p 4000:4000 test/jekyll sh
```

## Jekyll

```bash
jekyll new ./
jekyll new blog
jekyll serve -H 0.0.0.0 -P 4000
jekyll build
```

## Bundle - Ruby Dependency Management

```bash
bundle install
bundle show minima
```

## rougify

```bash
rougify help style
rougify style monokai > syntax.css
```

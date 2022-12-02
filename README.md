# Blog

## Build

```bash
docker build -t miiy/jekyll _deploy
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

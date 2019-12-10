## Install

docker run --rm -volumn=./jekyll:/srv/jekyll -it jekyll bash

## Commands

```bash
# start
docker-compose up

# build
docker-compose run --rm jekyll jekyll build

# bash
docker-compose run --rm jekyll bash

# jekyll
jekyll new ./
jekyll new blog
jekyll serve -H 0.0.0.0 -P 4000
jekyll build

bundle install
bundle show minima

# rougify
rougify help style
rougify style monokai > syntax.css
```

Working with GitHub Pages: https://help.github.com/en/github/working-with-github-pages

Jekyll: https://jekyllrb.com

Jekyll: https://github.com/jekyll/jekyll

minima: https://github.com/jekyll/minima

Liquid: https://shopify.github.io/liquid

Jekyll Sass Converter: https://github.com/jekyll/jekyll-sass-converter

Rouge: https://github.com/rouge-ruby/rouge
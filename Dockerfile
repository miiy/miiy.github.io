FROM debian:latest

RUN sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list.d/debian.sources \
    && cat /etc/apt/sources.list.d/debian.sources \
    && echo "\n" \
    && cat /etc/apt/sources.list.d/debian.sources

RUN apt-get update \
    && apt-get install -y git vim \
    && apt-get install -y ruby ruby-dev gcc g++ make pkg-config libssl-dev \
    && rm -rf /var/lib/apt/lists/* \
    && apt-get clean

RUN gem sources --add https://mirrors.ustc.edu.cn/rubygems/ --remove https://rubygems.org/ \
    && gem sources -u \
    && gem install bundler jekyll \
    && bundle config set --global mirror.https://rubygems.org https://mirrors.ustc.edu.cn/rubygems/ \
    && bundle config set --global path /usr/local/bundle

WORKDIR /app

# Gems live in the image, not in /app/vendor: /app is replaced by the
# `-v "$PWD":/app` bind mount, so anything installed there would be hidden.
ENV BUNDLE_PATH=/usr/local/bundle

# Install dependencies at build time; bump Gemfile.lock then rebuild the image.
COPY Gemfile Gemfile.lock /app/
RUN bundle install --jobs 4 --retry 3 \
    && rm -f /usr/local/bundle/ruby/*/cache/*.gem

# Copied after the bundle layer on purpose: editing the entrypoint should not
# invalidate the expensive gem installation above.
COPY ./entrypoint.sh /bin
RUN chmod +x /bin/entrypoint.sh

EXPOSE 4000

ENTRYPOINT ["/bin/entrypoint.sh"]

CMD ["bundle", "exec", "jekyll", "serve", "-H", "0.0.0.0", "-P", "4000"]

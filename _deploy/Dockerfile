FROM debian:bookworm

RUN sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list.d/debian.sources \
    && cat /etc/apt/sources.list.d/debian.sources \
    && echo "\n" \
    && cat /etc/apt/sources.list.d/debian.sources

RUN apt-get update \
    && apt-get install -y git vim \
    && apt-get install -y ruby ruby-dev gcc g++ make \
    && rm -rf /var/lib/apt/lists/* \
    && apt-get clean

RUN gem sources --add https://gems.ruby-china.com/ --remove https://rubygems.org/ \
    && gem sources -u \
    && gem install bundler jekyll \
    && bundle config mirror.https://rubygems.org https://gems.ruby-china.com \
    && bundle config set path vendor

COPY ./entrypoint.sh /bin
RUN chmod +x /bin/entrypoint.sh

WORKDIR /app

EXPOSE 4000

ENTRYPOINT ["/bin/entrypoint.sh"]

CMD ["bundle", "exec", "jekyll", "serve", "-H", "0.0.0.0", "-P", "4000"]
FROM alpine:latest

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add bash git nodejs \
    ruby ruby-dev ruby-rdoc ruby-bigdecimal ruby-webrick ruby-etc build-base && \
    gem sources --add https://gems.ruby-china.com/ --remove https://rubygems.org/ && \
    gem sources -u && \
    gem install bundle jekyll

WORKDIR /app

COPY ./Gemfile /app

RUN bundle install

EXPOSE 4000

CMD [ "jekyll", "serve", "-H", "0.0.0.0", "-P", "4000"]
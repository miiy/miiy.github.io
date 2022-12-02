FROM alpine:3.17

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk update && \
    apk add \
    ruby ruby-dev ruby-rdoc ruby-bigdecimal ruby-webrick ruby-etc build-base && \
    gem sources --add https://gems.ruby-china.com/ --remove https://rubygems.org/ && \
    gem sources -u && \
    gem install bundle jekyll && \
    bundle config mirror.https://rubygems.org https://gems.ruby-china.com

COPY ./entrypoint.sh /bin
RUN chmod +x /bin/entrypoint.sh

WORKDIR /app

EXPOSE 4000

ENTRYPOINT ["/bin/entrypoint.sh"]

CMD ["jekyll", "serve", "-H", "0.0.0.0", "-P", "4000"]
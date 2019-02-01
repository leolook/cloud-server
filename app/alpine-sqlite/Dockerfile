FROM alpine

MAINTAINER hwt
LABEL cloud-server 1.0.0

RUN apk --update upgrade
RUN apk add sqlite
# See http://stackoverflow.com/questions/34729748/installed-go-binary-not-found-in-path-on-alpine-linux-docker
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

# removing apk cache
RUN rm -rf /var/cache/apk/*
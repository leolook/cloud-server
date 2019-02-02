# 基础镜像
FROM hwtdocker/cloud:alpine-sqlite

MAINTAINER hwt
LABEL tool 1.0.0

RUN mkdir /app/

COPY tool /app/tool
COPY online_conf.json /app/conf.json

WORKDIR /app/
EXPOSE 3030

CMD ["./tool", "-conf=conf.json"]
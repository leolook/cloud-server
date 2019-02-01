# 基础镜像
FROM hwtdocker/cloud:alpine-sqlite

MAINTAINER hwt
LABEL cloud-server 1.0.0

RUN mkdir /app/
RUN mkdir /app/db/

COPY tool /app/tool
COPY online_conf.json /app/conf.json
COPY tool.db /app/tool.db

WORKDIR /app/
EXPOSE 3030

CMD ["./tool", "-conf=conf.json"]
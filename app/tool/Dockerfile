# 基础镜像
FROM hwtdocker/cloud:alpine-sqlite

MAINTAINER hwt
LABEL cloud-server 1.0.0

COPY tool /
COPY online_conf.json conf.json
COPY tool.db /

VOLUME ["tool.db"]
EXPOSE 3030

CMD ["./tool", "-conf=conf.json"]
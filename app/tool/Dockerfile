FROM alpine

MAINTAINER hwt
LABEL cloud-server 1.0.0

COPY tool /
COPY online_conf.json conf.json
COPY tool.db /

EXPOSE 3030

CMD ["./tool", "-conf=conf.json"]
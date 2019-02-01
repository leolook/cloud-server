FROM scratch

MAINTAINER hwt
LABEL cloud-server 1.0.0

COPY tool tool
COPY online_conf.json conf.json

EXPOSE 3030

CMD ["./tool", "-conf=conf.json"]
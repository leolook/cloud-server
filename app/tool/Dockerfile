FROM scratch

MAINTAINER hwt
LABEL cloud-server 1.0.0

WORKDIR /app/tool

COPY tool tool
COPY online_conf.json conf.json

ENTRYPOINT ["./tool", "-conf=conf.json"]

EXPOSE 2030

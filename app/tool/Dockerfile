FROM scratch
MAINTAINER hwt
LABEL cloud 1.0

WORKDIR /app/tool

RUN mkdir logs
COPY tool tool
COPY online_conf.json

CMD ["/tool", "-conf=online_conf.json"]

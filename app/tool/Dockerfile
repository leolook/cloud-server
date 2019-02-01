FROM scratch

MAINTAINER hwt
LABEL cloud-server 1.0.0

RUN mkdir /app
RUN mkdir /app/logs
COPY tool /app/tool
COPY online_conf.json /app/conf.json

WORKDIR /app/
EXPOSE 3030

RUN cp -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

ENTRYPOINT ["./tool", "-conf=conf.json"]



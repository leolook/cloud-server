FROM scratch
MAINTAINER hwt
LABEL cloud-server 1.0.0
WORKDIR /app/tool
RUN mkdir logs
COPY /tool /bin
COPY online_conf.json
CMD ["/bin", "-conf=online_conf.json"]

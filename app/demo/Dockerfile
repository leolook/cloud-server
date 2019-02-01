FROM alpine

MAINTAINER hwt
LABEL cloud-server 1.0.0

RUN mkdir /demo
RUN mkdir /demo/logs

COPY demo /demo/bin
COPY Dockfile /demo/Dockfile

WORKDIR /demo/
EXPOSE 330

CMD ["./bin"]
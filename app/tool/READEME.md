
#### docker 编译
1. docker build --rm -t hwtdocker/cloud:tool .

#### docker 推送
1. docker login
2. docker push hwtdocker/cloud:tool

#### docker 运行
1. docker run -itd -p 0.0.0.0:3030:3030/tcp -v tool.db:/root/go/src/cloud-server/app/tool --name=tool 9c047ad754b4

#### docker 进入容器
1. docker exec -it 5f31b7d05cf8 sh
2.如果出现报错:
  例如: rpc error: code = 2 desc = oci runtime error: exec failed:
      container_linux.go:247: starting container process caused "process_linux.go:110:
      decoding init error from pipe caused \"read parent: connection reset by peer\""
  解决办法: 1. yum update 2.yum downgrade docker docker-client docker-common


#### linux 系统，修改或者添加环境变量
1. vim /etc/profile
2. source /etc/profile 生效

#### golang 代码编译
1. GOOS=linux GOARCH=amd64 go build -tags netgo
2. 注意: 使用apline作为基础镜像,需要带上-tags netgo,不然编译的文件不能在该基础镜像上运行

#### docker 编译
1. docker build --rm -t hwtdocker/cloud:tool .

#### docker 推送
1. docker login
2. docker push hwtdocker/cloud:tool

#### docker 运行
1. docker run -itd -p 0.0.0.0:3030:3030/tcp -v /root/db/app/tool/tool.db:/app/tool.db --name=tool 8a01ed45d39c

#### docker 进入容器
1. docker exec -it 5f31b7d05cf8 sh
2.如果出现报错:
  例如: rpc error: code = 2 desc = oci runtime error: exec failed:
      container_linux.go:247: starting container process caused "process_linux.go:110:
      decoding init error from pipe caused \"read parent: connection reset by peer\""
  解决办法: 1. yum update 2.yum downgrade docker docker-client docker-common

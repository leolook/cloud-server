
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
1. docker run -itd -p 3030:3030/tcp -v /root/db/app/tool/tool.db:/app/tool.db --name=tool 8a01ed45d39c
2. -v 是指将宿主机的文件或者目录挂载到docker内部,这样docker就能访问宿主机的文件;冒号前面是宿主机路径,冒号后面是容器路径,必须都是绝对路径
3. -p 冒号前是宿主机端口,冒号后是容器开放端口

#### docker 进入容器
1. docker exec -it 5f31b7d05cf8 sh
2.如果出现报错:
  例如: rpc error: code = 2 desc = oci runtime error: exec failed:
      container_linux.go:247: starting container process caused "process_linux.go:110:
      decoding init error from pipe caused \"read parent: connection reset by peer\""
  解决办法: 1. yum update 2.yum downgrade docker docker-client docker-common
3. dokcer 进入容器前提是，容器运行是具有-d 后台运行的条件

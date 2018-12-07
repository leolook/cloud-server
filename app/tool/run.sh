#!/bin/sh
###编译
name=tool

mkdir -p ${name}/conf
mkdir -p ${name}/bin

go build main.go

###移动
mv main ${name}/bin/${name}
mv online_conf.json ${name}/conf/conf.json
mv pm2.json ${name}/conf/pm2.json

path=/root/hwt/project/${name}

####移动
mv ${name} ${path}

###切入目录
cd ${path}/conf/

###执行命令
###pm2 start pm2.json

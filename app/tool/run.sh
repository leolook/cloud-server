#!/bin/sh
name=tool

path=/root/hwt/project/${name}

###如果文件夹不存在，创建文件夹
if [ ! -d "${path}/bin" ]; then
   mkdir -p ${path}/bin
fi

if [ ! -d "${path}/conf" ]; then
   mkdir -p ${path}/conf
fi

if [ ! -d "${path}/logs" ]; then
   mkdir -p ${path}/logs
fi

###编译
go build main.go

###移动
mv main ${path}/bin/${name}
cp online_conf.json ${path}/conf/conf.json
cp pm2.json ${path}/conf/pm2.json

###切入目录
cd ${path}/conf/

###执行命令
pm2 start pm2.json

###日志查看
pm2 logs ${name}
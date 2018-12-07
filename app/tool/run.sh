#!/bin/sh
###编译
go build

path=/root/hwt/project/tool

###移动
mv tool ${path}/bin/tool
mv online_conf.json ${path}/conf/conf.json
mv pm2.json ${path}/conf/pm2.json

###切入目录
cd ${path}/conf/

###执行命令
###pm2 start pm2.json

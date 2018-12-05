#!/bin/sh
###编译
go build

###移动
mv tool /root/hwt/project/tool/tool

###切入目录
cd /root/hwt/project/tool

###停止服务
value=`netstat -lnp | grep 2030`
if [ "${value}" != "" ]
then
  tmp=${value#*LISTEN}
  pid=${tmp%%/*}
  echo ${pid}
  kill ${pid}
else
  echo "no server"
fi

###启动
nohup ./tool -server=172.17.199.149:2030 -sqlite=db/tool.db &

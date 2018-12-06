package main

import (
	"cloud-server/app/tool/controller"
	"cloud-server/common/flag"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.New()
	route := g.Group("tool")
	{
		db := controller.DB{}
		dbRoute := route.Group("db")
		dbRoute.POST("createOrModify", db.CreateOrModify) //创建或者编辑
		dbRoute.POST("detail", db.Detail)                 //详情
		dbRoute.POST("page", db.Page)                     //分页
		dbRoute.GET("allName", db.AllName)                //获取db名称
		dbRoute.POST("connect", db.Connect)               //连接
		dbRoute.POST("tableModel", db.TableModel)         //表模型
		dbRoute.POST("del", db.Del)                       //删除
	}
	g.Run(flag.ServerAddr)
}

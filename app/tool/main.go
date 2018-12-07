package main

import (
	"cloud-server/app/tool/controller"
	conf "cloud-server/common/flag"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	var str string
	flag.StringVar(&str, "flagfile", "", "flagfile")
	flag.Parse()
	fmt.Println(str)
}

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
	g.Run(conf.ServerAddr)
}

package main

import (
	"cloud-server/app/tool/web"
	"cloud-server/lib/log"
	"github.com/gin-gonic/gin"
)

func init() {
}

func main() {
	log.Infof("start init")
	g := gin.New()
	r := g.Group("tool")
	r.POST("db/createOrModify", web.DB{}.CreateOrModify) //创建或者编辑
	r.POST("db/detail", web.DB{}.Detail)                 //详情
	r.POST("db/page", web.DB{}.Page)                     //分页
	r.GET("db/allName", web.DB{}.AllName)                //获取db名称
	r.POST("db/connect", web.DB{}.Connect)               //连接
	r.POST("db/tableModel", web.DB{}.TableModel)         //表模型
	r.POST("db/del", web.DB{}.Del)                       //删除
	g.Run("0.0.0.0:2030")
	log.Infof("start server")
}

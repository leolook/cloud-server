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
	r.POST("db/createOrModify", web.DB{}.CreateOrModify)
	r.POST("db/detail", web.DB{}.Detail)
	r.POST("db/page", web.DB{}.Page)
	r.GET("db/allName", web.DB{}.AllName)
	r.POST("db/connect", web.DB{}.Connect)
	r.POST("db/tableModel", web.DB{}.TableModel)
	g.Run("0.0.0.0:2030")
	log.Infof("start server")
}

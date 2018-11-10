package main

import (
	"cloud-server/app/tool/web"
	"github.com/gin-gonic/gin"
)

func init() {
}

func main() {
	g := gin.New()
	r := g.Group("tool")
	r.POST("db/createOrModify", web.DB{}.CreateOrModify)
	r.POST("db/detail", web.DB{}.Detail)
	r.POST("db/page", web.DB{}.Page)
	r.GET("db/allName", web.DB{}.AllName)
	g.Run("0.0.0.0:2030")
}

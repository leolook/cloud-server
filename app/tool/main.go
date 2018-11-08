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
	r.POST("create", web.DB{}.Create)
	g.Run("0.0.0.0:2030")
}

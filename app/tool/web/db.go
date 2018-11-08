package web

import (
	"cloud-server/app/tool/service"
	"cloud-server/common/pb"
	log "cloud-server/lib/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DB struct{}

func (DB) Create(c *gin.Context) {

	//参数解析
	var req pb.CreateReq
	err := c.BindJSON(&req)
	if err != nil {
		log.Errorf("failed to bind json,[err=%v]", err)
		c.JSON(http.StatusOK, pb.ResponseToFail(pb.E_INVALID_ARGUMENT, pb.P_INVALID_ARGUMENT))
		return
	}

	if req.Db == nil {
		c.JSON(http.StatusOK, pb.ResponseToFail(pb.E_INVALID_ARGUMENT, pb.P_INVALID_ARGUMENT))
		return
	}

	//参数校验
	if req.Db.Name == "" {
		log.Warnf("name can no be empty,[name=%v]", req.Db.Name)
		c.JSON(http.StatusOK, pb.ResponseToFail(pb.E_INVALID_ARGUMENT, pb.P_NAME_EMPTY))
		return
	}

	if req.Db.Ip == "" {
		log.Warnf("ip can no be empty,[ip=%v]", req.Db.Ip)
		c.JSON(http.StatusOK, pb.ResponseToFail(pb.E_INVALID_ARGUMENT, pb.P_IP_EMPTY))
		return
	}

	if req.Db.Port <= 0 {
		log.Warnf("port can no be empty,[port=%v]", req.Db.Port)
		c.JSON(http.StatusOK, pb.ResponseToFail(pb.E_INVALID_ARGUMENT, pb.P_PORT_EMPTY))
		return
	}

	if req.Db.UserName == "" {
		log.Warnf("user name can no be empty,[userName=%v]", req.Db.UserName)
		c.JSON(http.StatusOK, pb.ResponseToFail(pb.E_INVALID_ARGUMENT, pb.P_USER_NAME_EMPTY))
		return
	}

	if req.Db.Password == "" {
		log.Warnf("password can no be empty,[password=%v]", req.Db.Password)
		c.JSON(http.StatusOK, pb.ResponseToFail(pb.E_INVALID_ARGUMENT, pb.P_PASSWORD_EMPTY))
		return
	}

	rsp, err := service.DB.Create(pb.NewContext(c), &req)
	if err != nil {
		c.JSON(http.StatusOK, pb.ResponseWithError(err))
		return
	}

	c.JSON(http.StatusOK, pb.ResponseToWin(rsp))
}

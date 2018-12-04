package web

import (
	"cloud-server/app/tool/service"
	"cloud-server/common/pb"
	"cloud-server/lib/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

type DB struct{}

//创建数据库
func (DB) CreateOrModify(c *gin.Context) {

	//参数解析
	var req pb.DbCreateOrModifyReq
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

	//ip校验
	reg := regexp.MustCompile(`((25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d)))\.){3}(25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d)))`)
	isValid := reg.MatchString(req.Db.Ip)
	if !isValid {
		log.Warnf("ip is wrong,[ip=%v]", req.Db.Ip)
		c.JSON(http.StatusOK, pb.ResponseToFail(pb.E_INVALID_ARGUMENT, pb.P_IP_IN_VALID))
		return
	}

	//端口校验
	if req.Db.Port < 1 || req.Db.Port > 65535 {
		log.Warnf("port is not in(1-65532),[port=%d]", req.Db.Port)
		c.JSON(http.StatusOK, pb.ResponseToFail(pb.E_INVALID_ARGUMENT, pb.P_PORT_IN_VALID))
		return
	}

	ctx := pb.NewContext(c)
	var rsp *pb.DbCreateOrModifyRsp

	if req.Db.Id > 0 {
		rsp, err = service.DB.Update(ctx, &req)
	} else {
		rsp, err = service.DB.Create(ctx, &req)
	}

	if err != nil {
		c.JSON(http.StatusOK, pb.ResponseWithError(err))
		return
	}

	c.JSON(http.StatusOK, pb.ResponseToWin(rsp))
}

//数据库详情
func (DB) Detail(c *gin.Context) {
	//参数解析
	var req pb.DbDetailByIDReq
	err := c.BindJSON(&req)
	if err != nil {
		log.Errorf("failed to bind json,[err=%v]", err)
		c.JSON(http.StatusOK, pb.ResponseToFail(pb.E_INVALID_ARGUMENT, pb.P_INVALID_ARGUMENT))
		return
	}

	if req.Id <= 0 {
		c.JSON(http.StatusOK, pb.ResponseToFail(pb.E_INVALID_ARGUMENT, pb.P_INVALID_ARGUMENT))
		return
	}

	rsp, err := service.DB.Detail(pb.NewContext(c), &req)
	if err != nil {
		c.JSON(http.StatusOK, pb.ResponseWithError(err))
		return
	}

	c.JSON(http.StatusOK, pb.ResponseToWin(rsp))
}

//数据库分页
func (DB) Page(c *gin.Context) {
	//参数解析
	var req pb.DbPageReq
	err := c.BindJSON(&req)
	if err != nil {
		log.Errorf("failed to bind json,[err=%v]", err)
		c.JSON(http.StatusOK, pb.ResponseToFail(pb.E_INVALID_ARGUMENT, pb.P_INVALID_ARGUMENT))
		return
	}

	//参数默认
	if req.PageNo <= 0 {
		req.PageNo = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	log.Infof("pageNo=%d,pageSize=%d", req.PageNo, req.PageSize)

	rsp, err := service.DB.Page(pb.NewContext(c), &req)
	if err != nil {
		c.JSON(http.StatusOK, pb.ResponseWithError(err))
		return
	}

	c.JSON(http.StatusOK, pb.ResponseToWin(rsp))
}

//获取所有数据库
func (DB) AllName(c *gin.Context) {
	rsp, err := service.DB.AllName(pb.NewContext(c), &pb.DbAllNameReq{})
	if err != nil {
		c.JSON(http.StatusOK, pb.ResponseWithError(err))
		return
	}

	c.JSON(http.StatusOK, pb.ResponseToWin(rsp))
}

//连接
func (DB) Connect(c *gin.Context) {
	//参数解析
	var req pb.DbConnectReq
	err := c.BindJSON(&req)
	if err != nil {
		log.Errorf("failed to bind json,[err=%v]", err)
		c.JSON(http.StatusOK, pb.ResponseToFail(pb.E_INVALID_ARGUMENT, pb.P_INVALID_ARGUMENT))
		return
	}

	if req.Id <= 0 {
		c.JSON(http.StatusOK, pb.ResponseToFail(pb.E_INVALID_ARGUMENT, pb.P_OPTIN_EMPTY))
		return
	}

	if req.DbName == "" {
		c.JSON(http.StatusOK, pb.ResponseToFail(pb.E_INVALID_ARGUMENT, pb.P_NAME_EMPTY))
		return
	}

	rsp, err := service.DB.Connect(pb.NewContext(c), &req)
	if err != nil {
		c.JSON(http.StatusOK, pb.ResponseWithError(err))
		return
	}

	c.JSON(http.StatusOK, pb.ResponseToWin(rsp))
}

//表模型
func (DB) TableModel(c *gin.Context) {
	//参数解析
	var req pb.DbTableModelReq
	err := c.BindJSON(&req)
	if err != nil {
		log.Errorf("failed to bind json,[err=%v]", err)
		c.JSON(http.StatusOK, pb.ResponseToFail(pb.E_INVALID_ARGUMENT, pb.P_INVALID_ARGUMENT))
		return
	}

	if req.Key == "" {
		c.JSON(http.StatusOK, pb.ResponseToFail(pb.E_INVALID_ARGUMENT, pb.P_KEY_EMPTY))
		return
	}

	if req.Name == "" {
		c.JSON(http.StatusOK, pb.ResponseToFail(pb.E_INVALID_ARGUMENT, pb.P_TABLE_NAME_EMPTY))
		return
	}

	rsp, err := service.DB.TableModel(pb.NewContext(c), &req)
	if err != nil {
		c.JSON(http.StatusOK, pb.ResponseWithError(err))
		return
	}

	c.JSON(http.StatusOK, pb.ResponseToWin(rsp))
}

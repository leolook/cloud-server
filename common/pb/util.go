package pb

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

type Context map[string]string

func (c *Context) Get(k string) string {
	var mp map[string]string = *c
	return mp[k]
}

func NewContext(c *gin.Context) *Context {

	mp := make(map[string]string)
	header := c.Request.Header
	for k, v := range header {
		ru := []rune(k)
		mp[strings.ToLower(string(ru[0]))+string(ru[1:])] = v[0]
	}

	var tmp Context = mp
	return &tmp
}

type rsp struct {
	Code    int32       `json:"code"`
	Message string      `json:"msg,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func ToError(code E, message string) error {
	return fmt.Errorf(`{"code":%d,"msg":"%s"}`, code, message)
}

func ResponseToWin(data interface{}) interface{} {
	return &rsp{
		Code:    int32(E_OK),
		Message: "",
		Data:    data,
	}
}

func ResponseToFail(code E, message string) interface{} {
	return &rsp{
		Code:    int32(E_OK),
		Message: message,
		Data:    nil,
	}
}

func ResponseWithError(err error) interface{} {
	var rsp rsp
	er := json.Unmarshal([]byte(err.Error()), &rsp)
	if er != nil {
		return err.Error()
	}
	return rsp
}

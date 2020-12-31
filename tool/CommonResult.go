package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS int = 0
	FAILED int = 1
)

func Success(context *gin.Context, data interface{}){
	context.JSON(http.StatusOK,map[string]interface{}{
		"code":SUCCESS,
		"msg":"成功",
		"data":data,
	})
}

func Failed(context *gin.Context, data interface{}){
	context.JSON(http.StatusOK,map[string]interface{}{
		"code":FAILED,
		"msg":"失败",
		"data":data,
	})
}

func ReturnData(message string,status bool,redirect bool, args ...interface{}) interface{}{
	data:=make(map[string]interface{})
	data["message"]=message
	data["status"]=status
	data["redirect"]=redirect	//重定向 true为需要重新回到某个地方，false为不需要，默认为false
	data["extra"]=args
	return data
}
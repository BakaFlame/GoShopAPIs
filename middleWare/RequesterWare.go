package middleWare

import (
	"GoShop/model/RequesterModel"
	"github.com/gin-gonic/gin"
)

func RequesterWare() gin.HandlerFunc{
	return func(context *gin.Context) {
		go func(){
			remoteAddress:=context.ClientIP()
			requestUrl:=context.Request.URL.Path
			requestMethod:=context.Request.Method
			//fmt.Println("访问接口\t"+requestUrl)
			//fmt.Println("来访ip\t"+remoteAddress)
			requesterBool,requesterId := RequesterModel.CheckRequester(remoteAddress)
			if requesterBool {
				RequesterModel.AddRequestCount(requesterId,requestUrl,requestMethod)
			} else {
				RequesterModel.AddNewRequester(remoteAddress,requestUrl,requestMethod)
			}
		}()
	}
}

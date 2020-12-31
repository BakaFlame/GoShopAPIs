package middleWare

import (
	"GoShop/model/UserModel"
	"fmt"
	"github.com/gin-gonic/gin"
)

func IsLogin() gin.HandlerFunc{
	return func(context *gin.Context) {
		token:=context.GetHeader("thisisnotatoken")
		fmt.Println(token)
		fmt.Println("寻找token")
		if token == "" {
			context.Abort()
		} else {
			fmt.Println("找到token密钥")
			fmt.Println(token)
			if UserModel.CheckTokenInRedis(token){ //开始匹配cookie
				fmt.Println("redis找到token")
			} else {
				context.Abort()
			}
		}
	}
}

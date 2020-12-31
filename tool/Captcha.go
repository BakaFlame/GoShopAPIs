package tool

import (
	"github.com/gin-gonic/gin"
)

type CaptchaResult struct{
	Id string `json:"id"`
	Base64Blob string `json:"base_64_blob"`
	VerifyValue string `json:"code"`
}

func GenerateCaptcha(context *gin.Context){
	//parameters:=base64Captcha.Driver(base64Captcha.Item())
}

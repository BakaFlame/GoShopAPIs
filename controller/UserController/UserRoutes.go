package UserController

import (
	"GoShop/middleWare"
	"github.com/gin-gonic/gin"
)

func (user *UserStruct) Route(EngineRouter *gin.Engine)  {
	userRoute := EngineRouter.Group("/api/user")
	//userRoute.Use(tool.Cors())	跨域需要在app.default()获取的时候使用，放在这里的话优先级会低，导致option过不去
	{
		userRoute.GET("/loginpage",user.LoginPage)	//登录页面
		userRoute.GET("/registerpage",user.RegisterPage) //注册页面
		userRoute.POST("/register",user.Register)	//注册接口
		userRoute.POST("/sendemailcode",user.SendEmailCode)	//注册接口
		userRoute.POST("/checkusername",user.CheckUsername)	//输入框检测用户名
		userRoute.POST("/checkuseremail",user.CheckUserEmail) //输入框检测邮箱
		userRoute.POST("/login",user.Login)	//登录接口
		userRoute.POST("/logout",middleWare.IsLogin(),user.Logout)	//注销接口
		userRoute.POST("/usersendemailcode",middleWare.IsLogin(),user.UserSendEmailCode)	//用户发送邮箱验证码接口
		userRoute.POST("/resetpasswordwitholdpassword",middleWare.IsLogin(),user.ResetPasswordWithOldPassword)	//重置密码接口
		userRoute.POST("/resetpasswordwithemail",middleWare.IsLogin(),user.ResetPasswordWithEmail)	//重置密码接口
		userRoute.POST("/resetemail",middleWare.IsLogin(),user.ResetEmail)	//重置邮箱接口
	}
}
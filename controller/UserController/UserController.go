package UserController

import (
	"GoShop/model/UserModel"
	"GoShop/tool"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserStruct struct {
}

//登录页面
func (user *UserStruct) LoginPage(context *gin.Context) { //登录页面
	context.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Main website",
	})
}

//登录接口
func (user *UserStruct) Login(context *gin.Context) { //用户登录
	token := context.GetHeader("thisisnotatoken")
	fmt.Println("登录先寻找token避免重复登录设置token")
	if token == "" {
		fmt.Println("token为空")
		username := context.PostForm("username") //获取发送来的数据
		password := context.PostForm("password")
		if username != "" && password != "" {
			password = tool.MD5Encode(context.PostForm("password"))  //加密密码
			loginBool, userId := UserModel.Login(username, password) // 将username和加密后的密码带入登录方法
			if loginBool {                                           //设置cookie
				fmt.Println("设置token")
				token := tool.RandMD5Encode(password) //加密
				userData := tool.RedisUserInfo{
					UserId:    userId,
					Username:  username,
					UserToken: token,
				}
				tokenRes, err := json.Marshal(userData)
				if err != nil {
					fmt.Println(err)
				}
				go UserModel.SetTokenInRedis(token, string(tokenRes))
				fmt.Println(token)
				context.JSON(200, tool.ReturnData("登录成功", loginBool, true, token)) //传给前端
			} else {
				context.JSON(200, tool.ReturnData("登录失败，请检查用户名和密码", loginBool, true)) //传给前端
			}
		} else {
			context.JSON(200, tool.ReturnData("请填写用户名和密码", false, true)) //传给前端
		}
	} else {
		fmt.Println("找到token密钥")
		fmt.Println(token)
		if UserModel.CheckTokenInRedis(token) { //开始匹配cookie
			fmt.Println("redis找到token")
			context.JSON(200, tool.ReturnData("登录成功", true, true, token)) //传给前端
		} else {
			context.JSON(200, tool.ReturnData("token错误", false, false, "")) //传给前端
		}
	}
}

//注册页面
func (user *UserStruct) RegisterPage(context *gin.Context) {
	context.HTML(http.StatusOK, "register.html", gin.H{
		"title": "Main website",
	})
}

//注册时发送邮箱验证码
func (user *UserStruct) RegisterSendEmailCode(context *gin.Context) {
	email := context.PostForm("email")
	if email != "" {
		if UserModel.CheckUserEmailIsTaken(email) == false {
			go tool.SendEmail(email)
			context.JSON(200, tool.ReturnData("已发送验证码", true, false))
		} else {
			context.JSON(200, tool.ReturnData("邮箱已被占用", true, false))
		}
	} else {
		context.JSON(200, tool.ReturnData("请填写邮箱", true, false))
	}
}

//注册接口
func (user *UserStruct) Register(context *gin.Context) {
	email := context.PostForm("email") //接受传入数据
	emailCode := context.PostForm("emailcode")
	username := context.PostForm("username")
	password := context.PostForm("password")
	if email != "" && emailCode != "" && username != "" && password != "" {
		fmt.Println("内容完成进入注册")
		if emailCode == UserModel.VerifyEmailCode(email) {
			fmt.Println("邮箱验证码成功")
			go UserModel.DelCodeInRedis(email)
			password = tool.MD5Encode(password)                              //开始加密密码
			message, status := UserModel.Register(username, password, email) //将username,加密后密码,邮箱放入注册方法
			context.JSON(200, tool.ReturnData(message, status, false))
		} else {
			context.JSON(200, tool.ReturnData("邮箱验证码错误", false, false))
		}
	} else {
		context.JSON(200, tool.ReturnData("请填写完整信息", false, false))
	}
}

//检查用户名是否存在(输入框)
func (user *UserStruct) CheckUsername(context *gin.Context) {
	username := context.PostForm("username") //接受传入数据
	if UserModel.CheckUserNameIsTaken(username) {
		context.JSON(200, tool.ReturnData("用户名已存在", false, false))
	} else {
		context.JSON(200, tool.ReturnData("用户名可用", true, false))
	}
}

//检查邮箱是否存在(输入框)
func (user *UserStruct) CheckUserEmail(context *gin.Context) {
	email := context.PostForm("email") //接受传入数据
	if UserModel.CheckUserEmailIsTaken(email) {
		context.JSON(200, tool.ReturnData("邮箱已存在", false, false))
	} else {
		context.JSON(200, tool.ReturnData("邮箱可用", true, false))
	}
}

//注销接口
func (user *UserStruct) Logout(context *gin.Context) {
	thisisnotacookie := context.GetHeader("thisisnotatoken")
	if thisisnotacookie == "" {
		context.JSON(200, tool.ReturnData("注销失败", true, false))
	} else {
		UserModel.LogoutUserInRedis(thisisnotacookie)
		context.JSON(200, tool.ReturnData("成功注销", true, false))
	}
}

//用户发送邮箱验证码接口
func (user *UserStruct) UserSendEmailCode(context *gin.Context) {
	token := context.GetHeader("thisisnotatoken")
	userId, _ := tool.GetUserInfoFromRedis(token)
	userInfo := UserModel.GetUserInfo(userId)
	go tool.SendEmail(userInfo.Email)
	context.JSON(200, tool.ReturnData("发送成功", true, false))
}

//用户重置密码接口(使用旧密码验证)
func (user *UserStruct) ResetPasswordWithOldPassword(context *gin.Context) { //需要配合前端js 获取referer 从指定页面请求
	oldPassword := context.PostForm("oldpassword")
	newPassword := context.PostForm("newpassword")
	token := context.GetHeader("thisisnotatoken")
	userId, _ := tool.GetUserInfoFromRedis(token)
	oldOne := tool.MD5Encode(oldPassword)
	if UserModel.CheckUserPassword(userId, oldOne) {
		newPassword = tool.MD5Encode(newPassword)
		if UserModel.ResetPassword(userId, newPassword) {
			UserModel.LogoutUserInRedis(token)
			context.JSON(200, tool.ReturnData("成功修改密码，请用新密码登录", true, true))
		} else {
			context.JSON(200, tool.ReturnData("修改失败，请联系管理员", false, false))
		}
	}
}

//用户重置密码接口(使用邮箱验证)
func (user *UserStruct) ResetPasswordWithEmail(context *gin.Context) { //需要配合前端js 获取referer 从指定页面请求
	emailCode := context.PostForm("emailcode")
	newPassword := context.PostForm("newpassword")
	token := context.GetHeader("thisisnotatoken")
	userId, _ := tool.GetUserInfoFromRedis(token)
	userInfo := UserModel.GetUserInfo(userId)
	if emailCode == UserModel.VerifyEmailCode(userInfo.Email) {
		if UserModel.ResetPassword(userId, newPassword) {
			context.JSON(200, tool.ReturnData("成功修改密码，请用新密码登录", true, true))
		} else {
			context.JSON(200, tool.ReturnData("修改失败，请联系管理员", false, false))
		}
	} else {
		context.JSON(200, tool.ReturnData("邮箱验证码错误", false, false))
	}
}

//重置邮箱接口
func (user *UserStruct) ResetEmail(context *gin.Context) {
	emailCode := context.PostForm("emailcode")
	newEmail := context.PostForm("newemail")
	token := context.GetHeader("thisisnotatoken")
	userId, _ := tool.GetUserInfoFromRedis(token)
	userInfo := UserModel.GetUserInfo(userId)
	if emailCode == UserModel.VerifyEmailCode(userInfo.Email) {
		if UserModel.ResetEmail(userId, newEmail) {
			context.JSON(200, tool.ReturnData("成功修改邮箱", true, true))
		} else {
			context.JSON(200, tool.ReturnData("修改失败，请联系管理员", false, false))
		}
	} else {
		context.JSON(200, tool.ReturnData("邮箱验证码错误", false, false))
	}
}

func (user *UserStruct) NoLoginResetPassword(context *gin.Context) {
	email := context.PostForm("email")
	emailCode := context.PostForm("emailcode")
	if emailCode == UserModel.VerifyEmailCode(email) {
		newPassword := context.PostForm("newpassword")
		newPassword = tool.MD5Encode(newPassword)
		if UserModel.NoLoginResetPassword(email, newPassword) {
			context.JSON(200, tool.ReturnData("成功修改密码", true, true))
		} else {
			context.JSON(200, tool.ReturnData("修改失败，请联系管理员", false, false))
		}
	} else {
		context.JSON(200, tool.ReturnData("邮箱验证码错误", false, false))
	}
}

func (user *UserStruct) SendEmailCode(context *gin.Context) {
	email := context.PostForm("email")
	if email != "" {
		if UserModel.CheckUserEmailIsTaken(email) == false {
			go tool.SendEmail(email)
			context.JSON(200, tool.ReturnData("已发送验证码", true, false))
		} else {
			context.JSON(200, tool.ReturnData("没有此邮箱", true, false))
		}
	} else {
		context.JSON(200, tool.ReturnData("请填写邮箱", true, false))
	}
}

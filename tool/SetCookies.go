package tool

//func SetLoginCookies(userId string,username string,token string,context *gin.Context) {
//	fmt.Println("设置cookie")
//	cookieToken:=RandMD5Encode(token) //加密
//	cookie:=UserModel.Cookies{
//		UserId: userId,
//		Username: username,
//		CookieToken: cookieToken,
//	}
//	cookieRes, err := json.Marshal(&cookie)
//	if err != nil {
//		fmt.Println(err)
//	}
//	go UserModel.SetCookieInRedis(cookieToken,string(cookieRes))
//	context.SetCookie("username",username,2592000,"/","127.0.0.1",false, true) //2592000 30天的秒数
//	context.SetCookie("thisisnotacookie",cookieToken,2592000,"/","127.0.0.1",false, true) //2592000 30天的秒数
//	context.SetCookie("userid",userId,2592000,"/","127.0.0.1",false, true) //2592000 30天的秒数
//}

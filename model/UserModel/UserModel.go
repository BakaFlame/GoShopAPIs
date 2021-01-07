package UserModel

import (
	"GoShop/model"
	"GoShop/tool"
	"context"
	"fmt"
	"strconv"
	"time"
)

func Register(username string, password string, email string) (string, bool) {
	if CheckUserNameIsTaken(username) { //检查用户名是否存在，如果是将执行if里面的语句
		fmt.Println("用户名存在")
		return "用户名已存在", false
	}
	if CheckUserEmailIsTaken(email) { //检查邮箱是否存在，如果是将执行if里面的语句
		fmt.Println("邮箱存在")
		return "邮箱已存在", false
	}
	user := &User{
		Username:   username,
		Password:   password,
		Email:      email,
		CreateTime: model.BetterTime{time.Now()},
		UpdateTime: model.BetterTime{time.Now()},
		Confirmed:  1,
	}
	go model.DB.Create(&user) //协程创建用户
	return "注册成功", true
}

func Login(username string, password string) (bool, string) {
	user := &User{}
	sql := "select * from users where username = ? and password = ?"
	model.DB.Raw(sql, username, password).Scan(&user)
	//fmt.Println("loginmodel"+user.Username)
	if user.ID != 0 {
		return true, strconv.Itoa(user.ID)
	} else {
		return false, ""
	}
}

func CheckTokenInRedis(cookieToken string) bool {
	model.SwitchRedisDB(1)
	redisToken, err := model.RedisDB.Get(context.Background(), tool.MD5EncodeWithSalt(cookieToken)).Result()
	if err != nil {
		fmt.Println(err)
	}
	if redisToken != "" {
		return true
	} else {
		fmt.Println("token错误")
		fmt.Println(redisToken)
		return false
	}
}

func SetTokenInRedis(cookieToken string, cookieData string) {
	model.SwitchRedisDB(1)
	token := tool.MD5EncodeWithSalt(cookieToken)
	model.RedisDB.Set(context.Background(), token, cookieData, time.Hour*72)
	//fmt.Println("传入的键:"+cookieToken)
	////token, err :=model.RedisDB.Get(context.Background(),tool.MD5EncodeWithSalt(cookieToken)).Result()
	////if err != nil {
	////	fmt.Println(token)
	////}
	////fmt.Println(token)
	//redisInfo:=tool.RedisUserInfo{}
	//err = json.Unmarshal([]byte(token),&redisInfo)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//
	//}
	//if redisInfo.UserToken == tool.MD5EncodeWithSalt(cookieToken) {
	//	fmt.Println("等于")
	//	model.RedisDB.Set(context.Background(),redisInfo.UserToken,cookieData,time.Hour*72)
	//	fmt.Println(redisInfo.UserToken)
	//	fmt.Println(cookieData)
	//	return redisInfo.UserToken
	//} else {
	//	fmt.Println("不等于 token已经过期")
	//
	//	fmt.Println(token)
	//	fmt.Println(cookieData)
	//	return token
	//}
}

func LogoutUserInRedis(token string) {
	model.SwitchRedisDB(1)
	model.RedisDB.Del(context.Background(), tool.MD5EncodeWithSalt(token))
}

func CheckUserNameIsTaken(username string) bool {
	user := &User{}
	sql := "select id from users where username = ? and status = 0"
	model.DB.Raw(sql, username).Scan(&user)
	if user.ID != 0 {
		return true
	} else {
		return false
	}
}

func CheckUserEmailIsTaken(email string) bool {
	user := &User{}
	sql := "select id from users where email = ? and status = 0 and confirmed = 1"
	model.DB.Raw(sql, email).Scan(&user)
	//fmt.Println(user.ID)
	if user.ID != 0 { //0表示查到的数据为空，而int默认值为0，所以代表找不到数据，返回true
		return true
	} else {
		return false
	}
}

func CheckUserPassword(userId string, password string) bool {
	user := &User{}
	sql := "select id from users where id = ? and password = ?"
	model.DB.Raw(sql, userId, password).Scan(&user)
	if user.ID != 0 {
		return true
	} else {
		return false
	}
}

func GetUserInfo(userId string) User {
	var user User
	sql := "select * from users where id = ? and status = 0"
	model.DB.Raw(sql, userId).Scan(&user)
	return user
}

func VerifyEmailCode(email string) string {
	model.SwitchRedisDB(0)
	code, err := model.RedisDB.Get(context.Background(), email).Result()
	if err != nil {
		fmt.Println(err)
	}
	return code
}

func DelCodeInRedis(email string) bool {
	model.SwitchRedisDB(0)
	row, err := model.RedisDB.Del(context.Background(), email).Result()
	if row == 1 {
		return true
	} else {
		fmt.Println(err)
		return false
	}
}

func ResetPassword(userId string, password string) bool {
	sql := "update users set password = ? where id = ?"
	row := model.DB.Exec(sql, password, userId).RowsAffected
	if row == 1 {
		return true
	} else {
		return false
	}
}

func ResetEmail(userId string, email string) bool {
	sql := "update users set email = ? where id = ?"
	row := model.DB.Exec(sql, email, userId).RowsAffected
	if row == 1 {
		return true
	} else {
		return false
	}
}

func NoLoginResetPassword(email string, newPassword string) bool {
	sql := "update users set password = ? where email = ?"
	row := model.DB.Exec(sql, newPassword, email).RowsAffected
	if row == 1 {
		return true
	} else {
		return false
	}
}

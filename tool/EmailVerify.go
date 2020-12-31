package tool

import (
	"GoShop/model"
	"context"
	"fmt"
	"gopkg.in/gomail.v2"
	"math/rand"
	"strconv"
	"time"
)

func SendEmail(email string) (string,bool){
	m := gomail.NewMessage()
	rand.Seed(time.Now().Unix())
	num := rand.Intn(9000)+1000
	m.SetHeader("From", "2794827531@qq.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "验证码!")
	m.SetBody("text/html", "<p>这是注册验证码哦，时效只有15分钟</p>"+strconv.Itoa(num))
	d := gomail.NewDialer("smtp.qq.com", 587, "2794827531@qq.com", "afgakpieuzkbdcdh")
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	}
	model.SwitchRedisDB(0)
	var context = context.Background()
	row,err:=model.RedisDB.Set(context,email,strconv.Itoa(num),time.Minute*15).Result() //设置15分钟的过期时间
	if err!=nil{
		fmt.Println(err)
		return err.Error(),false
	}
	fmt.Println(row)
	return "已经发送到您的邮箱，请立刻查看",true
}
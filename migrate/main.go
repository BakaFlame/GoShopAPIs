package main

import (
	"GoShop/model/TestModel"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

func main()  {
	DB, err := gorm.Open("mysql", "root:123456@/go_gin?charset=utf8&parseTime=True&loc=Local")
	defer DB.Close() //压入栈中，执行完所有函数后执行
	if err != nil {
		log.Fatal(err.Error())
	}
	DB.AutoMigrate(TestModel.User{},TestModel.Tag{})//添加迁移结构体
	fmt.Println("迁移完毕")
}

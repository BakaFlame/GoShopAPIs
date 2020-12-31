package TestModel

import (
	"GoShop/model"
)

type User struct {
	//属性名一定要大写，不然无法写入数据库，因为没权限
	ID int `gorm:"column:id" json:"id"` //设置列名和传json的属性名，也包括结构体直接传(gorm自带结构体转json格式进行发送数据content.JSON(状态码.结构体))
	Name string `gorm:"column:name" json:"name"`
	Age int `gorm:"column:age" json:"age"`
	Content string `gorm:"column:content" json:"content"`
	Testdelete string `gorm:"column:testdelete" json:"testdelete"`
	CreateTime model.BetterTime `gorm:"column:createTime" json:"createTime"`
	UpdateTime model.BetterTime `gorm:"column:updateTime" json:"updateTime"`
	Operator int `gorm:"column:operator" json:"operator"`
	Tag int `gorm:"column:tag" json:"tag"`
}

type Tag struct {
	//属性名一定要大写，不然无法写入数据库，因为没权限
	ID int `gorm:"column:id" json:"id"`
	TagName string `gorm:"column:tagName" json:"TagName"`
	CreateTime model.BetterTime `gorm:"column:createTime" json:"createTime"`
	UpdateTime model.BetterTime `gorm:"column:updateTime" json:"updateTime"`
	Operator int `gorm:"column:operator" json:"operator"`
}

type WtfResult struct {
	ID int `gorm:"column:id" json:"id"`
	ItemCount int `gorm:"column:itemCount" json:"itemCount"`
	Amount float64 `gorm:"column:amount" json:"amount"`
	CartId int `gorm:"column:cartId" json:"cartId"`
}

type TestStruct struct {
	ID int `gorm:"column:id" json:"id"`
	TestNull string `gorm:"column:testnull" json:"testnull"`
}
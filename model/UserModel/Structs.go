package UserModel

import "GoShop/model"

type User struct {
	ID int `gorm:"column:id" json:"id"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Email string `gorm:"column:email" json:"email"`
	Confirmed int `gorm:"column:confirmed" json:"confirmed"`
	CreateTime model.BetterTime `gorm:"column:create_time" json:"create_time"`
	UpdateTime model.BetterTime `gorm:"column:update_time" json:"update_time"`
	Modifier int `gorm:"column:modifier" json:"modifier"`
	Status int `gorm:"column:status" json:"status"`
}


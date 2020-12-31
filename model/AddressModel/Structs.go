package AddressModel

import "GoShop/model"

type Address struct {
	ID int `gorm:"column:id" json:"id"`
	RealName string `gorm:"column:real_name" json:"real_name"`
	Address string `gorm:"column:address" json:"address"`
	Phone string `gorm:"column:phone" json:"phone"`
	CreateTime model.BetterTime `gorm:"column:create_time" json:"create_time"`
	UpdateTime model.BetterTime `gorm:"column:update_time" json:"update_time"`
	Creator int `gorm:"column:creator" json:"creator"`
	Modifier int `gorm:"column:modifier" json:"modifier"`
	Status int `gorm:"column:status" json:"status"`
	UserId int `gorm:"column:user_id" json:"user_id"`
}

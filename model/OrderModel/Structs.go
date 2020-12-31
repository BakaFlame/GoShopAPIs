package OrderModel

import "GoShop/model"

type Order struct {
	ID int `gorm:"column:id" json:"id"`
	UserId int `gorm:"column:user_id" json:"user_id"`
	AddressId int `gorm:"column:address_id" json:"address_id"`
	TotalPrice float64 `gorm:"column:total_price" json:"total_price"`
	Creator int  `gorm:"column:creator" json:"creator"`
	Modifier int  `gorm:"column:modifier" json:"modifier"`
	CreateTime model.BetterTime `gorm:"column:create_time" json:"create_time"`
	UpdateTime model.BetterTime `gorm:"column:update_time" json:"update_time"`
	Status int `gorm:"column:status" json:"status"`
}

type OrderDetail struct {
	ID int `gorm:"column:id" json:"id"`
	ItemId int `gorm:"column:item_id" json:"item_id"`
	ItemName string `gorm:"column:item_name" json:"item_name"`
	ItemPrice float64 `gorm:"column:item_price" json:"item_price"`
	ItemCount int `gorm:"column:item_count" json:"item_count"`
	ItemImage string `gorm:"column:item_image" json:"item_image"`
	Creator int  `gorm:"column:creator" json:"creator"`
	Modifier int  `gorm:"column:modifier" json:"modifier"`
	CreateTime model.BetterTime `gorm:"column:create_time" json:"create_time"`
	UpdateTime model.BetterTime `gorm:"column:update_time" json:"update_time"`
	Status int `gorm:"column:status" json:"status"`
	OrderId int `gorm:"column:order_id" json:"order_id"`
}

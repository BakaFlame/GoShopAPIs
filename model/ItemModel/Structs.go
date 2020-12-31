package ItemModel

import "GoShop/model"

type Item struct {
	ID int `gorm:"column:id" json:"id"`
	Type int `gorm:"column:type" json:"type"`
	Name string `gorm:"column:name" json:"name"`
	Price float64 `gorm:"column:price" json:"price"`
	Description string `gorm:"column:description" json:"description"`
	ItemLeft int `gorm:"column:item_left" json:"item_left"`
	ItemDetail string `gorm:"column:item_detail" json:"item_detail"`
	CreateTime model.BetterTime `gorm:"column:create_time" json:"create_time"`
	UpdateTime model.BetterTime `gorm:"column:update_time" json:"update_time"`
	Creator int `gorm:"column:creator" json:"creator"`
	Modifier int `gorm:"column:modifier" json:"modifier"`
	Status int `gorm:"column:status" json:"status"`
	ItemImgInfo *ItemImg
	ItemType *ItemType
}

type ItemImg struct {
	ID int `gorm:"column:id" json:"id"`
	Uri string `gorm:"column:uri" json:"uri"`
	ItemId int `gorm:"column:item_id" json:"item_id"`
	CreateTime model.BetterTime `gorm:"column:create_time" json:"create_time"`
	UpdateTime model.BetterTime `gorm:"column:update_time" json:"update_time"`
	Creator int `gorm:"column:creator" json:"creator"`
	Modifier int `gorm:"column:modifier" json:"modifier"`
	Status int `gorm:"column:status" json:"status"`
}

type ItemType struct {
	ID int `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	Description string `gorm:"column:description" json:"description"`
	Creator int `gorm:"column:creator" json:"creator"`
	Modifier int `gorm:"column:modifier" json:"modifier"`
	CreateTime model.BetterTime `gorm:"column:create_time" json:"create_time"`
	UpdateTime model.BetterTime `gorm:"column:update_time" json:"update_time"`
	Status int `gorm:"column:status" json:"status"`
}

type ItemBrand struct {
	ID int `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	TypeId int `gorm:"column:type_id" json:"type_id"`
	CreateTime model.BetterTime `gorm:"column:create_time" json:"create_time"`
	UpdateTime model.BetterTime `gorm:"column:update_time" json:"update_time"`
	Creator int `gorm:"column:creator" json:"creator"`
	Modifier int `gorm:"column:modifier" json:"modifier"`
	Status int `gorm:"column:status" json:"status"`
}

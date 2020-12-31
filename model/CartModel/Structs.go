package CartModel

type Cart struct {
	ID int `gorm:"column:id" json:"id"`
	ItemCount int `gorm:"column:item_count" json:"item_count"`
	ItemId int `gorm:"column:item_id" json:"item_id"`
	UserId int `gorm:"column:user_id" json:"user_id"`
}


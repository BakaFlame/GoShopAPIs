package RequesterModel

import "GoShop/model"

type Requester struct {
	ID int `gorm:"column:id" json:"id"`
	RemoteAddress string `gorm:"column:remote_address" json:"remote_address"`
	RequestCount int `gorm:"column:request_count" json:"request_count"`
	FirstTime model.BetterTime `gorm:"column:first_time" json:"first_time"`
	LastTime model.BetterTime `gorm:"column:last_time" json:"last_time"`
}

type RequestInfo struct {
	ID int `gorm:"column:id" json:"id"`
	Url string `gorm:"column:url" json:"url"`
	RequestTime model.BetterTime `gorm:"column:request_time" json:"request_time"`
	RequestMethod string `gorm:"column:request_method" json:"request_method"`
	RequesterId int `gorm:"column:requester_id" json:"requester_id"`
}

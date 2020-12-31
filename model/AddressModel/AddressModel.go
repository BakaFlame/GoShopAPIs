package AddressModel

import (
	"GoShop/model"
	"fmt"
	"strconv"
	"time"
)

func GetAddressInfoByUserId(userId string) ([]*Address,bool){
	var address []*Address
	//fmt.Println(userId)
	sql:="select id,real_name,address,phone,create_time from addresses where user_id = ? and status = 0 order by create_time"
	model.DB.Raw(sql,userId).Scan(&address)
	if len(address) == 0 {
		return address,false
	} else {
		return address,true
	}
}

func GetAddressInfoById(id string) Address{
	address:=Address{}
	sql:="select id,real_name,address,phone from addresses from addresses where id = ?"
	model.DB.Raw(sql,id).Scan(&address)
	return address
}

func InsertAddressInfo(realName string,address string,phone string,userId string){
	userIntId,_:=strconv.Atoi(userId)
	addressData:=&Address{
		RealName:   realName,
		Address:    address,
		Phone:     phone,
		CreateTime: model.BetterTime{time.Now()},
		UpdateTime: model.BetterTime{time.Now()},
		Creator:    userIntId,
		Modifier:  userIntId,
		Status:     0,
		UserId:     userIntId,
	}
	tx:=model.DB.Begin()
	if err:=tx.Create(&addressData).Error;err != nil {
		fmt.Println(err)
		tx.Rollback()
	}
}

func UpdateAddressById(realName string,address string,phone string,modifier string,addressid string) bool {
	sql := "update addresses set real_name = ?, address = ?, phone = ?,modifier = ? where id = ?"
	row := model.DB.Exec(sql, realName, address, phone,modifier , addressid).RowsAffected
	if row == 1 {
		return true
	} else {
		return false
	}
}

func DeleteAddressById(addressId string) bool {
	sql := "update addresses set status = 1 where id = ?"
	row := model.DB.Exec(sql, addressId).RowsAffected
	if row == 1 {
		return true
	} else {
		return false
	}
}

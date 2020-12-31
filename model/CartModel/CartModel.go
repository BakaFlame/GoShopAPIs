package CartModel

import (
	"GoShop/model"
	"GoShop/model/ItemModel"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func GetCartListByUserId(page int,userId string) interface{} {
	var dataCount int
	model.DB.Model(&Cart{}).Where("user_id = ?",userId).Count(&dataCount)
	maxPage:=math.Ceil(float64(dataCount)/8)	//最大页数(前端能刷新到的)
	sql:="select items.id,items.name,items.price,item_imgs.uri,carts.id,carts.item_count,carts.create_time from items,item_imgs,carts where carts.status = 0 items.id = item_imgs.item_id and items.id = carts.item_id and carts.user_id = "+userId+" order by carts.create_time LIMIT 8 OFFSET "+strconv.Itoa(page*10)
	data,_:=model.QuerySql(sql)
	var dataMap = map[string]interface{}{"maxpage":maxPage,"datacount":dataCount,"data":data}
	return dataMap
}

func CheckItemInCartByItemId(itemId int,userId int) bool {
	cart:=&Cart{}
	fmt.Println(userId)
	model.DB.Raw("select item_id from carts where item_id = ? and user_id = ? and carts.status = 0",itemId,userId).Scan(&cart)
	if cart.ItemId == itemId{
		return true
	} else {
		return false
	}
}

//物品页面操作
func ItemPageAddToCart(itemCount int,itemId int,userId int,InOrUp int) {
	tx := model.DB.Begin()
	if InOrUp == 0 {
		//插入
		cart:=&Cart{
			ItemCount: itemCount,
			ItemId:    itemId,
			UserId:    userId,
		}
		if err:=tx.Create(&cart).Error;err != nil{
			tx.Rollback()
		} else {
			tx.Commit()
		}
	} else {
		itemLeft:=ItemModel.GetItemLeftByItemId(itemId)
		if itemLeft-itemCount>=0 {
			if err:=tx.Exec("update carts set item_count = item_count + ? where item_id = ? and user_id = ?",itemCount,itemId,userId).Error;err!=nil{
				tx.Rollback()
			} else {
				tx.Commit()
			}
		}
	}
}


//购物车页面操作
func DeleteItemInCart(itemId string,userId string){
	tx := model.DB.Begin()
	sql:="update carts set status = 1 where item_id = ? and user_id = ?"
	if err:=tx.Exec(sql,itemId,userId).Error;err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
}

func UnDeleteItemInCart(itemId string,userId string){
	tx := model.DB.Begin()
	sql:="update carts set status = 0 where item_id = ? and user_id = ?"
	if err:=tx.Exec(sql,itemId,userId).Error;err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
}

func MultiDeleteItemInCart(itemId string,userId string){
	itemIds:=strings.Split(itemId, ",")
	for _,v := range itemIds {
		DeleteItemInCart(v,userId)
	}
}

func ChangeItemCountInCart(itemCount string,itemId string,userId string){
	tx := model.DB.Begin()
	sql:="update carts set item_count = ? where item_id = ? user_id = ?"
	if err:=tx.Exec(sql,itemCount,itemId,userId).Error;err!=nil{
		tx.Rollback()
	} else {
		tx.Commit()
	}
}

func GetItemFromCartById(cartId string,userId string) string{
	sql:="select items.id,items.name,items.price,item_imgs.uri,carts.id,carts.item_count, from items,item_imgs,carts where carts.id = "+cartId+" items.id = item_imgs.item_id and items.id = carts.item_id and carts.user_id = "+userId
	str,_:=model.QuerySql(sql)
	return str
}




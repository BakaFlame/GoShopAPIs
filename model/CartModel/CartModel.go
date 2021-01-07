package CartModel

import (
	"GoShop/model"
	"GoShop/model/ItemModel"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func GetCartListByUserId(page int, userId string) interface{} {
	var dataCount int
	model.DB.Model(&Cart{}).Where("user_id = ?", userId).Count(&dataCount)
	maxPage := math.Ceil(float64(dataCount) / 8) //最大页数(前端能刷新到的)
	sql := "select items.id,items.name,items.price,item_imgs.uri,carts.id as 'carts_id',carts.item_count,items.item_left,carts.create_time from items,item_imgs,carts where carts.status = 0 and items.id = item_imgs.item_id and items.id = carts.item_id and carts.user_id = " + userId + " order by carts.create_time LIMIT 8 OFFSET " + strconv.Itoa(page*8)
	data, _ := model.QuerySql(sql)
	var dataMap = map[string]interface{}{"maxpage": maxPage, "datacount": dataCount, "data": data}
	return dataMap
}

func CheckItemInCartByItemId(itemId int, userId int) bool {
	cart := &Cart{}
	//fmt.Println(userId)
	model.DB.Raw("select item_id from carts where item_id = ? and user_id = ? and carts.status = 0", itemId, userId).Scan(&cart)
	if cart.ItemId == itemId {
		return true
	} else {
		return false
	}
}

//物品页面操作
func ItemPageAddToCart(itemCount int, itemId int, userId int, InOrUp int) {
	tx := model.DB.Begin()
	if InOrUp == 0 {
		fmt.Println("插入")
		//插入
		cart := &Cart{
			ItemCount: itemCount,
			ItemId:    itemId,
			UserId:    userId,
		}
		if err := tx.Create(&cart).Error; err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	} else {
		fmt.Println("更新")
		if err := tx.Exec("update carts set item_count = item_count + ? where item_id = ? and user_id = ?", itemCount, itemId, userId).Debug().Error; err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}
}

//购物车页面操作
func DeleteItemInCart(cartId string, userId string) {
	tx := model.DB.Begin()
	sql := "update carts set status = 1 where id = ? and user_id = ?"
	if err := tx.Exec(sql, cartId, userId).Error; err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
}

func UnDeleteItemInCart(cartId string, userId string) {
	tx := model.DB.Begin()
	sql := "update carts set status = 0 where id = ? and user_id = ?"
	if err := tx.Exec(sql, cartId, userId).Error; err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
}

func MultiDeleteItemInCart(itemId string, userId string) {
	itemIds := strings.Split(itemId, ",")
	for _, v := range itemIds {
		DeleteItemInCart(v, userId)
	}
}

func ChangeItemCountInCart(itemCount int, cartId int, userId string) {
	tx := model.DB.Begin()
	sql := "update carts set item_count = ? where cartId = ? user_id = ?"
	if err := tx.Exec(sql, itemCount, cartId, userId).Error; err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
}

func GetItemFromCartById(cartId string, userId string) string {
	//fmt.Println(cartId)
	//fmt.Println(userId)
	sql := "select items.id as 'item_id',items.name,items.price,item_imgs.uri,carts.id as 'cart_id',carts.item_count from items,item_imgs,carts where carts.id = " + cartId + " and items.id = item_imgs.item_id and items.id = carts.item_id and carts.user_id = " + userId
	str, _ := model.QuerySql(sql)
	//fmt.Println(sql)
	return str
}

func CheckItemLeftForCartByItemId(userId string, itemId int, itemCount int) bool {
	item := ItemModel.Item{}
	sql := "select id,item_left from items where id = ?"
	model.DB.Raw(sql, itemId).Scan(&item)
	//fmt.Println(item.ItemLeft)
	cart := Cart{}
	sql = "select id,item_count from carts where item_id = ? and user_id = ?"
	model.DB.Raw(sql, itemId, userId).Scan(&cart)
	if item.ID != 0 {
		if item.ItemLeft >= itemCount+cart.ItemCount {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func CheckItemLeftForCartByCartId(cartId int) (bool, int) {
	item := ItemModel.Item{}
	sql := "select items.id,items.item_left from carts,items where carts.item_id = items.id and carts.id = ?"
	model.DB.Raw(sql, cartId).Scan(&item)
	cart := Cart{}
	sql = "select item_count from carts where id = ?"
	model.DB.Raw(sql, cartId).Scan(&cart)
	if item.ID != 0 {
		if item.ItemLeft >= cart.ItemCount {
			return true, item.ItemLeft
		} else {
			return false, item.ItemLeft
		}
	} else {
		return false, item.ItemLeft
	}
}

//func CheckItemLeftForCartInRedis(userId string,itemId string, itemCount int) (bool,int) {
//	item:=ItemModel.Item{}
//	sql:="select id,item_left from items where id = ?"
//	model.DB.Raw(sql,itemId).Scan(&item)
//	itemCountFromRedis,err:=model.RedisDB.HGet(context.Background(),userId,itemId+":itemcount").Result()
//	itemCountInCart,_:=strconv.Atoi(itemCountFromRedis)
//	if err != nil {
//		fmt.Println(err)
//	}
//	if item.ID != 0 {
//		if item.ItemLeft > itemCount+itemCountInCart {
//			return true, 0
//		} else {
//			return false, item.ItemLeft
//		}
//	} else {
//		return false,0
//	}
//}
//
//func CheckItemInCartByItemIdInRedis(){
//
//}

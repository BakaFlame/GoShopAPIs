package IndexModel

import (
	"GoShop/model"
	"GoShop/model/ItemModel"
	"fmt"
	"math"
	"strconv"
)

//func GetIndexRandTypeItem(typeId int) interface{} {
//	data, _ := model.QuerySql("select items.id,items.name,items.price,items.description,item_imgs.uri FROM items,item_imgs,item_types where items.id = item_imgs.item_id and items.type = item_types.id and items.type = " + strconv.Itoa(typeId) + " and items.id >= (SELECT floor( RAND() * ((SELECT MAX(items.id) FROM items)-(SELECT MIN(items.id) FROM items)) + (SELECT MIN(items.id) FROM items))) ORDER BY items.id LIMIT 4;")
//	return data
//}

func GetIndexRandItem() interface{} {
	data, _ := model.QuerySql("select items.id,items.name,items.price,items.description,item_imgs.uri FROM items,item_imgs where items.id = item_imgs.item_id and items.id >= (SELECT floor( RAND() * ((SELECT MAX(items.id) FROM items)-(SELECT MIN(items.id) FROM items)) + (SELECT MIN(items.id) FROM items))) ORDER BY items.id LIMIT 8;")
	return data
}

func GetIndexNewestItem() interface{} {
	data, _ := model.QuerySql("select items.id,items.name,items.price,items.description,item_imgs.uri FROM items,item_imgs,item_types where items.id = item_imgs.item_id and items.type = item_types.id order by items.create_time desc limit 5")
	return data
}

func GetIndexItemLIst(page int) interface{} { //这个好，1ms内，但是需要前端配合parsejson
	var dataCount int
	fmt.Println(page)
	model.DB.Model(&ItemModel.Item{}).Count(&dataCount)
	maxPage := math.Ceil(float64(dataCount) / 4) //最大页数(前端能刷新到的)
	data, _ := model.QuerySql("select items.id,items.name,items.price,items.description,item_imgs.uri from items,item_imgs where items.status = 0 and items.id = item_imgs.item_id order by items.create_time desc LIMIT 4 OFFSET " + strconv.Itoa(page*4))
	var dataMap = map[string]interface{}{"maxpage": maxPage, "datacount": dataCount, "data": data}
	return dataMap
}

//获取首页banner
func GetIndexBanner() interface{} {
	data, _ := model.QuerySql("select id,uri,create_time from banners where type = 0 order by create_time desc limit 1")
	data2, _ := model.QuerySql("select id,uri,create_time from banners where type = 0 order by create_time desc  limit 4")
	var dataMap = map[string]interface{}{"banner1": data, "banner4": data2}
	return dataMap
}

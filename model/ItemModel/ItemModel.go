package ItemModel

import (
	"GoShop/model"
	"fmt"
	"math"
	"strconv"
)

func GetItemById(itemId string) string { //前端配合parsejson
	sql := "select items.id,items.name,items.price,items.item_left,items.description,items.item_detail,item_types.name from items,item_types where items.type = item_types.id and items.id = " + itemId
	result, _ := model.QuerySql(sql)
	return result
}

func GetItemImgById(itemId string) string { //前端配合parsejson
	sql := "select id,uri from item_imgs where item_id = " + itemId
	result, _ := model.QuerySql(sql)
	return result
}

func GetItemList(page int) interface{} { //这个好，1ms内，但是需要前端配合parsejson
	var dataCount int
	model.DB.Model(&Item{}).Count(&dataCount)
	maxPage := math.Ceil(float64(dataCount) / 32) //最大页数(前端能刷新到的)
	data, _ := model.QuerySql("select items.id,items.name,items.price,items.description,item_imgs.uri from items,item_imgs where items.status = 0 and items.id = item_imgs.item_id order by items.create_time desc LIMIT 32 OFFSET " + strconv.Itoa(page*32))
	var dataMap = map[string]interface{}{"maxpage": maxPage, "datacount": dataCount, "data": data}
	return dataMap
}

func GetItemListByType(typeId string, page int) interface{} { //但是需要前端配合parsejson
	var dataCount int
	model.DB.Model(&Item{}).Where("type = ? and status = ?", typeId, 0).Count(&dataCount)
	maxPage := math.Ceil(float64(dataCount) / 32) //最大页数(前端能刷新到的)
	data, err := model.QuerySql("select items.id,items.name,items.price,items.description,item_imgs.uri,items.create_time from items,item_imgs,item_types where items.status = 0 and items.id = item_imgs.item_id and items.type = item_types.id and items.type = " + typeId + " order by items.create_time desc LIMIT 32 OFFSET " + strconv.Itoa(page*32))
	if err != nil {
		fmt.Println(err)
	}
	var dataMap = map[string]interface{}{"maxpage": maxPage, "datacount": dataCount, "data": data}
	return dataMap
}

func GetItemListByItemName(itemName string, page int) interface{} {
	var dataCount int
	model.DB.Model(&Item{}).Where("name = ? and status = ?", itemName, 0).Count(&dataCount)
	maxPage := math.Ceil(float64(dataCount) / 32) //最大页数(前端能刷新到的)
	data, err := model.QuerySql("select DISTINCT items.id,items.name,items.price,items.description,item_imgs.uri,items.create_time from items,item_imgs where items.status = 0 and items.id = item_imgs.item_id and items.name like '%" + itemName + "%' order by items.create_time desc LIMIT 32 OFFSET " + strconv.Itoa(page*32))
	if err != nil {
		fmt.Println(err)
	}
	var dataMap = map[string]interface{}{"maxpage": maxPage, "datacount": dataCount, "data": data}
	return dataMap
}

func GetItemListByBrand(brandId string, page int) interface{} {
	var dataCount int
	sql := "select DISTINCT items.id,items.name,items.price,items.description,item_imgs.uri,items.create_time from items,item_imgs,item_brands where items.status = 0 and items.id = item_imgs.item_id and items.brand = item_brands.id and items.brand = " + brandId + " order by items.create_time desc LIMIT 32 OFFSET " + strconv.Itoa(page*32)
	//fmt.Println(sql)
	data, err := model.QuerySql(sql)
	model.DB.Model(&Item{}).Where("brand = ? and status = ?", brandId, 0).Count(&dataCount)
	maxPage := math.Ceil(float64(dataCount) / 32) //最大页数(前端能刷新到的)
	if err != nil {
		fmt.Println(err)
	}
	var dataMap = map[string]interface{}{"maxpage": maxPage, "datacount": dataCount, "data": data}
	return dataMap
}

func GetItemListByTypeAndBrand(typeId string, brandId string, page int) interface{} {
	var dataCount int
	sql := "select DISTINCT items.id,items.name,items.price,items.description,item_imgs.uri,items.create_time from items,item_imgs,item_brands,item_types where items.id = item_imgs.item_id and items.brand = item_brands.id and items.type = item_types.id and items.brand = " + brandId + " and items.type = " + typeId + " order by items.create_time desc LIMIT 32 OFFSET " + strconv.Itoa(page*32)
	//fmt.Println(sql)
	data, err := model.QuerySql(sql)
	maxPage := math.Ceil(float64(dataCount) / 32) //最大页数(前端能刷新到的)
	if err != nil {
		fmt.Println(err)
	}
	var dataMap = map[string]interface{}{"maxpage": maxPage, "datacount": dataCount, "data": data}
	return dataMap
}

func GetItemLeftByItemId(itemId int) int {
	item := &Item{}
	sql := "select item_left from items where id = ?"
	model.DB.Raw(sql, itemId).Scan(&item)
	if item.ID != 0 {
		return item.ItemLeft
	} else {
		return 0
	}
}

func GetIBrandList() interface{} {
	itemBrand := []ItemBrand{}
	sql := "select id,name,type_id from item_brands where status = 0"
	model.DB.Raw(sql).Scan(&itemBrand)
	return itemBrand
}

func GetTypeList() interface{} {
	itemType := []ItemType{}
	sql := "select id,name,description from item_types where status = 0"
	model.DB.Raw(sql).Scan(&itemType)
	return itemType
}

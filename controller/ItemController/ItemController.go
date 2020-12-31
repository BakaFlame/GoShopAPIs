package ItemController

import (
	"GoShop/model/ItemModel"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ItemStruct struct {
}

func (itemStruct *ItemStruct) JsonItemPage(context *gin.Context) {
	context.HTML(http.StatusOK, "json.html", gin.H{
		"title": "Main website",
	})
}

//按照种类和品牌查询商品列表
func (itemStruct *ItemStruct) GetItemListByTypeAndBrand(context *gin.Context) {
	pageData := context.PostForm("page") //page从0开始
	fmt.Println("这是页数" + pageData)
	page, _ := strconv.Atoi(pageData)
	typeBool := false
	brandBool := false
	typeId := context.PostForm("typeid")
	fmt.Println("这是种类id" + typeId)
	if typeId != "" {
		typeBool = true
	}
	brandId := context.PostForm("brandid")
	fmt.Println("这是品牌id" + brandId)
	if brandId != "" {
		brandBool = true
	}
	fmt.Println(page)
	fmt.Println(typeBool)
	fmt.Println(brandBool)
	if typeBool && brandBool {
		fmt.Println("两个都查")
		context.JSON(200, ItemModel.GetItemListByTypeAndBrand(typeId, brandId, page))
	} else if typeBool {
		fmt.Println("种类")
		context.JSON(200, ItemModel.GetItemListByType(typeId, page))
	} else if brandBool {
		fmt.Println("品牌")
		context.JSON(200, ItemModel.GetItemListByBrand(brandId, page))
	} else {
		fmt.Println("普通")
		context.JSON(200, ItemModel.GetItemList(page))
	}
}

//按照名字来查询商品列表
func (itemStruct *ItemStruct) GetItemListByItemName(context *gin.Context) {
	itemName := context.PostForm("itemname")
	pageData := context.PostForm("page") //page从0开始
	page, _ := strconv.Atoi(pageData)
	context.JSON(200, ItemModel.GetItemListByItemName(itemName, page))
}

//获取商品详情信息
func (itemStruct *ItemStruct) GetItemInfoById(context *gin.Context) {
	itemId := context.PostForm("itemid")
	var dataMap = map[string]interface{}{"item": ItemModel.GetItemById(itemId),"item_img":ItemModel.GetItemImgById(itemId)}
	context.JSON(200, dataMap)
}

//获取品牌列表
func (itemStruct *ItemStruct) GetBrandList(context *gin.Context) {
	context.JSON(200, ItemModel.GetIBrandList())
}

//获取种类列表
func (itemStruct *ItemStruct) GetTypeList(context *gin.Context) {
	context.JSON(200, ItemModel.GetTypeList())
}

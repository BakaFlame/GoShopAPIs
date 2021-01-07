package CartController

import (
	"GoShop/model/CartModel"
	"GoShop/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CartStruct struct {
}

//获取用户得购物车列表
func (cart CartStruct) GetCartListByUserId(context *gin.Context) { //通过
	token := context.GetHeader("thisisnotatoken")
	userId, _ := tool.GetUserInfoFromRedis(token)
	fmt.Println("userId为:" + userId)
	page, _ := strconv.Atoi(context.PostForm("page"))
	context.JSON(200, CartModel.GetCartListByUserId(page, userId))
}

//在商品详情进入购物车
func (cart CartStruct) ItemPageAddToCart(context *gin.Context) { //通过
	token := context.GetHeader("thisisnotatoken")
	redisUserId, _ := tool.GetUserInfoFromRedis(token)
	userId, _ := strconv.Atoi(redisUserId)
	itemId, _ := strconv.Atoi(context.PostForm("itemid"))
	itemCount, _ := strconv.Atoi(context.PostForm("itemcount"))
	itemBool := CartModel.CheckItemLeftForCartByItemId(redisUserId, itemId, itemCount)
	//fmt.Println(itemCount)
	var message string
	if itemBool {
		message = "已加入购物车"
	} else {
		context.JSON(200, tool.ReturnData("已经超出当前物品库存数量", false, false))
		return
	}
	insertOrUpdate := 0
	if CartModel.CheckItemInCartByItemId(itemId, userId) {
		insertOrUpdate = 1
	}
	go CartModel.ItemPageAddToCart(itemCount, itemId, userId, insertOrUpdate)
	context.JSON(200, tool.ReturnData(message, true, false))
}

//购物车界面操作	参考淘宝
func (cart CartStruct) UpdateCart(context *gin.Context) {
	operation := context.PostForm("operation")
	token := context.GetHeader("thisisnotatoken")
	userId, _ := tool.GetUserInfoFromRedis(token)
	if operation == "delete" { //发送operation来确认操作，delete单个删除
		cartId := context.PostForm("cartid")
		go CartModel.DeleteItemInCart(cartId, userId)
		context.JSON(200, tool.ReturnData("已删除物品", true, false))
	} else if operation == "update" { //update更改当前物品的数量(从数量框里的数量获取直接set)
		cartId, _ := strconv.Atoi(context.PostForm("cartid"))
		itemCount, _ := strconv.Atoi(context.PostForm("itemcount"))
		if itemCount >= 1 {
			itemBool := CartModel.CheckItemLeftForCartByItemId(userId, cartId, itemCount)
			var message string
			if itemBool {
				message = "已修改数量"
			} else {
				context.JSON(200, tool.ReturnData("已经超出当前物品库存数量", false, false))
				return
			}
			CartModel.ChangeItemCountInCart(itemCount, cartId, userId)
			context.JSON(200, tool.ReturnData(message, true, false))
		} else {
			context.JSON(200, tool.ReturnData("输入的数量有误", false, false))
			return
		}
	} else if operation == "undelete" { //undelete撤回刚刚删除的，需要前端保存上一次删除的物品id，一次能多个撤回，取决于上一次删除的物品数量，前端进行隐藏处理或者保存被删除的html代码于vue中
		cartId := context.PostForm("cartid")
		go CartModel.UnDeleteItemInCart(cartId, userId)
		context.JSON(200, tool.ReturnData("已撤回刚刚删除的物品", true, false))
	} else if operation == "multidelete" { //multidelete多选删除，itemid格式为1,2,3 使用逗号隔开发送数据
		cartId := context.PostForm("cartid")
		go CartModel.MultiDeleteItemInCart(cartId, userId)
		context.JSON(200, tool.ReturnData("已删除物品", true, false))
	}
}

//func (cart CartStruct) DeleteItemInCart(context *gin.Context)  {
//	itemId:=context.PostForm("itemid")
//	userId,_:=context.Cookie("userid")
//	go CartModel.DeleteItemInCart(itemId,userId)
//	context.JSON(200,tool.ReturnData("已删除物品",true,false))
//}
//
//func (cart CartStruct) MultiDeleteItemInCart(context *gin.Context)  {
//	itemId:=context.PostForm("itemid")
//	userId,_:=context.Cookie("userid")
//	go CartModel.MultiDeleteItemInCart(itemId,userId)
//	context.JSON(200,tool.ReturnData("已删除物品",true,false))
//}
//
//func (cart CartStruct) ChangeItemCountInCart(context *gin.Context)  {
//	itemId:=context.PostForm("itemid")
//	userId,_:=context.Cookie("userid")
//	itemCount:=context.PostForm("itemcount")
//	CartModel.ChangeItemCountInCart(itemCount,itemId,userId)
//	context.JSON(200,tool.ReturnData("已修改数量",true,false))
//}

//func (cart CartStruct) ItemPageAddToRedisCart(context *gin.Context) {
//	token:=context.GetHeader("thisisnotatoken")
//	redisUserId, _ := tool.GetUserInfoFromRedis(token)
//	userId,_:=strconv.Atoi(redisUserId)
//	itemId:=context.PostForm("itemid")
//	itemCount,_:=strconv.Atoi(context.PostForm("itemcount"))
//	itemBool,itemLeft:=CartModel.CheckItemLeftForCartInRedis(redisUserId,itemId,itemCount)
//	var message string
//	if itemBool {
//		message = "已加入购物车"
//	} else {
//		message = "已加入物品的最大数量"
//		itemCount = itemLeft
//	}
//	insertOrUpdate:=0
//	if CartModel.CheckItemInCartByItemIdInRedis(itemId,userId) {
//		fmt.Println("更新")
//		insertOrUpdate=1
//	} else {
//		fmt.Println("插入")
//	}
//	go CartModel.ItemPageAddToCartInRedis(itemCount,itemId,userId,insertOrUpdate)
//	context.JSON(200,tool.ReturnData(message,true,false))
//}

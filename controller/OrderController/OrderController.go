package OrderController

import (
	"GoShop/model/AddressModel"
	"GoShop/model/CartModel"
	"GoShop/model/OrderModel"
	"GoShop/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type OrderStruct struct {
}

//获取订单列表接口
func (order OrderStruct) GetOrderList(context *gin.Context) {
	orderType := context.PostForm("ordertype")
	pageData := context.PostForm("page")
	page, _ := strconv.Atoi(pageData)
	token := context.GetHeader("thisisnotatoken")
	userId, _ := tool.GetUserInfoFromRedis(token)
	//userId := "16"
	context.JSON(200, OrderModel.GetOrderList(userId, orderType, page))
}

//确认订单前页面接口
func (order OrderStruct) OrderPage(context *gin.Context) {
	token := context.GetHeader("thisisnotatoken")
	userId, _ := tool.GetUserInfoFromRedis(token)
	address, existBool := AddressModel.GetAddressInfoByUserId(userId)
	itemId := context.PostForm("itemid")
	itemCount := context.PostForm("itemcount")
	//fmt.Println(itemCount)
	if itemId == "0" { //根据js get拿到上级页面传参的itemid之后判断是否是0 是0就直接判断从cartpage过来的
		cartDatas := context.PostForm("cartid")
		fmt.Println(cartDatas)
		cartData := strings.Split(cartDatas, ",")
		str := ""
		itemBool := false
		for i := 0; i < len(cartData); i++ {
			cartId, _ := strconv.Atoi(cartData[i])
			itemBool, _ = CartModel.CheckItemLeftForCartByCartId(cartId)
			//fmt.Println(cartData[i])
			if itemBool {
				str += CartModel.GetItemFromCartById(cartData[i], userId)
				//fmt.Println("成功")
			}
		}
		finalstr := strings.Replace(str, "][", ",", -1)
		data := map[string]interface{}{"address:": address, "item": finalstr, "exist": existBool}
		context.JSON(200, data)
	} else { //如果有参数就是从itempage直接购买来的 并且获取itemcount的数量进行查询之后前端自行计算
		item_id, _ := strconv.Atoi(itemId)
		item_count, _ := strconv.Atoi(itemCount)
		itemBool := CartModel.CheckItemLeftForCartByItemId(userId, item_id, item_count)
		if itemBool {
			fmt.Println("进来itembool")
			str := OrderModel.OrderPageList(itemId)
			finalstr := strings.Replace(str, "}", `,"itemcount":"`+itemCount+`"}`, -1)
			data := map[string]interface{}{"address:": address, "item": finalstr}
			context.JSON(200, data)
		} else {
			context.JSON(200, tool.ReturnData("当前物品超过仓库数量", false, false))
		}
	}
}

//提交订单 fromWhere:1.cartpage购物车页面结算,2.itempage从商品详情直接购买
func (order OrderStruct) PutOrder(context *gin.Context) {
	token := context.GetHeader("thisisnotatoken")
	redisUserId, _ := tool.GetUserInfoFromRedis(token)
	//redisUserId:="16"
	addressData := context.PostForm("addressid")
	itemIds := context.PostForm("itemid")
	itemCount := context.PostForm("itemcount") //如果从商品直接购买来的应该有参数发送，如果为0就是从购物车来的
	fromWhere := context.PostForm("fromwhere")
	addressId, _ := strconv.Atoi(addressData)
	userId, _ := strconv.Atoi(redisUserId)
	itemData := strings.Split(itemIds, ",")
	if len(itemData) == 1 && fromWhere == "itempage" {
		go OrderModel.PutOrder(userId, addressId, itemData[0], itemCount)
		context.JSON(200, tool.ReturnData("提交订单成功", true, true))
	}
	if len(itemData) >= 1 && fromWhere == "cartpage" {
		go OrderModel.PutMultipleOrder(userId, addressId, itemData)
		context.JSON(200, tool.ReturnData("提交订单成功", true, true))
	}
	//if len(itemData) == 1 {
	//
	//	context.JSON(200,tool.ReturnData("提交订单成功",true,true))
	//} else if len(itemData) > 1 {
	//	for i := 0; i < len(itemData); i++ {
	//		OrderModel.PutOrder(userId,addressId,itemData,fromWhere,itemCount)
	//	}
	//	context.JSON(200,tool.ReturnData("提交订单成功",true,true))
	//} else {
	//	context.JSON(200,tool.ReturnData("出现错误",false,false))
	//}
}

func (order OrderStruct) Refund(context *gin.Context) {
	orderId := context.PostForm("orderid")
	orderDetailId := context.PostForm("orderdetailid")
	if orderId != "" {
		orderInfo := OrderModel.GetOrderInfo(orderId)
		if orderInfo.Status == 3 {
			if OrderModel.RefundAllItems(orderId) {
				context.JSON(200, tool.ReturnData("退款已提交，请等待处理", true, false))
			} else {
				context.JSON(200, tool.ReturnData("退款全部物品出现错误", false, false))
			}
		} else {
			context.JSON(200, tool.ReturnData("退款状态出现错误", false, false))
		}
	} else if orderDetailId != "" {
		orderDetail := OrderModel.GetOrderIdByDetailId(orderDetailId)
		orderInfo := OrderModel.GetOrderInfo(strconv.Itoa(orderDetail.ID))
		if orderInfo.Status == 3 && orderDetail.Status == 0 {
			if OrderModel.RefundSingleItem(orderDetailId) {
				context.JSON(200, tool.ReturnData("退款已提交，请等待处理", true, false))
			} else {
				context.JSON(200, tool.ReturnData("退款全部物品出现错误", false, false))
			}
		} else {
			context.JSON(200, tool.ReturnData("退款状态出现错误", false, false))
		}
	} else {
		context.JSON(200, tool.ReturnData("你来这干啥的", false, true))
	}
}

func (order OrderStruct) DeleteOrder(context *gin.Context) {
	orderId := context.PostForm("orderid")
	if OrderModel.DeleteOrder(orderId) {
		context.JSON(200, tool.ReturnData("删除订单成功", true, false))
	} else {
		context.JSON(200, tool.ReturnData("删除订单失败", false, false))
	}
}

func (order OrderStruct) ChangeAddress(context *gin.Context) {
	orderId := context.PostForm("orderid")
	orderInfo := OrderModel.GetOrderInfo(orderId)
	if orderInfo.Status == 1 {
		addressId := context.PostForm("addressid")
		if OrderModel.ChangeAddress(addressId, orderId) {
			context.JSON(200, tool.ReturnData("更改地址成功", true, false))
		} else {
			context.JSON(200, tool.ReturnData("更改地址失败", false, false))
		}
	} else {
		context.JSON(200, tool.ReturnData("状态有误，更改地址", false, false))
	}
}

func (order OrderStruct) FinishOrder(context *gin.Context) {
	orderId := context.PostForm("orderid")
	orderInfo := OrderModel.GetOrderInfo(orderId)
	if orderInfo.Status == 2 {
		if OrderModel.FinishOrder(orderId) {
			context.JSON(200, tool.ReturnData("确认收货成功", true, false))
		} else {
			context.JSON(200, tool.ReturnData("确认收货失败", false, false))
		}
	} else {
		context.JSON(200, tool.ReturnData("状态有误，更改地址", false, false))
	}
}

func (order OrderStruct) OrderDetailPage(context *gin.Context) {
	orderId := context.PostForm("orderid")
	context.JSON(200, OrderModel.OrderDetailPage(orderId))
}

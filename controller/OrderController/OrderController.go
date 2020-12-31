package OrderController

import (
	"GoShop/model/AddressModel"
	"GoShop/model/CartModel"
	"GoShop/model/ItemModel"
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
func (order OrderStruct) GetOrderList(context *gin.Context)  {
	orderType:=context.PostForm("ordertype")
	pageData:=context.PostForm("page")
	page,_:=strconv.Atoi(pageData)
	token:=context.GetHeader("thisisnotatoken")
	userId, _ := tool.GetUserInfoFromRedis(token)
	context.JSON(200,OrderModel.GetOrderList(userId,orderType,page))
}

//确认订单前页面接口
func (order OrderStruct) OrderPage(context *gin.Context)  {
	token:=context.GetHeader("thisisnotatoken")
	userId, _ := tool.GetUserInfoFromRedis(token)
	address,existBool:=AddressModel.GetAddressInfoByUserId(userId)
	itemId:=context.PostForm("itemid")
	itemCount:=context.PostForm("itemcount")
	fmt.Println(itemCount)
	if itemId == "0"{	//根据js get拿到上级页面传参的itemid之后判断是否是0 是0就直接判断从cartpage过来的
		cartDatas:=context.PostForm("cartid")
		cartData:=strings.Split(cartDatas,",")
		str:=""
		for i := 0; i < len(cartData) ; i++ {
			str+=CartModel.GetItemFromCartById(cartData[i],userId)
		}
		finalstr:=strings.Replace(str,"][",",",-1)
		data:=map[string]interface{}{"address:":address,"item":finalstr,"exist":existBool}
		context.JSON(200,data)
	} else {	//如果有参数就是从itempage直接购买来的 并且获取itemcount的数量进行查询之后前端自行计算
		str:=ItemModel.GetItemById(itemId)
		finalstr:=strings.Replace(str,"}",",itemcount:"+itemCount+"}]",-1)
		data:=map[string]interface{}{"address:":address,"item":finalstr}
		context.JSON(200,data)
	}
}

//提交订单
func (order OrderStruct) PutOrder(context *gin.Context)  {
	token:=context.GetHeader("thisisnotatoken")
	redisUserId, _ := tool.GetUserInfoFromRedis(token)
	userId,_:=strconv.Atoi(redisUserId)
	addressData:=context.PostForm("addressid")
	addressId,_:=strconv.Atoi(addressData)
	itemDatas:=context.PostForm("itemid")
	itemData:=strings.Split(itemDatas,",")
	fromWhere:=context.PostForm("fromwhere")	//从哪里过来的，直接购买还是购物车，如果直接购买就不做购物车查询
	itemCount:=context.PostForm("itemcount")//如果从商品直接购买来的应该有参数发送，如果为0就是从购物车来的
	if len(itemData) == 1 {
		go OrderModel.PutOrder(userId,addressId,itemDatas,fromWhere,itemCount)
		context.JSON(200,tool.ReturnData("提交订单成功",true,true))
	} else if len(itemData) > 1 {
		for i := 0; i < len(itemData); i++ {
			go OrderModel.PutOrder(userId,addressId,itemData[i],fromWhere,itemCount)
		}
		context.JSON(200,tool.ReturnData("提交订单成功",true,true))
	} else {
		context.JSON(200,tool.ReturnData("出现错误",false,false))
	}
}

func (order OrderStruct) Refund(context *gin.Context)  {
	orderId:=context.PostForm("orderid")
	orderDetailId:=context.PostForm("orderdetailid")
	if orderId != "" {
		orderInfo:=OrderModel.GetOrderInfo(orderId)
		if orderInfo.Status == 3 {
			if OrderModel.RefundAllItems(orderId) {
				context.JSON(200,tool.ReturnData("退款已提交，请等待处理",true,false))
			} else {
				context.JSON(200,tool.ReturnData("退款全部物品出现错误",false,false))
			}
		} else {
			context.JSON(200,tool.ReturnData("退款状态出现错误",false,false))
		}
	} else if orderDetailId != "" {
		orderDetail:=OrderModel.GetOrderIdByDetailId(orderDetailId)
		orderInfo:=OrderModel.GetOrderInfo(strconv.Itoa(orderDetail.ID))
		if orderInfo.Status == 3 && orderDetail.Status == 1 {
			if OrderModel.RefundSingleItem(orderDetailId) {
				context.JSON(200,tool.ReturnData("退款已提交，请等待处理",true,false))
			} else {
				context.JSON(200,tool.ReturnData("退款全部物品出现错误",false,false))
			}
		} else {
			context.JSON(200,tool.ReturnData("退款状态出现错误",false,false))
		}
	} else {
		context.JSON(200,tool.ReturnData("你来这干啥的",false,true))
	}
}

func (order OrderStruct) ChangeAddress(context *gin.Context)  {
	orderId:=context.PostForm("orderid")
	orderInfo:=OrderModel.GetOrderInfo(orderId)
	if orderInfo.Status == 1 {
		addressId:=context.PostForm("addressid")
		if OrderModel.ChangeAddress(addressId,orderId) {
			context.JSON(200,tool.ReturnData("更改地址成功",true,false))
		} else {
			context.JSON(200,tool.ReturnData("更改地址失败",false,false))
		}
	} else {
		context.JSON(200,tool.ReturnData("状态有误，更改地址",false,false))
	}
}

////直接购买提交订单
//func (order OrderStruct) PutOrder(context *gin.Context)  {
//	cookieUserId,_:=context.Cookie("userid")
//	userId,_:=strconv.Atoi(cookieUserId)
//	addressData:=context.PostForm("addressid")
//	addressId,_:=strconv.Atoi(addressData)
//	itemId:=context.PostForm("itemid")
//	fromWhere:=context.PostForm("fromwhere")
//	itemPageCount:=context.PostForm("itempagecount")
//	go OrderModel.PutOrder(userId,addressId,itemId,fromWhere,itemPageCount)
//	context.JSON(200,tool.ReturnData("提交订单成功",true,true))
//}

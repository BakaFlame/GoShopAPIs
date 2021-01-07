package OrderController

import (
	"github.com/gin-gonic/gin"
)

func (order *OrderStruct) Route(EngineRouter *gin.Engine) {
	orderRoutes := EngineRouter.Group("/api/order")
	orderRoutes.Use()
	{
		orderRoutes.POST("/getorderlist", order.GetOrderList)   //获取订单列表接口
		orderRoutes.POST("/putorder", order.PutOrder)           //提交订单接口
		orderRoutes.POST("/orderpage", order.OrderPage)         //提交订单前页面数据
		orderRoutes.POST("/refund", order.Refund)               //退款接口
		orderRoutes.POST("/changeaddress", order.ChangeAddress) //更改地址接口
		orderRoutes.POST("/deleteorder", order.DeleteOrder)     //删除订单接口
	}
}

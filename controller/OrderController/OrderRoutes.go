package OrderController

import (
	"GoShop/middleWare"
	"GoShop/tool"
	"github.com/gin-gonic/gin"
)

func (order *OrderStruct) Route(EngineRouter *gin.Engine)  {
	orderRoutes := EngineRouter.Group("/api/order")
	orderRoutes.Use(tool.Cors(),middleWare.IsLogin())
	{
		orderRoutes.POST("/getorderlist",order.GetOrderList)	//获取订单列表接口
		orderRoutes.POST("/putorder",order.PutOrder)	//购物车提交订单
		orderRoutes.POST("/orderpage",order.OrderPage)	//提交订单前页面
		orderRoutes.POST("/refund",order.Refund)	//提交订单前页面
		orderRoutes.POST("/changeaddress",order.ChangeAddress)	//提交订单前页面
	}
}

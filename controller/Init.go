package controller

import (
	"GoShop/controller/AddressController"
	"GoShop/controller/CartController"
	"GoShop/controller/IndexController"
	"GoShop/controller/ItemController"
	"GoShop/controller/OrderController"
	"GoShop/controller/TestController"
	"GoShop/controller/UserController"
	"github.com/gin-gonic/gin"
)

//在此手动添加路由 new(包名.结构体).Route(EngineRouter)
func RegisterRouter(EngineRouter *gin.Engine){
	new(TestController.TestStruct).Route(EngineRouter)
	new(CartController.CartStruct).Route(EngineRouter)
	new(UserController.UserStruct).Route(EngineRouter)
	new(AddressController.AddressStruct).Route(EngineRouter)
	new(ItemController.ItemStruct).Route(EngineRouter)
	new(OrderController.OrderStruct).Route(EngineRouter)
	new(IndexController.IndexStruct).Route(EngineRouter)
}
package OrderModel

import (
	"GoShop/model"
	"GoShop/model/CartModel"
	"GoShop/model/ItemModel"
	"fmt"
	"math"
	"strconv"
	"time"
)

func GetOrderList(userId string,orderType string,page int) interface{}{
	var sql string
	var dataCount int
	if orderType == "all" {	//查询所有的物品
		sql="select orders.id,orders.total_price,orders.create_time,order_details.item_id,order_details.item_name,order_details.item_price,order_details.item_count,order_details.item_image,order_details.status from orders,order_details where orders.id = order_details.order_id and orders.status <> 4 and orders.user_id = "+userId+" order by orders.create_time desc LIMIT 10 OFFSET "+strconv.Itoa(page*10)
		model.DB.Model(&Order{}).Where("status <> ?",4).Count(&dataCount)
	} else if orderType == "unfinished" {	//查询未完成的订单（待付款，待发货，已发货）
		sql="select orders.id,orders.total_price,orders.create_time,order_details.item_id,order_details.item_name,order_details.item_price,order_details.item_count,order_details.item_image,order_details.status from orders,order_details where orders.id = order_details.order_id and orders.status = 0 or orders.status = 1 or orders.status = 2 and orders.user_id = "+userId+" order by orders.create_time desc LIMIT 10 OFFSET "+strconv.Itoa(page*10)
		model.DB.Model(&Order{}).Where("status = ? or status = ? or status = ?",0,1,2).Count(&dataCount)
	} else if orderType == "finished" {	//查询完成的订单
		sql="select orders.id,orders.total_price,orders.create_time,order_details.item_id,order_details.item_name,order_details.item_price,order_details.item_count,order_details.item_image,order_details.status from orders,order_details where orders.id = order_details.order_id and orders.status = 3  and orders.user_id = "+userId+" order by orders.create_time desc LIMIT 10 OFFSET "+strconv.Itoa(page*10)
		model.DB.Model(&Order{}).Where("status = ?",3).Count(&dataCount)
	}
	maxPage:=math.Ceil(float64(dataCount)/10)
	data,_:=model.QuerySql(sql)
	var dataMap = map[string]interface{}{"maxpage":maxPage,"datacount":dataCount,"data":data}
	return dataMap
}

//差删除购物车里对应的项	fromWhere:1.cartpage购物车页面结算,2.itempage从商品详情直接购买
func PutOrder(userId int,addressId int,itemId string,fromWhere string,itemCount string){
	item_id,_:=strconv.Atoi(itemId)
	item:=ItemModel.Item{}
	itemImg:=ItemModel.ItemImg{}
	cart:=CartModel.Cart{}
	model.DB.Raw("select * from items where id = ?",itemId).Scan(&item)
	model.DB.Raw("select * from item_imgs where item_Id = ? order by id desc LIMIT 1",itemId).Scan(&itemImg)
	order_detail:=OrderDetail{
		ItemId:     item_id,
		ItemName:   item.Name,
		ItemPrice:  item.Price,
		ItemImage:  itemImg.Uri,
		Creator:    userId,
		Modifier:   userId,
		CreateTime: model.BetterTime{time.Now()},
		UpdateTime: model.BetterTime{time.Now()},
		Status:     0,
	}
	order:=&Order{
		UserId:     userId ,
		AddressId:  addressId,
		Creator:userId,
		Modifier:userId,
		CreateTime: model.BetterTime{time.Now()},
		UpdateTime: model.BetterTime{time.Now()},
		Status:     0,
	}
	if fromWhere == "cartpage" {
		model.DB.Raw("select * from carts where item_id = ? and user_id = ?",itemId,userId).Scan(&cart)
		order.TotalPrice=float64(cart.ItemCount) * item.Price
		order_detail.ItemCount = cart.ItemCount
		model.DB.Exec("delete from carts where item_id = ? and user_id = ?",itemId,userId)
	} else if fromWhere == "itempage"{
		itemcount,_ := strconv.Atoi(itemCount)
		order.TotalPrice=float64(itemcount) * item.Price
		order_detail.ItemCount = itemcount
	} else {
		fmt.Println("你从哪里来的???")
		return
	}
	tx := model.DB.Begin()
	if err:=tx.Create(&order).Error;err != nil {
		tx.Rollback()
		fmt.Println("order错误")
		fmt.Println(err)
		return
	}
	order_detail.OrderId=order.ID
	//fmt.Println(order_detail)
	if err:=tx.Create(&order_detail).Error;err != nil {
		tx.Rollback()
		fmt.Println("detail错误")
		fmt.Println(err)
		return
	}
	tx.Commit()	//commit一定是最后执行，不能上传完一个发现没有回滚立刻commit，不然会报错事物回滚。
}

func GetOrderInfo(orderId string) Order {
	order:=Order{}
	sql:="select * from orders where id = ?"
	model.DB.Raw(sql,orderId).Scan(&order)
	return order
}

func GetOrderIdByDetailId(orderDetailId string) OrderDetail {
	orderDetail:=OrderDetail{}
	sql:="select order_id,status from order_details where id = ?"
	model.DB.Raw(sql,orderDetailId).Scan(&orderDetail)
	return orderDetail
}

func GetIOrderItemInfo(detailId string) OrderDetail {
	orderDetail:=OrderDetail{}
	sql:="select * from order_details where id = ?"
	model.DB.Raw(sql,detailId).Scan(&orderDetail)
	return orderDetail
}

func RefundAllItems(orderId string) bool {
	sql:="update order_details set status = 1 where order_id = ?"
	rows:=model.DB.Exec(sql,orderId).RowsAffected
	if rows >=1 {
		return true
	} else {
		return false
	}
}

func RefundSingleItem(orderDetailId string) bool {
	sql:="update order_details set status = 1 where id = ?"
	rows:=model.DB.Exec(sql,orderDetailId).RowsAffected
	if rows ==1 {
		return true
	} else {
		return false
	}
}

func ChangeAddress(addressId string, orderId string) bool {
	sql := "update orders set address_id = ? where id = ?"
	rows:=model.DB.Exec(sql,addressId,orderId).RowsAffected
	if rows >=1 {
		return true
	} else {
		return false
	}
}
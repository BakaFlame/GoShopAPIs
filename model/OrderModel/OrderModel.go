package OrderModel

import (
	"GoShop/model"
	"GoShop/model/AddressModel"
	"GoShop/model/CartModel"
	"GoShop/model/ItemModel"
	"fmt"
	"math"
	"strconv"
	"time"
)

func GetOrderList(userId string, orderType string, page int) interface{} {
	var orderSql string
	var dataCount int
	fmt.Println("ordertype为" + orderType)
	if orderType == "all" { //查询所有的物品
		orderSql = "select id,user_id,address_id,create_time,update_time from orders where orders.status <> 4 and user_id = " + userId + " order by create_time desc LIMIT 10 OFFSET " + strconv.Itoa(page*10)
		model.DB.Model(&Order{}).Where("status <> ? and user_id = ?", 4, userId).Count(&dataCount)
	} else if orderType == "unfinished" { //查询未完成的订单（待付款，待发货，已发货）
		//orderSql = "select orders.id,orders.create_time,order_details.item_id,order_details.item_name,order_details.item_price,order_details.item_count,order_details.item_image,order_details.status from orders,order_details where orders.id = order_details.order_id and (orders.status = 0 or orders.status = 1 or orders.status = 2) and orders.user_id = " + userId + " order by orders.create_time desc LIMIT 10 OFFSET " + strconv.Itoa(page*10)
		orderSql = "select id,user_id,address_id,create_time,update_time from orders where (orders.status = 0 or orders.status = 1 or orders.status = 2) and orders.user_id = " + userId + " order by orders.create_time desc LIMIT 10 OFFSET " + strconv.Itoa(page*10)
		model.DB.Model(&Order{}).Where("status = ? or status = ? or status = ? and user_id = ?", 0, 1, 2, userId).Count(&dataCount)
	} else if orderType == "finished" { //查询完成的订单
		orderSql = "select id,user_id,address_id,create_time,update_time from orders where status = 3  and user_id = " + userId + " order by create_time desc LIMIT 10 OFFSET " + strconv.Itoa(page*10)
		model.DB.Model(&Order{}).Where("status = ? and user_id = ?", 3, userId).Count(&dataCount)
	}
	fmt.Println(orderSql)
	maxPage := math.Ceil(float64(dataCount) / 10)
	//data:=[]interface{}{}
	order := []Order{}
	orderDetail := []OrderDetail{}
	model.DB.Raw(orderSql).Scan(&order)
	for i := 0; i < len(order); i++ {
		sql := "select id,item_id,item_name,item_price,item_count,item_image from order_details where order_id = " + strconv.Itoa(order[i].ID)
		model.DB.Raw(sql).Scan(&orderDetail)
		order[i].OderDetails = orderDetail
	}
	//data,_:=model.QuerySql(sql)
	var dataMap = map[string]interface{}{"maxpage": maxPage, "datacount": dataCount, "data": order}
	return dataMap
}

func PutOrder(userId int, addressId int, itemId string, itemCount string) {
	item_id, _ := strconv.Atoi(itemId)
	item := ItemModel.Item{}
	itemImg := ItemModel.ItemImg{}
	model.DB.Raw("select * from items where id = ?", itemId).Scan(&item)
	model.DB.Raw("select * from item_imgs where item_Id = ? order by id desc LIMIT 1", itemId).Scan(&itemImg)
	orderDetail := OrderDetail{
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
	order := &Order{
		UserId:     userId,
		AddressId:  addressId,
		Creator:    userId,
		Modifier:   userId,
		CreateTime: model.BetterTime{time.Now()},
		UpdateTime: model.BetterTime{time.Now()},
		Status:     0,
	}
	itemcount, _ := strconv.Atoi(itemCount)
	orderDetail.ItemCount = itemcount
	tx := model.DB.Begin()
	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		fmt.Println("order错误")
		fmt.Println(err)
		return
	}
	orderDetail.OrderId = order.ID
	if err := tx.Create(&orderDetail).Error; err != nil {
		tx.Rollback()
		fmt.Println("detail错误")
		fmt.Println(err)
		return
	}
	tx.Commit() //commit一定是最后执行，不能上传完一个发现没有回滚立刻commit，不然会报错事物回滚。
}

//fromWhere:1.cartpage购物车页面结算,2.itempage从商品详情直接购买
func PutMultipleOrder(userId int, addressId int, itemIds []string) {
	order := &Order{
		UserId:     userId,
		AddressId:  addressId,
		Creator:    userId,
		Modifier:   userId,
		CreateTime: model.BetterTime{time.Now()},
		UpdateTime: model.BetterTime{time.Now()},
		Status:     0,
	}
	tx := model.DB.Begin()
	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		fmt.Println("order错误")
		fmt.Println(err)
		return
	}
	var orderId = order.ID
	for i := 0; i < len(itemIds); i++ {
		id := itemIds[i]
		itemId, _ := strconv.Atoi(id)
		item := ItemModel.Item{}
		itemImg := ItemModel.ItemImg{}
		cart := CartModel.Cart{}
		model.DB.Raw("select id,name,price,price from items where id = ?", itemId).Scan(&item)
		model.DB.Raw("select uri from item_imgs where item_Id = ? order by id desc LIMIT 1", itemId).Scan(&itemImg)
		model.DB.Raw("select * from carts where item_id = ? and user_id = ?", itemId, userId).Scan(&cart)
		model.DB.Exec("delete from carts where item_id = ? and user_id = ?", itemId, userId)
		orderDetail := OrderDetail{
			ItemId:     itemId,
			ItemName:   item.Name,
			ItemPrice:  item.Price,
			ItemImage:  itemImg.Uri,
			Creator:    userId,
			Modifier:   userId,
			CreateTime: model.BetterTime{time.Now()},
			UpdateTime: model.BetterTime{time.Now()},
			Status:     0,
		}
		orderDetail.ItemCount = cart.ItemCount
		orderDetail.OrderId = orderId
		if err := tx.Create(&orderDetail).Error; err != nil {
			tx.Rollback()
			fmt.Println("detail错误")
			fmt.Println(err)
			return
		}
	}
	tx.Commit() //commit一定是最后执行，不能上传完一个发现没有回滚立刻commit，不然会报错事物回滚。
}

func GetOrderInfo(orderId string) Order {
	order := Order{}
	sql := "select * from orders where id = ?"
	model.DB.Raw(sql, orderId).Scan(&order)
	return order
}

func GetOrderIdByDetailId(orderDetailId string) OrderDetail {
	orderDetail := OrderDetail{}
	sql := "select order_id,status from order_details where id = ?"
	model.DB.Raw(sql, orderDetailId).Scan(&orderDetail)
	return orderDetail
}

//func GetIOrderItemInfo(detailId string) OrderDetail {
//	orderDetail := OrderDetail{}
//	sql := "select * from order_details where id = ?"
//	model.DB.Raw(sql, detailId).Scan(&orderDetail)
//	return orderDetail
//}

func RefundAllItems(orderId string) bool {
	sql := "update order_details set status = 1 where order_id = ?"
	rows := model.DB.Exec(sql, orderId).RowsAffected
	if rows >= 1 {
		return true
	} else {
		return false
	}
}

func RefundSingleItem(orderDetailId string) bool {
	sql := "update order_details set status = 1 where id = ?"
	rows := model.DB.Exec(sql, orderDetailId).RowsAffected
	if rows == 1 {
		return true
	} else {
		return false
	}
}

func ChangeAddress(addressId string, orderId string) bool {
	sql := "update orders set address_id = ? where id = ?"
	rows := model.DB.Exec(sql, addressId, orderId).RowsAffected
	if rows >= 1 {
		return true
	} else {
		return false
	}
}

func OrderPageList(itemId string) string {
	sql := "select items.id,items.name,items.price,items.description,item_imgs.uri from items,item_imgs where items.id = item_imgs.item_id and items.id = " + itemId + " limit 1"
	result, _ := model.QuerySql(sql)
	return result
}

func DeleteOrder(orderId string) bool {
	sql := "update set orders status = 4 where id = ?"
	row := model.DB.Exec(sql, orderId).RowsAffected
	if row == 1 {
		return true
	} else {
		return false
	}
}

func FinishOrder(orderId string) bool {
	sql := "update set orders status = 3 where id = ?"
	row := model.DB.Exec(sql, orderId).RowsAffected
	if row == 1 {
		return true
	} else {
		return false
	}
}

func OrderDetailPage(orderId string) interface{} {
	order := Order{}
	sql := "select id,address_id,user_id,status from orders where id = ?"
	model.DB.Raw(sql, orderId).Scan(&order)
	address := AddressModel.Address{}
	sql = "select id,real_name,address,phone from addresses where id = ?"
	model.DB.Raw(sql, order.AddressId).Scan(&address)
	orderDetail := []OrderDetail{}
	sql = "select id,item_id,item_name,item_price,item_count,item_image,status from order_details where order_id = ?"
	model.DB.Raw(sql, orderId).Scan(&orderDetail)
	var dataMap = map[string]interface{}{"address": address, "order": order, "order_detail": orderDetail}
	return dataMap
}

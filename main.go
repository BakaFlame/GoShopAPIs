package main

import (
	"GoShop/initPack"
)

func main() {
	app, cfg, db, db_query, redisDB := initPack.InitWeb()
	defer db.Close()                         //压入栈中，主线程结束时关闭数据库防止内存泄漏
	defer db_query.Close()                   //压入栈中，主线程结束时关闭数据库防止内存泄漏
	defer redisDB.Close()                    //压入栈中，主线程结束时关闭数据库防止内存泄漏
	app.Run(cfg.AppHost + ":" + cfg.AppPort) //（注意，一定要关防火墙）如果要开公网，把cfg.AppHost删除 使用:端口开启 默认开启公网和内网
}

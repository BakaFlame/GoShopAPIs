package RequesterModel

import (
	"GoShop/model"
	"fmt"
	"time"
)

func CheckRequester(remoteAdress string) (bool,int) {
	requester:=Requester{}
	sql:="select id,remote_address from requesters where remote_address = ?"
	//fmt.Println(sql)
	model.DB.Raw(sql,remoteAdress).Scan(&requester)
	//fmt.Println(site.ID)
	//fmt.Println(remoteAdress)
	if requester.ID != 0 {
		//fmt.Println("有这个ip")
		return true,requester.ID
	} else {
		//fmt.Println("无这个ip")
		return false,0
	}
}

func AddRequestCount(requesterId int,requestUrl string,requestMethod string) {
	sql:="update requesters set request_count = request_count + 1 where id = ?"
	tx:=model.DB.Begin()
	updateRequesterError:=tx.Exec(sql,requesterId).Error
	if  updateRequesterError != nil{
		tx.Rollback()
		fmt.Println("更新请求者失败")
	}

	requestInfo:=RequestInfo{
		Url:         requestUrl,
		RequestTime: model.BetterTime{time.Now()},
		RequestMethod: requestMethod,
		RequesterId:   requesterId,
	}
	AddRequestInfoError:=tx.Create(&requestInfo).Error
	if AddRequestInfoError != nil {
		tx.Rollback()
		fmt.Println("新增请求信息失败")
	}
	tx.Commit()
}

func AddNewRequester(remoteAddress string,requestUrl string,requestMethod string) {
	requester:=Requester{
		RemoteAddress: remoteAddress,
		RequestCount:    1,
		FirstTime:    model.BetterTime{time.Now()},
		LastTime:    model.BetterTime{time.Now()},
	}
	tx:=model.DB.Begin()
	AddRequesterError:=tx.Create(&requester).Error
	if AddRequesterError != nil {
		tx.Rollback()
		fmt.Println("新增请求者失败")
	} else {
		requestInfo:=RequestInfo{
			Url:         requestUrl,
			RequestTime: model.BetterTime{time.Now()},
			RequestMethod: requestMethod,
			RequesterId:  requester.ID,
		}
		AddRequestInfoError:=tx.Create(&requestInfo).Error
		if AddRequestInfoError != nil {
			tx.Rollback()
			fmt.Println("新增请求信息失败")
		}
		tx.Commit()
	}
}

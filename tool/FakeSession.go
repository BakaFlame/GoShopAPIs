package tool

import (
	"GoShop/model"
	"context"
	"encoding/json"
	"fmt"
)

type RedisUserInfo struct {
	UserId string		`json:"userid"`
	Username string	`json:"username"`
}

func GetUserInfoFromRedis(token string) (string,string){
	model.SwitchRedisDB(1)
	redisToken,err:=model.RedisDB.Get(context.Background(),MD5EncodeWithSalt(token)).Result()
	if err != nil {
		fmt.Println(err)
	}
	userData:=RedisUserInfo{}
	json.Unmarshal([]byte(redisToken),&userData)
	return userData.UserId,userData.Username
}

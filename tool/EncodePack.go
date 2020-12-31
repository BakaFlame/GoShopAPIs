package tool

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var salt = "salty salt"

func MD5Encode (text string) string{
	var md5ctx = md5.New()
	md5ctx.Write([]byte(text))
	return hex.EncodeToString(md5ctx.Sum(nil))
}

func RandMD5Encode (text string) string{
	var md5ctx = md5.New()
	md5ctx.Write([]byte(text))
	rand.Seed(time.Now().Unix())
	randNum:=rand.Intn(10000)
	return MD5Encode(MD5Encode(text)+strconv.Itoa(randNum))
}

func MD5EncodeWithSalt (text string) string{
	var md5ctx = md5.New()
	md5ctx.Write([]byte(text))
	return hex.EncodeToString(md5ctx.Sum([]byte(salt)))
}

func SHA256EncodeWithSalt(text string) string {
	var sha256ctx = sha256.New()
	sha256ctx.Write([]byte(text))
	return hex.EncodeToString(sha256ctx.Sum([]byte(salt)))
}

func SHA256Decode(text string) {
	var sha256ctx = sha256.New()
	sha256ctx.Write([]byte(text))
	var temp,_ = hex.DecodeString(hex.EncodeToString(sha256ctx.Sum([]byte(salt))))
	fmt.Printf("%s\n", temp)
}

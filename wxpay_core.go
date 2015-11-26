package go_wxpay

import (
	"encoding/xml"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
)

func NewOrderNo() string {
	return fmt.Sprintf("%s%d", time.Now().Format("20060102150405"), RandInt(10000, 99999))
}

func RandInt(min int, max int) int {
	if max-min <= 0 {
		return min
	}
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func unifiedorderSign(kv *Kvpairs, wxPayConfig *WxPayConfig) string {

	//貌似现在只有md5？
	stringA := createLinkStringNoUrl(kv)
	stringSignTemp := stringA + "&key=" + wxPayConfig.KEY
	mysign := md5String(stringSignTemp)
	return strings.ToUpper(mysign)

}

func ParseUResponse(result []byte) (*UnifiedorderResp, error) {

	uResp := &UnifiedorderResp{}
	var err error

	err = xml.Unmarshal([]byte(result), uResp)
	if err != nil {
		log.Printf("xml Unmarshal err is %v", err)
		return uResp, err
	}

	log.Printf("Appid is %v", uResp.Appid)

	return uResp, nil
}

func CallbackUResponse(result []byte) (*UnifiedorderResp, error) {
	u, err := ParseUResponse(result)
	if err != nil {
		log.Println("解析返回参数错误")
		return u, err
	}
	return u, err
}

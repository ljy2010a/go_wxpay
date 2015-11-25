package wxpay

import (
	// "encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"
)

type Unifiedorder struct {
	//appid	 appid  是	String(32)	wxd678efh567hg6787	微信分配的公众账号ID（企业号corpid即为此appId）
	Appid string `xml:"appid"`
	//商户号	mch_id	是	String(32)	1230000109	微信支付分配的商户号
	Mch_id string `xml:"mch_id"`
	//随机字符串	nonce_str 	是	String(32)	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	随机字符串，不长于32位
	Nonce_str string `xml:"nonce_str"`
	//签名	sign	是	String(32)	C380BEC2BFD727A4B6845133519F3AD6	签名，详见签名生成算法
	Sign string `xml:"sign"`
	//
	Total_fee int64 `xml:"total_fee"`
}

func (u *Unifiedorder) SetSign() string {

	p := Kvpairs{}
	ref := reflect.ValueOf(u).Elem()
	typ := reflect.TypeOf(u).Elem()
	n := typ.NumField()
	for i := 0; i < n; i++ {
		f := typ.Field(i)
		name := f.Tag.Get("xml")
		v := fmt.Sprintf("%v", ref.Field(i))
		// log.Println(name)
		// log.Println(f.Name)
		// log.Printf("%v = %v", name, v)
		p = append(p, Kvpair{name, v})
	}

	paraFilter(&p)
	argSort(&p)

	sign := unifiedorderSign(&p, GWxPayConfig)
	p = append(p, Kvpair{`sign`, sign})
	xml := createXmlString(&p)
	log.Println("UnifiedorderRequestXml : ", xml)
	return xml
}

func (u *Unifiedorder) TakeOrder() (*UnifiedorderResp, error) {
	reqParam := u.SetSign()
	log.Println("微信预下单参数 : ", reqParam)
	req, _ := http.NewRequest("POST", UnifiedorderApi, strings.NewReader(reqParam))
	req.Header.Set("Content-Type", "application/xml ")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("请求微信预下单失败 ", err)
	}
	bodyByte, _ := ioutil.ReadAll(resp.Body)
	log.Println("微信预下单返回结果 : ", string(bodyByte))
	return CallbackUResponse(bodyByte)
}

type UnifiedorderResp struct {
	Return_code string `xml:"return_code"`
	Return_msg  string `xml:"return_msg"`
	Appid       string `xml:"appid"`
	Mch_id      string `xml:"mch_id"`
	Nonce_str   string `xml:"nonce_str"`
	Sign        string `xml:"sign"`
	Result_code string `xml:"result_code"`
	Prepay_id   string `xml:"prepay_id"`
	Trade_type  string `xml:"trade_type"`
}

// 设备号	device_info	否	String(32)	013467007045764	终端设备号(门店号或收银设备ID)，注意：PC网页或公众号内支付请传"WEB"
// 商品描述	body	是	String(32)	Ipad mini  16G  白色	商品或支付单简要描述
// 商品详情	detail	否	String(8192)	Ipad mini  16G  白色	商品名称明细列表
// 附加数据	attach	否	String(127)	深圳分店	附加数据，在查询API和支付通知中原样返回，该字段主要用于商户携带订单的自定义数据
// 商户订单号	out_trade_no	是	String(32)	20150806125346	商户系统内部的订单号,32个字符内、可包含字母, 其他说明见商户订单号
// 货币类型	fee_type	否	String(16)	CNY	符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
// 总金额	total_fee	是	Int	888	订单总金额，单位为分，详见支付金额
// 终端IP	spbill_create_ip	是	String(16)	123.12.12.123	APP和网页支付提交用户端ip，Native支付填调用微信支付API的机器IP。
// 交易起始时间	time_start	否	String(14)	20091225091010	订单生成时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。其他详见时间规则
// 交易结束时间	time_expire	否	String(14)	20091227091010
// 订单失效时间，格式为yyyyMMddHHmmss，如2009年12月27日9点10分10秒表示为20091227091010。其他详见时间规则
// 注意：最短失效时间间隔必须大于5分钟
// 商品标记	goods_tag	否	String(32)	WXG	商品标记，代金券或立减优惠功能的参数，说明详见代金券或立减优惠
// 通知地址	notify_url	是	String(256)	http://www.weixin.qq.com/wxpay/pay.php	接收微信支付异步通知回调地址
// 交易类型	trade_type	是	String(16)	JSAPI	取值如下：JSAPI，NATIVE，APP，详细说明见参数规定
// 商品ID	product_id	否	String(32)	12235413214070356458058	trade_type=NATIVE，此参数必传。此id为二维码中包含的商品ID，商户自行定义。
// 指定支付方式	limit_pay	否	String(32)	no_credit	no_credit--指定不能使用信用卡支付
// 用户标识	openid	否	String(128)	oUpF8uMuAJO_M2pxb1Q9zNjWeS6o	trade_type=JSAPI，此参数必传，用户在商户appid下的唯一标识。openid如何获取，可参考【获取openid】。企业号请使用【企业号OAuth2.0接口】获取企业号内成员userid，再调用【企业号userid转openid接口】进行转换

package wxpay

import (
	"fmt"
	"log"
	"reflect"
)

type NaviteNotify struct {
	// 返回状态码	return_code	是	String(16)	SUCCESS
	// SUCCESS/FAIL
	// 此字段是通信标识，非交易标识，交易是否成功需要查看result_code来判断
	return_code string `xml:"return_code"`
	// 返回信息	return_msg	否	String(128)	签名失败
	// 返回信息，如非空，为错误原因
	// 签名失败
	// 参数格式校验错误
	Return_msg string `xml:"return_msg"`
	// 公众账号ID	appid	是	String(32)	wx8888888888888888	微信分配的公众账号ID（企业号corpid即为此appId）
	Appid string `xml:"appid"`
	// 商户号	mch_id	是	String(32)	1900000109	微信支付分配的商户号
	Mch_id string `xml:"mch_id"`
	// 设备号	device_info	否	String(32)	013467007045764	微信支付分配的终端设备号，
	Device_info string `xml:"device_info"`
	// 随机字符串	nonce_str	是	String(32)	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	随机字符串，不长于32位
	Nonce_str string `xml:"nonce_str"`
	// 签名	sign	是	String(32)	C380BEC2BFD727A4B6845133519F3AD6	签名，详见签名算法
	Sign string `xml:"sign"`
	// 业务结果	result_code	是	String(16)	SUCCESS	SUCCESS/FAIL
	Result_code string `xml:"result_code"`
	// 错误代码	err_code	否	String(32)	SYSTEMERROR	错误返回的信息描述
	Err_code string `xml:"err_code"`
	// 错误代码描述	err_code_des	否	String(128)	系统错误	错误返回的信息描述
	Err_code_des string `xml:"err_code_des"`
	// 用户标识	openid	是	String(128)	wxd930ea5d5a258f4f	用户在商户appid下的唯一标识
	Openid string `xml:"openid"`
	// 是否关注公众账号	is_subscribe	否	String(1)	Y	用户是否关注公众账号，Y-关注，N-未关注，仅在公众账号类型支付有效
	Is_subscribe string `xml:"is_subscribe"`
	// 交易类型	trade_type	是	String(16)	JSAPI	JSAPI、NATIVE、APP
	Trade_type string `xml:"trade_type"`
	// 付款银行	bank_type	是	String(16)	CMC	银行类型，采用字符串类型的银行标识，银行类型见银行列表
	Bank_type string `xml:"bank_type"`
	// 总金额	total_fee	是	Int	100	订单总金额，单位为分
	Total_fee string `xml:"total_fee"`
	// 货币种类	fee_type	否	String(8)	CNY	货币类型，符合ISO4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	Fee_type string `xml:"fee_type"`
	// 现金支付金额	cash_fee	是	Int	100	现金支付金额订单现金支付金额，详见支付金额
	Cash_fee string `xml:"cash_fee"`
	// 现金支付货币类型	cash_fee_type	否	String(16)	CNY	货币类型，符合ISO4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	Cash_fee_type string `xml:"cash_fee_type"`
	// 代金券或立减优惠金额	coupon_fee	否	Int	10	代金券或立减优惠金额<=订单总金额，订单总金额-代金券或立减优惠金额=现金支付金额，详见支付金额
	Coupon_fee string `xml:"coupon_fee"`
	// 代金券或立减优惠使用数量	coupon_count	否	Int	1	代金券或立减优惠使用数量
	Return_code string `xml:"coupon_count"`
	// 代金券或立减优惠ID	coupon_id_$n	否	String(20)	10000	代金券或立减优惠ID,$n为下标，从0开始编号
	Coupon_id_1 string `xml:"coupon_id_1"`
	// 单个代金券或立减优惠支付金额	coupon_fee_$n	否	Int	100	单个代金券或立减优惠支付金额,$n为下标，从0开始编号
	Coupon_fee_1 string `xml:"coupon_fee_1"`
	// 微信支付订单号	transaction_id	是	String(32)	1217752501201407033233368018	微信支付订单号
	Transaction_id string `xml:"transaction_id"`
	// 商户订单号	out_trade_no	是	String(32)	1212321211201407033568112322	商户系统的订单号，与请求一致。
	Out_trade_no string `xml:"out_trade_no"`
	// 商家数据包	attach	否	String(128)	123456	商家数据包，原样返回
	Attach string `xml:"attach"`
	// 支付完成时间	time_end	是	String(14)	20141030133525	支付完成时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。其他详见时间规则
	Time_end string `xml:"time_end"`
}

func NewNaviteNotify(b []byte, wxPayConfig *WxPayConfig) (*NaviteNotify, bool, error) {
	n := &NaviteNotify{}
	isSignValid := n.CheckSign(wxPayConfig)
	return n, isSignValid, nil
}

type NotyfyCallback struct {
	Return_code string
	Return_msg  string
}

func (n *NotyfyCallback) ToXML() string {
	return fmt.Sprintf(
		"<xml><return_code>%v</return_code><return_msg>%v</return_msg></xml>",
		n.Return_code,
		n.Return_msg)
}

func (n *NaviteNotify) CheckSign(wxPayConfig *WxPayConfig) bool {

	p := Kvpairs{}
	ref := reflect.ValueOf(n).Elem()
	typ := reflect.TypeOf(n).Elem()
	nf := typ.NumField()
	for i := 0; i < nf; i++ {
		f := typ.Field(i)
		name := f.Tag.Get("xml")
		v := fmt.Sprintf("%v", ref.Field(i))
		// log.Println(name)
		// log.Println(f.Name)
		// log.Printf("%v = %v", name, v)
		if v != "" && name != "sign" {
			p = append(p, Kvpair{name, v})
		}

	}

	paraFilter(&p)
	argSort(&p)

	sign := unifiedorderSign(&p, wxPayConfig)
	if n.Sign == sign {
		return true
	}
	log.Printf("验证签名错误 获取签名 %v , 计算签名 %v \n", n.Sign, sign)
	return false
}

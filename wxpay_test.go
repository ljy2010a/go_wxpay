package go_wxpay

import (
	"testing"
)

func TestUnifiedorders(t *testing.T) {
	// https://pay.weixin.qq.com/wiki/tools/signverify/
	// 验证地址
	t.Log("TestNewUnifiedorder")
	u := NewUnifiedorder(GWxPayConfig)
	// 随机字符串
	u.Nonce_str = md5String(NewOrderNo())
	// 商品描述
	u.Body = "xxxxxxxxx"
	// 商户订单号
	u.Out_trade_no = NewOrderNo()
	// 总金额
	u.Total_fee = 1
	// 终端IP
	u.Spbill_create_ip = "14.23.162.172"
	// 通知地址 Config 统一配置
	// u.Notify_url = "xxxxxxxxx"
	// 交易类型
	u.Trade_type = "NATIVE"
	// 商品ID
	u.Product_id = "xxxxxxxxx"

	//====== 选填
	// 设备号
	u.Device_info = "xxxxxxxxx"
	// 商品详情
	u.Detail = "xxxxxxxxx"
	// 附加数据
	u.Attach = "xxxxxxxxx"
	// 货币类型
	u.Fee_type = "xxxxxxxxx"
	// 交易起始时间
	// u.Time_start = "xxxxxxxxx"
	// 交易结束时间
	// u.Time_expire = "xxxxxxxxx"
	// 商品标记
	u.Goods_tag = "xxxxxxxxx"
	// 指定支付方式
	u.Limit_pay = "xxxxx"
	// 用户标识
	// u.Openid = "xxxxxxxxx"

	uresp, err := u.TakeOrder(GWxPayConfig)
	if err != nil {
		t.Fatalf("TakeOrder fail %v", err)
	}
	t.Logf("uresp.Return_code %v \n", uresp.Return_code)
}

func TestUnifiedorderResponse(t *testing.T) {
	t.Log("TestUnifiedorderResponse")

	tstring := []byte(`<xml>
   <return_code>SUCCESS<![CDATA[]]></return_code>
   <return_msg><![CDATA[OK]]></return_msg>
   <appid><![CDATA[wx2421b1c4370ec43b]]></appid>
   <mch_id><![CDATA[10000100]]></mch_id>
   <nonce_str><![CDATA[IITRi8Iabbblz1Jc]]></nonce_str>
   <sign><![CDATA[7921E432F65EB8ED0CE9755F0E86D72F]]></sign>
   <result_code><![CDATA[SUCCESS]]></result_code>
   <prepay_id><![CDATA[wx201411101639507cbf6ffd8b0779950874]]></prepay_id>
   <trade_type><![CDATA[JSAPI]]></trade_type>
</xml>`)

	_, err := ParseUResponse(tstring)
	if err != nil {
		t.Fatalf("ParseUnifiedorderResponse err %v", err)
	}

}

func TestUnifiedorder(t *testing.T) {
	t.Log("TestUnifiedorderSign")
	u := &Unifiedorder{}
	u.Appid = "xxx"
	u.Mch_id = "xxx"
	u.Nonce_str = "xzz"
	u.Total_fee = 1234
	xml := u.SetSign(GWxPayConfig)
	t.Logf("xml : %v", xml)

	uresp := &UnifiedorderResp{}
	uresp, err := u.TakeOrder(GWxPayConfig)
	if err != nil {
		t.Fatalf("TakeOrder fail %v", err)
	}

	t.Logf("uresp.Return_code %v \n", uresp.Return_code)
}

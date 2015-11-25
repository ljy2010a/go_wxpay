package wxpay

import (
	"testing"
)

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
	xml := u.SetSign()
	t.Logf("xml : %v", xml)

	uresp := &UnifiedorderResp{}
	uresp, err := u.TakeOrder()
	if err != nil {
		t.Fatalf("TakeOrder fail %v", err)
	}

	t.Logf("uresp.Return_code %v \n", uresp.Return_code)
}

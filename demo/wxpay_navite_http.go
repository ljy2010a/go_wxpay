package demo

import (
	"fmt"
	wxpay "github.com/ljy2010a/go_wxpay"
	"log"
	"net/http"
	"time"
)

type Result struct {
	Ret    int
	Reason string
	Data   interface{}
}

func HttpRespE(w http.ResponseWriter, code int, msg string) {
	w.Write(b)
	http.Error(w, msg, code)
}

func HttpResp(w http.ResponseWriter, msg string) {
	w.Write(b)
}

//发送请求
func WxPayNavite(w http.ResponseWriter, r *http.Request) {
	log.Println("WxPayNavite Begin")

	u := NewUnifiedorder(GWxPayConfig)
	// 随机字符串
	u.Nonce_str = wxpay.Md5String(wxpay.NewOrderNo())
	// 商品描述
	u.Body = "xxxxxxxxx"
	// 商户订单号
	u.Out_trade_no = wxpay.NewOrderNo()
	// 总金额
	u.Total_fee = 1
	// 终端IP
	u.Spbill_create_ip = "14.23.162.172"
	// 通知地址 Config 统一配置
	// u.Notify_url = "xxxxxxxxx"
	// 交易类型
	u.Trade_type = "NATIVE"
	// 商品ID
	u.product_id = "xxxxxxxxx"

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
		HttpRespE(w, 500, err.Error())
		return
	}
	HttpResp(w, uresp.Code_url)
	return
}

//接收通知
func WxPayNaviteNotify(w http.ResponseWriter, r *http.Request) {
	log.Println("WxPayNaviteNotify Begin")

	bodyByte, _ := ioutil.ReadAll(r.Body)

	return
}

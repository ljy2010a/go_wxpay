package demo

import (
	wxpay "github.com/ljy2010a/go_wxpay"
)

func initWxConfig() {
	wxpay.GWxPayConfig.Notify_url = "www.baidu.com"
}

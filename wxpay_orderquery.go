package go_wxpay

const (
	OrderQueryApi = "https://api.mch.weixin.qq.com/pay/orderquery"
)

type OrderQuery struct {
	// 	公众账号ID	appid	是	String(32)	wxd678efh567hg6787	微信分配的公众账号ID（企业号corpid即为此appId）
	Appid string `xml:"appid"`
	// 商户号	mch_id	是	String(32)	1230000109	微信支付分配的商户号
	Mch_id string `xml:"mch_id"`
	// 微信订单号	transaction_id	二选一	String(32)	1009660380201506130728806387	微信的订单号，优先使用
	Transaction_id string `xml:"transaction_id"`
	// 商户订单号	out_trade_no	String(32)	20150806125346	商户系统内部的订单号，当没提供transaction_id时需要传这个。
	Out_trade_no string `xml:"out_trade_no"`
	// 随机字符串	nonce_str	是	String(32)	C380BEC2BFD727A4B6845133519F3AD6	随机字符串，不长于32位。推荐随机数生成算法
	Nonce_str string `xml:"nonce_str"`
	// 签名	sign	是	String(32)	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	签名，详见签名生成算法
	Sign string `xml:"sign"`
}

type OrderQueryBack struct {
	// 	返回状态码	return_code	是	String(16)	SUCCESS
	// SUCCESS/FAIL
	// 此字段是通信标识，非交易标识，交易是否成功需要查看trade_state来判断
	Return_code string `xml:"return_code"`
	// 返回信息	return_msg	否	String(128)	签名失败
	// 返回信息，如非空，为错误原因
	// 签名失败
	// 参数格式校验错误
	Return_msg string `xml:"return_msg"`
	// 公众账号ID	appid	是	String(32)	wxd678efh567hg6787	微信分配的公众账号ID
	Appid string `xml:"appid"`
	// 商户号	mch_id	是	String(32)	1230000109	微信支付分配的商户号
	Mch_id string `xml:"mch_id"`
	// 随机字符串	nonce_str	是	String(32)	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	随机字符串，不长于32位。推荐随机数生成算法
	Nonce_str string `xml:"nonce_str"`
	// 签名	sign	是	String(32)	C380BEC2BFD727A4B6845133519F3AD6	签名，详见签名生成算法
	Sign string `xml:"sign"`
	// 业务结果	result_code	是	String(16)	SUCCESS	SUCCESS/FAIL
	Result_code string `xml:"result_code"`
	// 错误代码	err_code	否	String(32)	SYSTEMERROR	详细参见第6节错误列表
	Err_code string `xml:"err_code"`
	// 错误代码描述	err_code_des	否	String(128)	系统错误	结果信息描述
	Err_code_des string `xml:"err_code_des"`
	// 设备号	device_info	否	String(32)	013467007045764	微信支付分配的终端设备号，
	Device_info string `xml:"device_info"`
	// 用户标识	openid	是	String(128)	oUpF8uMuAJO_M2pxb1Q9zNjWeS6o	用户在商户appid下的唯一标识
	Openid string `xml:"openid"`
	// 是否关注公众账号	is_subscribe	否	String(1)	Y	用户是否关注公众账号，Y-关注，N-未关注，仅在公众账号类型支付有效
	Is_subscribe string `xml:"is_subscribe"`
	// 交易类型	trade_type	是	String(16)	JSAPI	调用接口提交的交易类型，取值如下：JSAPI，NATIVE，APP，MICROPAY，详细说明见参数规定
	Trade_type string `xml:"trade_type"`
	// 交易状态	trade_state	是	String(32)	SUCCESS
	// SUCCESS—支付成功
	// REFUND—转入退款
	// NOTPAY—未支付
	// CLOSED—已关闭
	// REVOKED—已撤销（刷卡支付）
	// USERPAYING--用户支付中
	// PAYERROR--支付失败(其他原因，如银行返回失败)
	err_code string `xml:"err_code"`
	// 付款银行	bank_type	是	String(16)	CMC	银行类型，采用字符串类型的银行标识
	Bank_type string `xml:"bank_type"`
	// 总金额	total_fee	是	Int	100	订单总金额，单位为分
	Total_fee string `xml:"total_fee"`
	// 货币种类	fee_type	否	String(8)	CNY	货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	Fee_type string `xml:"fee_type"`
	// 现金支付金额	cash_fee	是	Int	100	现金支付金额订单现金支付金额，详见支付金额
	Cash_fee string `xml:"cash_fee"`
	// 现金支付货币类型	cash_fee_type	否	String(16)	CNY	货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	Cash_fee_type string `xml:"cash_fee_type"`
	// 代金券或立减优惠金额	coupon_fee	否	Int	100	“代金券或立减优惠”金额<=订单总金额，订单总金额-“代金券或立减优惠”金额=现金支付金额，详见支付金额
	Coupon_fee string `xml:"coupon_fee"`
	// 代金券或立减优惠使用数量	coupon_count	否	Int	1	代金券或立减优惠使用数量
	Coupon_count string `xml:"coupon_count"`
	// 代金券或立减优惠批次ID	coupon_batch_id_$n	否	String(20)	100	代金券或立减优惠批次ID ,$n为下标，从0开始编号
	Coupon_batch_id_1 string `xml:"coupon_batch_id_1"`
	// 代金券或立减优惠ID	coupon_id_$n	否	String(20)	10000 	代金券或立减优惠ID, $n为下标，从0开始编号
	Coupon_id_1 string `xml:"coupon_id_"`
	// 单个代金券或立减优惠支付金额	coupon_fee_$n	否	Int	100	单个代金券或立减优惠支付金额, $n为下标，从0开始编号
	Coupon_fee_1 string `xml:"coupon_fee_1"`
	// 微信支付订单号	transaction_id	是	String(32)	1009660380201506130728806387	微信支付订单号
	Transaction_id string `xml:"transaction_id"`
	// 商户订单号	out_trade_no	是	String(32)	20150806125346	商户系统的订单号，与请求一致。
	Out_trade_no string `xml:"out_trade_no"`
	// 附加数据	attach	否	String(128)	深圳分店	附加数据，原样返回
	Attach string `xml:"attach"`
	// 支付完成时间	time_end	是	String(14)	20141030133525	订单支付时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。其他详见时间规则
	Time_end string `xml:"time_end"`
	// 交易状态描述	trade_state_desc	是	String(256)	支付失败，请重新下单支付	对当前查询订单状态的描述和下一步操作的指引
	Trade_state_desc string `xml:"trade_state_desc"`
}

func (o *OrderQuery) Query() (*OrderQueryBack, error) {

	return nil, nil
}

func (o *OrderQuery) CheckSign() bool {
	return false
}

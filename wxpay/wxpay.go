package wxpay

import "encoding/xml"

const (
	// Success status success
	Success = "SUCCESS"
	// Fail status fail
	Fail = "FAIL"
)

type (
	// Sign sign
	Sign struct {
		XMLName   xml.Name `xml:"xml" url:"-"`                                     // 是 公共头部
		AppID     string   `xml:"appid,omitempty" url:"appid,omitempty"`           // 是 公众账号ID
		MchID     string   `xml:"mch_id,omitempty" url:"mch_id,omitempty"`         // 是 商户号
		NonceStr  string   `xml:"nonce_str,omitempty" url:"nonce_str,omitempty"`   // 是 随机字符串
		Sign      string   `xml:"sign,omitempty" url:"-"`                          // 是 签名
		SignType  string   `xml:"sign_type,omitempty" url:"sign_type,omitempty"`   // 否 签名类型
		NotifyURL string   `xml:"notify_url,omitempty" url:"notify_url,omitempty"` // 是 通知地址
		Request
	}
	// Request request
	Request struct {
		TransactionID string `xml:"transaction_id,omitempty" url:"transaction_id,omitempty"`                                   // 是 微信支付订单号
		OutTradeNo    string `xml:"out_trade_no,omitempty" url:"out_trade_no,omitempty"`                                       // 是 商户订单号
		OutRefundNo   string `xml:"out_refund_no,omitempty" url:"out_refund_no,omitempty"`                                     // 是 商户退款单号
		TotalFee      int    `xml:"total_fee,omitempty" url:"total_fee,omitempty" json:"total_fee,omitempty"`                  // 是 总金额
		RefundFee     int    `xml:"refund_fee,omitempty" url:"refund_fee,omitempty" json:"refund_fee,omitempty"`               // 是 退款金额
		RefundFeeType string `xml:"refund_fee_type,omitempty" url:"refund_fee_type,omitempty" json:"efund_fee_type,omitempty"` // 否 货币类型
		RefundDesc    string `xml:"refund_desc,omitempty" url:"refund_desc,omitempty" json:"refund_desc,omitempty"`            // 否 退款原因
		RefundAccount string `xml:"refund_account,omitempty" url:"refund_account,omitempty"`                                   // 否 资金来源
		Offset        int    `xml:"offset,omitempty" url:"offset,omitempty"`                                                   // 否 偏移量
	}
	// Response response
	Response struct {
		XMLName    xml.Name `xml:"xml" url:"-"`                                         // 是 公共头部
		ReturnCode string   `xml:"return_code,omitempty" url:"return_code,omitempty"`   // 是 返回状态码
		ReturnMsg  string   `xml:"return_msg,omitempty" url:"return_msg,omitempty"`     // 否 返回信息
		ResultCode string   `xml:"result_code,omitempty" url:"result_code,omitempty"`   // 是 业务结果
		ErrCode    string   `xml:"err_code,omitempty" url:"err_code,omitempty"`         // 否 错误代码
		ErrCodeDes string   `xml:"err_code_des,omitempty" url:"err_code_des,omitempty"` // 否 错误代码
	}
	// Result result
	Result struct {
		Response
		AppID               string `xml:"appid,omitempty" url:"appid,omitempty"`                                 // 是 公众账号ID
		MchID               string `xml:"mch_id,omitempty" url:"mch_id,omitempty"`                               // 是 商户号
		NonceStr            string `xml:"nonce_str,omitempty" url:"nonce_str,omitempty"`                         // 是 随机字符串
		Sign                string `xml:"sign,omitempty" url:"-"`                                                // 是 签名
		SignType            string `xml:"sign_type,omitempty" url:"sign_type,omitempty"`                         // 否 签名类型
		TransactionID       string `xml:"transaction_id,omitempty" url:"transaction_id,omitempty"`               // 是 微信支付订单号
		OutTradeNo          string `xml:"out_trade_no,omitempty" url:"out_trade_no,omitempty"`                   // 是 商户订单号
		OutRefundNo         string `xml:"out_refund_no,omitempty" url:"out_refund_no,omitempty"`                 // 是 商户退款单号
		RefundID            string `xml:"refund_id,omitempty" url:"refund_id,omitempty"`                         // 是 微信退款订单号
		RefundFee           int    `xml:"refund_fee,omitempty" url:"refund_fee,omitempty"`                       // 是 退款金额
		SettlementRefundFee int    `xml:"settlement_refund_fee,omitempty" url:"settlement_refund_fee,omitempty"` // 否 应结退款金额
		TotalFee            int    `xml:"total_fee,omitempty" url:"total_fee,omitempty"`                         // 是 标价金额
		SettlementTotalFee  int    `xml:"settlement_total_fee,omitempty" url:"settlement_total_fee,omitempty"`   // 是 应结订单金额
		FeeType             string `xml:"fee_type,omitempty" url:"fee_type,omitempty"`                           // 否 标价币种
		CashFee             int    `xml:"cash_fee,omitempty" url:"cash_fee,omitempty"`                           // 是 现金支付金额
		CashFeeType         string `xml:"cash_fee_type,omitempty" url:"cash_fee_type,omitempty"`                 // 否 现金支付币种
		CashRefundFee       int    `xml:"cash_refund_fee,omitempty" url:"cash_refund_fee,omitempty"`             // 否 现金退款金额
		CouponType0         string `xml:"coupon_type_0,omitempty" url:"coupon_type_0,omitempty"`                 // 否 代金券类型
		CouponType1         string `xml:"coupon_type_1,omitempty" url:"coupon_type_1,omitempty"`                 // 否 代金券类型
		CouponType2         string `xml:"coupon_type_2,omitempty" url:"coupon_type_2,omitempty"`                 // 否 代金券类型
		CouponType3         string `xml:"coupon_type_3,omitempty" url:"coupon_type_3,omitempty"`                 // 否 代金券类型
		CouponType4         string `xml:"coupon_type_4,omitempty" url:"coupon_type_4,omitempty"`                 // 否 代金券类型
		CouponRefundFee     int    `xml:"coupon_refund_fee,omitempty" url:"coupon_refund_fee,omitempty"`         // 否 代金券退款总金额
		CouponRefundFee0    int    `xml:"coupon_refund_fee_0,omitempty" url:"coupon_refund_fee_0,omitempty"`     // 否 单个代金券退款金额
		CouponRefundFee1    int    `xml:"coupon_refund_fee_1,omitempty" url:"coupon_refund_fee_1,omitempty"`     // 否 单个代金券退款金额
		CouponRefundFee2    int    `xml:"coupon_refund_fee_2,omitempty" url:"coupon_refund_fee_2,omitempty"`     // 否 单个代金券退款金额
		CouponRefundFee3    int    `xml:"coupon_refund_fee_3,omitempty" url:"coupon_refund_fee_3,omitempty"`     // 否 单个代金券退款金额
		CouponRefundFee4    int    `xml:"coupon_refund_fee_4,omitempty" url:"coupon_refund_fee_4,omitempty"`     // 否 单个代金券退款金额
		CouponRefundCount   int    `xml:"coupon_refund_count,omitempty" url:"coupon_refund_count,omitempty"`     // 否 退款代金券使用数量
		CouponRefundID0     string `xml:"coupon_refund_id_0,omitempty" url:"coupon_refund_id_0,omitempty"`       // 否 退款代金券ID
		CouponRefundID1     string `xml:"coupon_refund_id_1,omitempty" url:"coupon_refund_id_1,omitempty"`       // 否 退款代金券ID
		CouponRefundID2     string `xml:"coupon_refund_id_2,omitempty" url:"coupon_refund_id_2,omitempty"`       // 否 退款代金券ID
		CouponRefundID3     string `xml:"coupon_refund_id_3,omitempty" url:"coupon_refund_id_3,omitempty"`       // 否 退款代金券ID
		CouponRefundID4     string `xml:"coupon_refund_id_4,omitempty" url:"coupon_refund_id_4,omitempty"`       // 否 退款代金券ID

	}
	// Notice notice
	Notice struct {
		Response
		AppID    string `xml:"appid,omitempty" url:"appid,omitempty"`         // 是 公众账号ID
		MchID    string `xml:"mch_id,omitempty" url:"mch_id,omitempty"`       // 是 商户号
		NonceStr string `xml:"nonce_str,omitempty" url:"nonce_str,omitempty"` // 是 随机字符串
		ReqInfo  string `xml:"req_info,omitempty" url:"req_info,omitempty"`   // 是 加密信息
		Ciphertext
	}
	// Ciphertext ciphertext
	Ciphertext struct {
		TransactionID       string `xml:"transaction_id,omitempty" url:"transaction_id,omitempty"`               // 是 微信支付订单号
		OutTradeNo          string `xml:"out_trade_no,omitempty" url:"out_trade_no,omitempty"`                   // 是 商户订单号
		OutRefundNo         string `xml:"out_refund_no,omitempty" url:"out_refund_no,omitempty"`                 // 是 商户退款单号
		RefundID            string `xml:"refund_id,omitempty" url:"refund_id,omitempty"`                         // 是 微信退款订单号
		RefundFee           int    `xml:"refund_fee,omitempty" url:"refund_fee,omitempty"`                       // 是 退款金额
		SettlementRefundFee int    `xml:"settlement_refund_fee,omitempty" url:"settlement_refund_fee,omitempty"` // 否 应结退款金额
		TotalFee            int    `xml:"total_fee,omitempty" url:"total_fee,omitempty"`                         // 是 标价金额
		SettlementTotalFee  int    `xml:"settlement_total_fee,omitempty" url:"settlement_total_fee,omitempty"`   // 是 应结订单金额
		RefundStatus        string `xml:"refund_status,omitempty" url:"refund_status,omitempty"`                 // 是 退款状态
		SuccessTime         string `xml:"success_time,omitempty" url:"success_time,omitempty" `                  // 否 退款成功时间
		RefundRecvAccout    string `xml:"refund_recv_accout,omitempty" url:"refund_recv_accout,omitempty"`       // 是 退款入账账户
		RefundAccount       string `xml:"refund_account,omitempty" url:"refund_account,omitempty"`               // 是 退款资金来源
		RefundRequestSource string `xml:"refund_request_source,omitempty" url:"refund_request_source,omitempty"` // 是 退款发起来源
	}
	// Query query
	Query struct {
		Response
		AppID                string `xml:"appid,omitempty" url:"appid,omitempty"`                                     // 是 公众账号ID
		MchID                string `xml:"mch_id,omitempty" url:"mch_id,omitempty"`                                   // 是 商户号
		NonceStr             string `xml:"nonce_str,omitempty" url:"nonce_str,omitempty"`                             // 是 随机字符串
		Sign                 string `xml:"sign,omitempty" url:"-"`                                                    // 是 签名
		SignType             string `xml:"sign_type,omitempty" url:"sign_type,omitempty"`                             // 否 签名类型
		TotalRefundCount     int    `xml:"total_refund_count,omitempty" url:"total_refund_countomitempty"`            // 否 退款次数
		TransactionID        string `xml:"transaction_id,omitempty" url:"transaction_id,omitempty"`                   // 是 微信支付订单号
		OutTradeNo           string `xml:"out_trade_no,omitempty" url:"out_trade_no,omitempty"`                       // 是 商户订单号
		TotalFee             int    `xml:"total_fee,omitempty" url:"total_fee,omitempty"`                             // 是 标价金额
		SettlementTotalFee   int    `xml:"settlement_total_fee,omitempty" url:"settlement_total_fee,omitempty"`       // 是 应结订单金额
		FeeType              string `xml:"fee_type,omitempty" url:"fee_type,omitempty"`                               // 否 标价币种
		CashFee              int    `xml:"cash_fee,omitempty" url:"cash_fee,omitempty"`                               // 是 现金支付金额
		RefundCount          int    `xml:"refund_count,omitempty" url:"refund_count,omitempty"`                       // 是 退款笔数
		OutRefundNo0         string `xml:"out_refund_no_0,omitempty" url:"out_refund_no_0,omitempty"`                 // 是 商户退款单号
		OutRefundNo1         string `xml:"out_refund_no_1,omitempty" url:"out_refund_no_1,omitempty"`                 // 是 商户退款单号
		OutRefundNo2         string `xml:"out_refund_no_2,omitempty" url:"out_refund_no_2,omitempty"`                 // 是 商户退款单号
		OutRefundNo3         string `xml:"out_refund_no_3,omitempty" url:"out_refund_no_3,omitempty"`                 // 是 商户退款单号
		OutRefundNo4         string `xml:"out_refund_no_4,omitempty" url:"out_refund_no_4,omitempty"`                 // 是 商户退款单号
		RefundID0            string `xml:"refund_id_0,omitempty" url:"refund_id_0,omitempty"`                         // 是 微信退款单号
		RefundID1            string `xml:"refund_id_1,omitempty" url:"refund_id_1,omitempty"`                         // 是 微信退款单号
		RefundID2            string `xml:"refund_id_2,omitempty" url:"refund_id_2,omitempty"`                         // 是 微信退款单号
		RefundID3            string `xml:"refund_id_3,omitempty" url:"refund_id_3,omitempty"`                         // 是 微信退款单号
		RefundID4            string `xml:"refund_id_4,omitempty" url:"refund_id_4,omitempty"`                         // 是 微信退款单号
		RefundChannel0       string `xml:"refund_channel_0,omitempty" url:"refund_channel_0,omitempty"`               // 否 退款渠道
		RefundChannel1       string `xml:"refund_channel_1,omitempty" url:"refund_channel_1,omitempty"`               // 否 退款渠道
		RefundChannel2       string `xml:"refund_channel_2,omitempty" url:"refund_channel_2,omitempty"`               // 否 退款渠道
		RefundChannel3       string `xml:"refund_channel_3,omitempty" url:"refund_channel_3,omitempty"`               // 否 退款渠道
		RefundChannel4       string `xml:"refund_channel_4,omitempty" url:"refund_channel_4,omitempty"`               // 否 退款渠道
		RefundFee0           int    `xml:"refund_fee_0,omitempty" url:"refund_fee_0,omitempty"`                       // 是 申请退款金额
		RefundFee1           int    `xml:"refund_fee_1,omitempty" url:"refund_fee_1,omitempty"`                       // 是 申请退款金额
		RefundFee2           int    `xml:"refund_fee_2,omitempty" url:"refund_fee_2,omitempty"`                       // 是 申请退款金额
		RefundFee3           int    `xml:"refund_fee_3,omitempty" url:"refund_fee_3,omitempty"`                       // 是 申请退款金额
		RefundFee4           int    `xml:"refund_fee_4,omitempty" url:"refund_fee_4,omitempty"`                       // 是 申请退款金额
		SettlementRefundFee0 int    `xml:"settlement_refund_fee_0,omitempty" url:"settlement_refund_fee_0,omitempty"` // 否 退款金额
		SettlementRefundFee1 int    `xml:"settlement_refund_fee_1,omitempty" url:"settlement_refund_fee_1,omitempty"` // 否 退款金额
		SettlementRefundFee2 int    `xml:"settlement_refund_fee_2,omitempty" url:"settlement_refund_fee_2,omitempty"` // 否 退款金额
		SettlementRefundFee3 int    `xml:"settlement_refund_fee_3,omitempty" url:"settlement_refund_fee_3,omitempty"` // 否 退款金额
		SettlementRefundFee4 int    `xml:"settlement_refund_fee_4,omitempty" url:"settlement_refund_fee_4,omitempty"` // 否 退款金额
		CouponTypeNM         string `xml:"coupon_type_$n_$m,omitempty" url:"coupon_type_$n_$m,omitempty"`             // 否 代金券类型
		CouponRefundFee0     int    `xml:"coupon_refund_fee_0,omitempty" url:"coupon_refund_fee_0,omitempty"`         // 否 总个代金券退款金额
		CouponRefundFee1     int    `xml:"coupon_refund_fee_1,omitempty" url:"coupon_refund_fee_1,omitempty"`         // 否 总个代金券退款金额
		CouponRefundFee2     int    `xml:"coupon_refund_fee_2,omitempty" url:"coupon_refund_fee_2,omitempty"`         // 否 总个代金券退款金额
		CouponRefundFee3     int    `xml:"coupon_refund_fee_3,omitempty" url:"coupon_refund_fee_3,omitempty"`         // 否 总个代金券退款金额
		CouponRefundFee4     int    `xml:"coupon_refund_fee_4,omitempty" url:"coupon_refund_fee_4,omitempty"`         // 否 总个代金券退款金额
		CouponRefundCount0   int    `xml:"coupon_refund_count_0,omitempty" url:"coupon_refund_count_0,omitempty"`     // 否 单个代金券退款金额
		CouponRefundCount1   int    `xml:"coupon_refund_count_1,omitempty" url:"coupon_refund_count_1,omitempty"`     // 否 单个代金券退款金额
		CouponRefundCount2   int    `xml:"coupon_refund_count_2,omitempty" url:"coupon_refund_count_2,omitempty"`     // 否 单个代金券退款金额
		CouponRefundCount3   int    `xml:"coupon_refund_count_3,omitempty" url:"coupon_refund_count_3,omitempty"`     // 否 单个代金券退款金额
		CouponRefundCount4   int    `xml:"coupon_refund_count_4,omitempty" url:"coupon_refund_count_4,omitempty"`     // 否 单个代金券退款金额
		CouponRefundIDNM     string `xml:"coupon_refund_id_$n_$m,omitempty" url:"coupon_refund_id_$n_$m,omitempty"`   // 否 退款代金券ID
		CouponRefundFeeNM    int    `xml:"coupon_refund_fee_$n_$m,omitempty" url:"coupon_refund_fee_$n_$m,omitempty"` // 否 单个代金券退款金额
		RefundStatus0        string `xml:"refund_status_0,omitempty" url:"refund_status_0,omitempty"`                 // 是 退款状态
		RefundStatus1        string `xml:"refund_status_1,omitempty" url:"refund_status_1,omitempty"`                 // 是 退款状态
		RefundStatus2        string `xml:"refund_status_2,omitempty" url:"refund_status_2,omitempty"`                 // 是 退款状态
		RefundStatus3        string `xml:"refund_status_3,omitempty" url:"refund_status_3,omitempty"`                 // 是 退款状态
		RefundStatus4        string `xml:"refund_status_4,omitempty" url:"refund_status_4,omitempty"`                 // 是 退款状态
		RefundAccount0       string `xml:"refund_account_0,omitempty" url:"refund_account_0,omitempty"`               // 否 退款资金来源
		RefundAccount1       string `xml:"refund_account_1,omitempty" url:"refund_account_1,omitempty"`               // 否 退款资金来源
		RefundAccount2       string `xml:"refund_account_2,omitempty" url:"refund_account_2,omitempty"`               // 否 退款资金来源
		RefundAccount3       string `xml:"refund_account_3,omitempty" url:"refund_account_3,omitempty"`               // 否 退款资金来源
		RefundAccount4       string `xml:"refund_account_4,omitempty" url:"refund_account_4,omitempty"`               // 否 退款资金来源
		RefundRecvAccout0    string `xml:"refund_recv_accout_0,omitempty" url:"refund_recv_accout_0,omitempty"`       // 是 退款入账账户
		RefundRecvAccout1    string `xml:"refund_recv_accout_1,omitempty" url:"refund_recv_accout_1,omitempty"`       // 是 退款入账账户
		RefundRecvAccout2    string `xml:"refund_recv_accout_2,omitempty" url:"refund_recv_accout_2,omitempty"`       // 是 退款入账账户
		RefundRecvAccout3    string `xml:"refund_recv_accout_3,omitempty" url:"refund_recv_accout_3,omitempty"`       // 是 退款入账账户
		RefundRecvAccout4    string `xml:"refund_recv_accout_4,omitempty" url:"refund_recv_accout_4,omitempty"`       // 是 退款入账账户
		RefundSuccessTime0   string `xml:"refund_success_time_0,omitempty" url:"refund_success_time_0,omitempty"`     // 否 退款成功时间
		RefundSuccessTime1   string `xml:"refund_success_time_1,omitempty" url:"refund_success_time_1,omitempty"`     // 否 退款成功时间
		RefundSuccessTime2   string `xml:"refund_success_time_2,omitempty" url:"refund_success_time_2,omitempty"`     // 否 退款成功时间
		RefundSuccessTime3   string `xml:"refund_success_time_3,omitempty" url:"refund_success_time_3,omitempty"`     // 否 退款成功时间
		RefundSuccessTime4   string `xml:"refund_success_time_4,omitempty" url:"refund_success_time_4,omitempty"`     // 否 退款成功时间
	}
)

package alipay

type (
	// Sign sign
	Sign struct {
		AppID        string `json:"app_id,omitempty" url:"app_id,omitempty"`                 // 是 应用ID
		Method       string `json:"method,omitempty" url:"method,omitempty"`                 // 是 接口名称
		Format       string `json:"format,omitempty" url:"format,omitempty"`                 // 否 JSON
		ReturnURL    string `json:"return_url,omitempty" url:"return_url,omitempty"`         // 否 跳转地址
		Charset      string `json:"charset,omitempty" url:"charset,omitempty"`               // 是 utf-8
		SignType     string `json:"sign_type,omitempty" url:"sign_type,omitempty"`           // 是 RSA2
		Sign         string `json:"sign,omitempty" url:"sign,omitempty"`                     // 是 签名
		TimeStamp    string `json:"timestamp,omitempty" url:"timestamp,omitempty"`           // 是 时间
		Version      string `json:"version,omitempty" url:"version,omitempty"`               // 是 1.0
		AppAuthToken string `json:"app_auth_token,omitempty" url:"app_auth_token,omitempty"` // 否 详见应用授权概述
		NotifyURL    string `json:"notify_url,omitempty" url:"notify_url,omitempty"`         // 否 通知地址
		BizContent   string `json:"biz_content,omitempty" url:"biz_content,omitempty"`       // 是 非公共参数
	}
	// Request request
	Request struct {
		OutTradeNo              string  `json:"out_trade_no,omitempty" url:"out_trade_no,omitempty"`                           // 是 单号
		TradeNo                 string  `json:"trade_no,omitempty" url:"trade_no,omitempty"`                                   // 否 支付宝交易号
		RefundAmount            float64 `json:"refund_amount,omitempty" url:"refund_amount,omitempty"`                         // 是 退款金额
		RefundCurrency          float64 `json:"refund_currency,omitempty" url:"refund_currency,omitempty"`                     // 是 退款币种
		RefundReason            string  `json:"refund_reason,omitempty" url:"refund_reason,omitempty"`                         // 否 退款说明
		OutRequestNo            string  `json:"out_request_id,omitempty" url:"out_request_id,omitempty"`                       // 否 标识一次退款请求
		OperatorID              string  `json:"operator_id,omitempty" url:"operator_id,omitempty"`                             // 否 操作员编号
		TerminalID              string  `json:"terminal_id,omitempty" url:"terminal_id,omitempty"`                             // 否 终端编号
		GoodsDetail             string  `json:"goods_detail,omitempty" url:"goods_detail,omitempty"`                           // 否 商品明细
		RefundRoyaltyParameters string  `json:"refund_royalty_parameters,omitempty" url:"refund_royalty_parameters,omitempty"` // 否 退账明细
		OrgPid                  string  `json:"org_pid,omitempty" url:"org_pid,omitempty"`                                     // 否 机构编号
	}
	// Result result
	Result struct {
		AlipayTradeFefundResponse            *Notice `json:"alipay_trade_refund_response,omitempty"`               // 内容
		AlipayTradeFastpayRfundQueryResponse *Query  `json:"alipay_trade_fastpay_refund_query_response,omitempty"` // 内容
		Sign                                 string  `json:"sign,omitempty"`                                       // 签名
	}
	// Response response
	Response struct {
		Code    string `json:"code,omitempty"`     // 网关返回码
		Msg     string `json:"msg,omitempty"`      // 网关返回码描述
		SubCode string `json:"sub_code,omitempty"` // 业务返回码
		SumbMsg string `json:"sub_msg,omitempty"`  // 业务描述
	}
	// Notice notice
	Notice struct {
		Response
		TradeNo                      string  `json:"trade_no,omitempty" url:"trade_no,omitempty"`                                               // 是 交易号
		OutTradeNo                   string  `json:"out_trade_no,omitempty" url:"out_trade_no,omitempty"`                                       // 是 单号
		BuyerLoginID                 string  `json:"buyer_login_id,omitempty" url:"buyer_login_id,omitempty"`                                   // 是 买家帐号
		FundChange                   string  `json:"fund_change,omitempty" url:"fund_change,omitempty"`                                         // 是 资金编号
		RefundFee                    float64 `json:"refund_fee,omitempty" url:"refund_fee,omitempty"`                                           // 是 退款金额
		RefundCurrency               string  `json:"refund_currency,omitempty" url:"refund_currency,omitempty"`                                 // 否 退款币种
		GmtRefundPay                 string  `json:"gmt_refund_pay,omitempty" url:"gmt_refund_pay,omitempty"`                                   // 是 退款支付时间
		RefundDetailItemList         string  `json:"refund_detail_item_list,omitempty" url:"refund_detail_item_list,omitempty"`                 // 否 退款使用的资金渠道
		StoreName                    string  `json:"store_name,omitempty" url:"store_name,omitempty"`                                           // 是 店铺名称
		BuyerUserID                  string  `json:"buyer_user_id,omitempty" url:"buyer_user_id,omitempty"`                                     // 否 卖家用户
		RefundPresetPaytoolList      string  `json:"refund_preset_paytool_list,omitempty" url:"refund_preset_paytool_list,omitempty"`           // 否 退回的前置资产列表
		RefundSettlementID           string  `json:"refund_settlement_id,omitempty" url:"refund_settlement_id,omitempty"`                       // 否 退款清算编号
		PresentRefundBuyerAmount     string  `json:"present_refund_buyer_amount,omitempty" url:"present_refund_buyer_amount,omitempty"`         // 否 本次退款金额中买家退款金额
		PresentRefundDiscountAmount  string  `json:"present_refund_discount_amount,omitempty" url:"present_refund_discount_amount,omitempty"`   // 否 本次退款金额中平台优惠退款金额
		PresentRefundMdiscountAmount string  `json:"present_refund_mdiscount_amount,omitempty" url:"present_refund_mdiscount_amount,omitempty"` // 否 本次退款金额中商家优惠退款金额
	}
	// Query query
	Query struct {
		Response
		TradeNo                      string  `json:"trade_no,omitempty" url:"trade_no,omitempty"`                                               // 是 交易号
		OutTradeNo                   string  `json:"out_trade_no,omitempty" url:"out_trade_no,omitempty"`                                       // 是 单号
		OutRequestNo                 string  `json:"out_request_no,omitempty" url:"out_request_no,omitempty"`                                   // 否 本笔退款对应的退款请求号
		RefundReason                 string  `json:"refund_reason,omitempty" url:"refund_reason,omitempty"`                                     // 否 发起退款时，传入的退款原因
		TotalAmount                  float64 `json:"total_amount,omitempty" url:"total_amount,omitempty"`                                       // 是 订单金额
		RefundAmount                 float64 `json:"refund_amount,omitempty" url:"refund_amount,omitempty"`                                     // 是 退款金额
		RefundRoyaltys               string  `json:"refund_royaltys,omitempty" url:"refund_royaltys,omitempty"`                                 // 否 退分账明细信息
		GmtRefundPay                 string  `json:"gmt_refund_pay,omitempty" url:"gmt_refund_pay,omitempty"`                                   // 否 退款时间
		RefundDetailItemList         string  `json:"refund_detail_item_list,omitempty" url:"refund_detail_item_list,omitempty"`                 // 否 本次退款使用的资金渠道；
		SendBackFee                  string  `json:"send_back_fee,omitempty" url:"send_back_fee,omitempty"`                                     // 否 本次商户实际退回金额；
		RefundSettlementID           string  `json:"refund_settlement_id,omitempty" url:"refund_settlement_id,omitempty"`                       // 否 退款清算编号
		PresentRefundBuyerAmount     string  `json:"present_refund_buyer_amount,omitempty" url:"present_refund_buyer_amount,omitempty"`         // 否 本次退款金额中买家退款金额
		PresentRefundDiscountAmount  string  `json:"present_refund_discount_amount,omitempty" url:"present_refund_discount_amount,omitempty"`   // 否 本次退款金额中平台优惠退款金额
		PresentRefundMdiscountAmount string  `json:"present_refund_mdiscount_amount,omitempty" url:"present_refund_mdiscount_amount,omitempty"` // 否 本次退款金额中商家优惠退款金额
	}
)

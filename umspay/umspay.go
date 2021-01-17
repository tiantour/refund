package umspay

const (
	// Success status success
	Success = "SUCCESS"
	// FAILED status failed
	FAILED = "FAILED"
)

var (
	// AccessToken access token
	AccessToken = ""
)

type (
	// Request request
	Request struct {
		MsgID            string   `json:"msgId,omitempty"`            // 否 消息ID
		RequestTimestamp string   `json:"requestTimestamp,omitempty"` // 是 报文请求时间 格式yyyy-MM-dd HH:mm:ss
		SrcReserve       string   `json:"srcReserve,omitempty"`       // 否 请求系统预留字段
		MerOrderID       string   `json:"merOrderId,omitempty"`       // 是 商户子订单号
		InstMID          string   `json:"instMid,omitempty"`          // 否 业务类型
		MID              string   `json:"mid,omitempty"`              // 是 商户号
		TID              string   `json:"tid,omitempty"`              // 是 机构商户号
		TargetOrderID    string   `json:"targetOrderId,omitempty"`    // 是 第三方订单号
		RefundAmount     int      `json:"refundAmount,omitempty"`     // 是 要退货的金额
		RefundOrderID    string   `json:"refundOrderId,omitempty"`    // 否 退款订单号
		PlatformAmount   int      `json:"platformAmount,omitempty"`   // 否 平台商户分账金额
		SubOrders        []*Order `json:"subOrders,omitempty"`        // 否 子订单信息
		RefundDesc       string   `json:"refundDesc,omitempty"`       // 否 退货说明
	}

	// Response response
	Response struct {
		ErrCode             string `json:"errCode,omitempty"`             // 平台错误码
		ErrMsg              string `json:"errMsg,omitempty"`              // 否 平台错误信息
		MsgID               string `json:"msgId,omitempty"`               // 否 消息ID
		RequestTimestamp    string `json:"requestTimestamp,omitempty"`    // 是 报文请求时间 格式yyyy-MM-dd HH:mm:ss
		SrcReserve          string `json:"srcReserve,omitempty"`          // 否 请求系统预留字段
		MID                 string `json:"mid,omitempty"`                 // 是 商户号
		TID                 string `json:"tid,omitempty"`                 // 是 机构商户号
		MerOrderID          string `json:"merOrderId,omitempty"`          // 否 商户订单号
		MerName             string `json:"merName,omitempty"`             // 否 商户名称
		SeqID               string `json:"seqId,omitempty"`               // 否 平台流水号
		SettleRefID         string `json:"settleRefId,omitempty"`         // 否 清分ID 字符串
		Status              string `json:"status,omitempty"`              // 否 交易状态
		TargetMID           string `json:"targetMid,omitempty"`           // 否 支付渠道商户号
		TargetOrderID       string `json:"targetOrderId,omitempty"`       // 否 第三方订单号
		TargetStatus        string `json:"targetStatus,omitempty"`        // 否 目标平台的状态
		TargetSys           string `json:"targetSys,omitempty"`           // 否 目标平台代码
		TotalAmount         int    `json:"totalAmount,omitempty"`         // 否 支付总金额
		RefundAmount        string `json:"refundAmount,omitempty"`        // 否 总退款金额
		RefundFunds         string `json:"refundFunds,omitempty"`         // 否 退款渠道列表
		RefundFundsDesc     string `json:"refundFundsDesc,omitempty"`     // 否 退款渠道描述
		RefundInvoiceAmount int    `json:"refundInvoiceAmount,omitempty"` // 否 实付部分退款金额
		RefundOrderID       string `json:"refundOrderId,omitempty"`       // 否 退货订单号
		RefundTargetOrderID string `json:"refundTargetOrderId,omitempty"` // 否 目标系统退货订单号
		YXLMAmount          int    `json:"yxlmAmount,omitempty"`          // 否 营销联盟优惠金额
		RefundStatus        string `json:"refundStatus,omitempty"`        // 否 退款状态
		BankCardNo          string `json:"bankCardNo,omitempty"`          // 否 银行卡号
		BankInfo            string `json:"bankInfo,omitempty"`            // 否 银行信息
		PayTime             string `json:"payTime,omitempty"`             // 否 支付时间 格式yyyy-MM-dd HH:mm:ss
		SettleDate          string `json:"settleDate,omitempty"`          // 否 结算日期 格式yyyy-MM-dd
		SendBackAmount      string `json:"sendBackAmount,omitempty"`      // 否 商户实退金额
	}

	// Order suborder
	Order struct {
		MID           string `json:"mid,omitempty"`           // 否 子商户号
		MerOrderID    string `json:"merOrderId,omitempty"`    // 否 商户子订单号
		RefundOrderID string `json:"refundOrderId,omitempty"` // 否 退款订单号
		TotalAmount   int    `json:"totalAmount,omitempty"`   // 否 子商户分账金额
	}

	// Query query
	Query struct {
		MsgID            string `json:"msgId,omitempty"`            // 否 消息ID
		RequestTimestamp string `json:"requestTimestamp,omitempty"` // 是 报文请求时间 格式yyyy-MM-dd HH:mm:ss
		SrcReserve       string `json:"srcReserve,omitempty"`       // 否 请求系统预留字段
		MerOrderID       string `json:"merOrderId,omitempty"`       // 是 商户子订单号
		InstMID          string `json:"instMid,omitempty"`          // 否 业务类型
		MID              string `json:"mid,omitempty"`              // 是 商户号
		TID              string `json:"tid,omitempty"`              // 是 机构商户号
	}
)

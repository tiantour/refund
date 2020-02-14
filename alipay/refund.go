package alipay

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/tiantour/fetch"
)

// Refund refund
type Refund struct{}

// NewRefund new refund
func NewRefund() *Refund {
	return &Refund{}
}

// Apply apply
func (r *Refund) Apply(str string) (*Notice, error) {
	body, err := fetch.Cmd(&fetch.Request{
		Method: "GET",
		URL:    fmt.Sprintf("https://openapi.alipay.com/gateway.do?%s", str),
	})
	if err != nil {
		return nil, err
	}
	result := Result{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	notice := result.AlipayTradeFefundResponse
	if notice.Code != "10000" {
		return nil, errors.New(notice.Msg)
	}
	return notice, nil
}

// Query query
func (r *Refund) Query(str string) (*Query, error) {
	body, err := fetch.Cmd(&fetch.Request{
		Method: "GET",
		URL:    fmt.Sprintf("https://openapi.alipay.com/gateway.do?%s", str),
	})
	if err != nil {
		return nil, err
	}
	result := Result{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	query := result.AlipayTradeFastpayRfundQueryResponse
	if query.Code != "10000" {
		return nil, errors.New(query.Msg)
	}
	return query, nil
}

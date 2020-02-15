package alipay

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/tiantour/fetch"
	"github.com/tiantour/imago"
	"github.com/tiantour/rsae"
)

// Refund refund
type Refund struct{}

// NewRefund new refund
func NewRefund() *Refund {
	return &Refund{}
}

// Sign trade sign
func (r *Refund) Sign(args *url.Values, privatePath string) (string, error) {
	query, err := url.QueryUnescape(args.Encode())
	if err != nil {
		return "", err
	}
	privateKey, err := imago.NewFile().Read(privatePath)
	if err != nil {
		return "", err
	}
	sign, err := rsae.NewRSA().Sign(query, privateKey)
	if err != nil {
		return "", err
	}
	args.Add("sign", sign)
	return args.Encode(), nil
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
		return nil, errors.New(notice.SumbMsg)
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
		return nil, errors.New(query.SumbMsg)
	}
	return query, nil
}

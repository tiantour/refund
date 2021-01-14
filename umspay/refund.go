package umspay

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/tiantour/fetch"
)

// Refund refund
type Refund struct{}

// NewRefund new refund
func NewRefund() *Refund {
	return &Refund{}
}

// Apply apply
func (r *Refund) Apply(args *Request) (*Response, error) {
	body, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	header := http.Header{}
	header.Add("Accept", "application/json")
	header.Add("Content-Type", "application/json;charset=utf-8")
	header.Add("Authorization", fmt.Sprintf("OPEN-ACCESS-TOKEN AccessToken=%s", AccessToken))
	body, err = fetch.Cmd(&fetch.Request{
		Method: "POST",
		URL:    "https://api-mop.chinaums.com/v1/netpay/refund",
		Body:   body,
		Header: header,
	})
	if err != nil {
		return nil, err
	}

	result := Response{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	if result.ErrCode != Success {
		return nil, errors.New(result.ErrMsg)
	}
	return &result, nil
}

// Query query
func (r *Refund) Query(args *Query) (*Response, error) {
	body, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	header := http.Header{}
	header.Add("Accept", "application/json")
	header.Add("Content-Type", "application/json;charset=utf-8")
	header.Add("Authorization", fmt.Sprintf("OPEN-ACCESS-TOKEN AccessToken=%s", AccessToken))
	body, err = fetch.Cmd(&fetch.Request{
		Method: "POST",
		URL:    "https://api-mop.chinaums.com/v1/netpay/refund-query",
		Body:   body,
		Header: header,
	})
	if err != nil {
		return nil, err
	}

	result := Response{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	if result.ErrCode != Success {
		return nil, errors.New(result.ErrMsg)
	}
	return &result, nil
}

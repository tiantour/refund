package wxpay

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/tiantour/fetch"
	"github.com/tiantour/rsae"
)

// Refund refund
type Refund struct{}

// NewRefund new refund
func NewRefund() *Refund {
	return &Refund{}
}

// Sign trade sign
func (r *Refund) Sign(args *url.Values, key string) (string, error) {
	query, err := url.QueryUnescape(args.Encode())
	if err != nil {
		return "", err
	}
	query = fmt.Sprintf("%s&key=%s", query, key)
	sign := rsae.NewMD5().Encode(query)
	return strings.ToUpper(sign), nil
}

// Apply apply
func (r *Refund) Apply(args *Sign, certPath, keyPath string) (*Result, error) {
	body, err := xml.Marshal(args)
	if err != nil {
		return nil, err
	}
	header := http.Header{}
	header.Add("Accept", "application/xml")
	header.Add("Content-Type", "application/xml;charset=utf-8")
	body, err = fetch.Cmd(&fetch.Request{
		Method: "POST",
		URL:    "https://api.mch.weixin.qq.com/secapi/pay/refund",
		Body:   body,
		Header: header,
		Cert:   certPath,
		Key:    keyPath,
	})
	if err != nil {
		return nil, err
	}
	result := Result{}
	err = xml.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	if result.ReturnCode != Success {
		return nil, errors.New(result.ReturnMsg)
	}
	if result.ResultCode != Success {
		return nil, errors.New(result.ErrCodeDes)
	}
	return &result, nil
}

// Verify verify
func (r *Refund) Verify(args string, key string) (*Ciphertext, error) {
	body, err := rsae.NewBase64().Decode(args)
	if err != nil {
		return nil, err
	}
	key = rsae.NewMD5().Encode(key)
	key = strings.ToLower(key)
	body, err = rsae.NewAES().ECBDecrypt(body, []byte(key))
	if err != nil {
		return nil, err
	}
	data := Ciphertext{}
	err = xml.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// Query query
func (r *Refund) Query(args *Sign) (*Query, error) {
	body, err := xml.Marshal(args)
	if err != nil {
		return nil, err
	}
	header := http.Header{}
	header.Add("Accept", "application/xml")
	header.Add("Content-Type", "application/xml;charset=utf-8")
	body, err = fetch.Cmd(&fetch.Request{
		Method: "POST",
		URL:    "https://api.mch.weixin.qq.com/pay/refundquery",
		Body:   body,
		Header: header,
	})
	if err != nil {
		return nil, err
	}
	result := Query{}
	err = xml.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	if result.ReturnCode != Success {
		return nil, errors.New(result.ReturnMsg)
	}
	if result.ResultCode != Success {
		return nil, errors.New(result.ErrCodeDes)
	}
	return &result, nil
}

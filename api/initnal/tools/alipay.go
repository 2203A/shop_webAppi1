package tools

import (
	"SX1/shop_api/api/cmd/global"
	"context"
	"github.com/smartwalle/alipay/v3"
	"net/url"
)

type Alipay struct {
	AppID     string
	PublicKey string
	PrickKey  string
}

func NewAlipayClient(a Alipay) *Alipay {
	return &Alipay{
		AppID:     a.AppID,
		PublicKey: a.PublicKey,
		PrickKey:  a.PrickKey,
	}
}

// 类方法接口
type AlipayInterface interface {
}

type Pay struct {
	NotifyURL   string
	ReturnURL   string
	Subject     string
	OutTradeNo  string
	TotalAmount string
	ProductCode string
}

// 支付方法
func (a *Alipay) Pay(p Pay) (*url.URL, error) {
	var y alipay.TradePagePay
	y.NotifyURL = p.NotifyURL
	y.ReturnURL = p.ReturnURL
	y.Subject = p.Subject
	y.OutTradeNo = p.OutTradeNo
	y.TotalAmount = p.TotalAmount
	y.ProductCode = p.ProductCode
	pay, err := global.AlipayClient.TradePagePay(y)
	return pay, err
}

// 退款方法
func (a *Alipay) Refund(code, price string) (*alipay.TradeRefundRsp, error) {
	refund := alipay.TradeRefund{
		OutTradeNo:   code,
		RefundAmount: price,
	}
	tradeRefund, err := global.AlipayClient.TradeRefund(context.Background(), refund)
	return tradeRefund, err
}

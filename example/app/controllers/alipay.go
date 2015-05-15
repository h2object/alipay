package controllers

import (
	"github.com/h2object/alipay"
	"github.com/revel/revel"
	"time"
)

type Alipay struct {
	*revel.Controller
}

func (c Alipay) Payment() revel.Result {
	payment := alipay.NewDirectPayment(time.Now().Format("2006-10-03_10_23_44"), "商品名称", "商品描述", 0.01)
	page := alipay.PaymentPage(payment, "", "")
	return c.RenderHtml(page)
}

func (c Alipay) Return() revel.Result {
	invoice, err := alipay.RevelReturn(c.Request)
	if err != nil {
		return c.RenderError(err)
	}
	revel.INFO.Printf("alipay return invoice: %v", invoice)
	return c.RenderJson(invoice)
}

func (c Alipay) Notify() revel.Result {
	invoice, err := alipay.RevelNotify(c.Request)
	if err != nil {
		return c.RenderError(err)
	}
	revel.INFO.Printf("alipay notify invoice: %v", invoice)
	return c.RenderJson(invoice)
}

func init() {
	alipay.SellerID = "111"
	alipay.SellerKey = "222"
	alipay.SellerEmail = "33"
	alipay.ReturnURL = "http://host:8089/alipay/return"
	alipay.NotifyURL = "http://host:8089/alipay/notify"
}

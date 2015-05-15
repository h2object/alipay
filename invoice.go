package alipay

import (
	"log"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"github.com/revel/revel"
)

type Invoice struct{
	OutTradeNo string `form:"out_trade_no"`
	TradeNo    string `form:"trade_no"`		
	BuyerID    string `form:"buyer_id"`
	BuyerEmail string `form:"buyer_email"`
	TotalFee   float64 `form:"total_fee"`
}

func Return(req *http.Request) (*Invoice, error) {
 	var invoice Invoice
 	queries := req.URL.Query()
 	trade_status := queries.Get("trade_status")
 	if trade_status == "TRADE_FINISHED" || trade_status == "TRADE_SUCCESS" {
 		invoice.OutTradeNo = queries.Get("out_trade_no")
	 	invoice.TradeNo = queries.Get("trade_no")
	 	invoice.BuyerID = queries.Get("buyer_id")
	 	invoice.BuyerEmail = queries.Get("buyer_email")		
		fee, _ := strconv.ParseFloat(queries.Get("total_fee"), 64)
		invoice.TotalFee = fee
		return &invoice, nil
 	}

 	return nil, errors.New("trade status is not success or finished.")
}

func Notify(req *http.Request) (*Invoice, error) {
	var invoice Invoice

 	body, err := ioutil.ReadAll(req.Body)
 	if err != nil {
 		return nil, err
 	}

 	queries, err := url.ParseQuery(string(body))
 	if err != nil {
 		return nil, err
 	}

 	trade_status := queries.Get("trade_status")
 	if trade_status == "TRADE_FINISHED" || trade_status == "TRADE_SUCCESS" {
 		invoice.OutTradeNo = queries.Get("out_trade_no")
	 	invoice.TradeNo = queries.Get("trade_no")
	 	invoice.BuyerID = queries.Get("buyer_id")
	 	invoice.BuyerEmail = queries.Get("buyer_email")		
		fee, _ := strconv.ParseFloat(queries.Get("total_fee"), 64)
		invoice.TotalFee = fee
		log.Println("notify invoice:", invoice)
		return &invoice, nil
 	}

 	return nil, errors.New("trade status is not success or finished.")
}


func RevelReturn(req *revel.Request) (*Invoice, error) {
 	return Return(req.Request)
}

func RevelNotify(req *revel.Request) (*Invoice, error) {
 	return Notify(req.Request)
}

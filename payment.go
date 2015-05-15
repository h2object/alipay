package alipay

import (
	"strconv"
	"fmt"
)

type Payment struct{
	InputCharset string 	`json:"_input_charset"` //网站编码
	Service 	 string 	`json:"service"`        //接口名称
	PaymentType  uint8      `json:"payment_type"`   //支付类型 1：商品购买
	
	SellerID     string 	`json:"partner"`        //合作者身份ID
	SellerEmail  string 	`json:"seller_email"`   //卖家支付宝邮箱

	OutTradeNo 	 string 	`json:"out_trade_no"`   //订单唯一id
	Subject    	 string 	`json:"subject"`        //商品名称
	Description  string 	`json:"body"`           //订单描述
	TotalFee 	 float64 	`json:"total_fee"`      //总价

	ReturnURL    string 	`json:"return_url"`     //同步通知页面
	NotifyURL    string 	`json:"notify_url"`     //异步通知页面

	Signature         string `json:"sign"`           //签名，生成签名时忽略
	SignatureType     string `json:"sign_type"`      //签名类型，生成签名时忽略
}

func NewDirectPayment(order string, subject string, description string, fee float64) *Payment {
	payment := &Payment{
		InputCharset: "utf-8",
		Service: "create_direct_pay_by_user",
		PaymentType: 1,
		SellerID: SellerID,
		SellerEmail: SellerEmail,

		OutTradeNo: order,
		Subject: subject,
		Description: description,
		TotalFee: fee,

		ReturnURL: ReturnURL,
		NotifyURL: NotifyURL,
	}

	payment.Signature = MD5Signature(payment)
	payment.SignatureType = "MD5"
	return payment
}

/* 按照支付宝规则生成sign */
func MD5Signature(payment *Payment) string {
	p := kvpairs {
		kvpair{`_input_charset`, payment.InputCharset},
		kvpair{`out_trade_no`, payment.OutTradeNo},
		kvpair{`partner`, payment.SellerID},
		kvpair{`payment_type`, fmt.Sprintf("%d",payment.PaymentType)},
		kvpair{`notify_url`, payment.NotifyURL},
		kvpair{`return_url`, payment.ReturnURL},
		kvpair{`subject`, payment.Subject},
		kvpair{`total_fee`, fmt.Sprintf("%.2f", payment.TotalFee)},
		kvpair{`body`, payment.Description},
		kvpair{`service`, payment.Service},
		kvpair{`seller_email`, payment.SellerEmail},
	}
	p = p.RemoveEmpty()
	p.Sort()

	sign := md5Sign(p.Join(), SellerKey)
	return sign
}

func PaymentPage(payment *Payment, returnURL string, notifyURL string) string {
	if returnURL != "" {
		payment.ReturnURL = returnURL
	}

	if notifyURL != "" {
		payment.NotifyURL = notifyURL
	}

	return `
		<html>
			<head>
				<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
			</head>
			<body>
				<form id="alipaysubmit" name="alipaysubmit" action="` + AlipayGateway + `?_input_charset=utf-8" method="get" style='display:none;'>
					<input type="hidden" name="_input_charset" value="` + payment.InputCharset + `">
					<input type="hidden" name="body" value="` + payment.Description + `">
					<input type="hidden" name="notify_url" value="` + payment.NotifyURL + `">
					<input type="hidden" name="out_trade_no" value="` + payment.OutTradeNo + `">
					<input type="hidden" name="partner" value="` + payment.SellerID + `">
					<input type="hidden" name="payment_type" value="` + strconv.Itoa(int(payment.PaymentType)) + `">
					<input type="hidden" name="return_url" value="` + payment.ReturnURL + `">
					<input type="hidden" name="seller_email" value="` + payment.SellerEmail + `">
					<input type="hidden" name="service" value="` + payment.Service + `">
					<input type="hidden" name="subject" value="` + payment.Subject + `">
					<input type="hidden" name="total_fee" value="` + fmt.Sprintf("%.2f", payment.TotalFee) + `">
					<input type="hidden" name="sign" value="` + payment.Signature + `">
					<input type="hidden" name="sign_type" value="` + payment.SignatureType + `">
				</form>
				<script>
					document.forms['alipaysubmit'].submit();
				</script>
		</body>
		</html>
	`
}
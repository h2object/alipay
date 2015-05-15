package alipay

const (
	AlipayGateway string = "https://mapi.alipay.com/gateway.do"
)

var (
	SellerID 	string
	SellerKey 	string
	SellerEmail string
	ReturnURL 	string
	NotifyURL 	string
)

/*
	
-	开始

	import "github.com/h2object/alipay"

-	配置

	func init() {
		alipay.SellerID = ""
		alipay.SellerKey = ""
		alipay.SellerEmail = ""
		// default return url
		alipay.ReturnURL = ""
		// default notify url	
		alipay.NotifyURL = ""
	}

-	发起支付

	payment := alipay.NewDirectPayment("订单号","商品名称","商品描述", 0.01)

	page := alipay.PaymentPage(payment, "http://ReturnURL", "http://NotifyURL")

-	账单同步通知

	alipay.Return(r *http.Request) (*Invoice, error)

-	账单异步通知

	alipay.Notify(r *http.Request) (*Invoice, error)

*/

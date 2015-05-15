# alipay
支付宝开发包（pure alipay package base on  net/http package ）

# 快速开发指南

## 引用

````
	import "github.com/h2object/alipay"

````
## 商户配置

````
//! 应用开始前 配置好商户参数

func init() {
	// 商户ID
	alipay.SellerID = ""
	// 商户密钥
	alipay.SellerKey = ""
	// 商户支付宝邮箱
	alipay.SellerEmail = ""
	// 默认同步回调URL
	alipay.ReturnURL = ""
	// 默认异步回调URL
	alipay.NotifyURL = ""
}


````

## 接口说明

### 发起即时支付请求

````
	payment := alipay.NewDirectPayment("订单号","商品名称","商品描述", 0.01)
	
	page := PaymentPage(payment, "http://domain:port/returnURL", "http://domain:port/notifyURL")

	//将 page 内容输出到 http response

````

### 支付同步响应

````
	alipay.Return(r *http.Request) (*Invoice, error)
	alipay.RevelReturn(r *revel.Request) (*Invoice, error)

````

### 支付异步响应

````
	alipay.Notify(r *http.Request) (*Invoice, error)
	alipay.RevelNotify(r *revel.Request) (*Invoice, error)

````

## 样例

样例基于revel框架开发, 请在外网机器进行样例测试:

````
	
	$: go get github.com/h2object/alipay

	//!!! 注意 设置支付宝商户相关参数

	$: revel run github.com/h2object/alipay/example

````
	
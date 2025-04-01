package designpatterns

type PayFactory interface {
	CreatePay() Pay
}

type AliPayFactory struct{}

func (AliPayFactory) CreatePay() Pay {
	return AliPay{}
}

type WeChatPayFactory struct{}

func (WeChatPayFactory) CreatePay() Pay {
	return WeChatPay{}
}

package designpatterns

import "fmt"

type Pay interface {
	PayValue(value int)
}

type AliPay struct{}

func (AliPay) PayValue(value int) {
	fmt.Println("AliPay: ", value)
}

type WeChatPay struct{}

func (WeChatPay) PayValue(value int) {
	fmt.Println("WeChatPay: ", value)
}

type PayType int8

const (
	AliPayType    PayType = 1
	WeChatPayType PayType = 2
)

func NewPay(pt PayType) Pay {
	switch pt {
	case AliPayType:
		return AliPay{}
	case WeChatPayType:
		return WeChatPay{}
	}
	return nil
}

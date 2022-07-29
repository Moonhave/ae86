package temp

import (
	"ae86/internal/enums"
	"ae86/internal/model"
	tele "gopkg.in/telebot.v3"
)

// CustomerInfo TempCustomerInfo - temporary customer info, for storing a customer's cart and other temporary data
type CustomerInfo struct {
	Cart                    []*model.OrderItem
	IsRequiredToSendAddress bool
	PreferredAddress        string
	PreferredPaymentMethod  enums.PaymentMethod
}

// temp user storage
var temporaryCustomerStorage = make(map[uint]*CustomerInfo)

func CreateCustomer(c tele.Context) {
	temporaryCustomerStorage[uint(c.Sender().ID)] = &CustomerInfo{
		Cart:                    []*model.OrderItem{},
		IsRequiredToSendAddress: false,
	}
}

func AddToCart(c tele.Context, item *model.OrderItem) {
	id := uint(c.Sender().ID)
	if temporaryCustomerStorage[id] == nil {
		temporaryCustomerStorage[id] = &CustomerInfo{}
	}
	temporaryCustomerStorage[id].Cart = append(temporaryCustomerStorage[id].Cart, item)
}

func GetCurrentCustomer(c tele.Context) *CustomerInfo {
	id := uint(c.Sender().ID)
	if temporaryCustomerStorage[id] == nil {
		temporaryCustomerStorage[id] = &CustomerInfo{}
	}
	return temporaryCustomerStorage[id]
}

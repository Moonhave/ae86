package handlers

import (
	"ae86/internal/enums"
	"ae86/internal/model"
	"ae86/internal/transport/adapter"
	"ae86/internal/transport/bot/temp"
	"ae86/internal/transport/bot/view"
	"fmt"
	tele "gopkg.in/telebot.v3"
)

func NewOrderHandler(service adapter.ServiceContainer) *OrderHandler {
	return &OrderHandler{service: service}
}

type OrderHandler struct {
	service adapter.ServiceContainer
}

func (h *OrderHandler) SendCart(c tele.Context) error {
	cart := temp.GetCurrentCustomer(c).Cart
	if len(cart) == 0 {
		return c.Send(view.CartEmptyMessage, view.EmptyMenu)
	}
	text := ""
	for _, orderItem := range cart {
		product := orderItem.Product
		text += fmt.Sprintf("%s\n%s\nЦена: %dx%d=%d тенге\n\n", product.Title, product.Description,
			product.Price, orderItem.Amount, product.Price*orderItem.Amount)
	}
	text += "Сумма: " + fmt.Sprintf("%d", priceSum(cart)) + " тенге"
	return c.Send(text, view.CartMenu)
}

func (h *OrderHandler) ClearCart(c tele.Context) error {
	temp.GetCurrentCustomer(c).Cart = []*model.OrderItem{}
	return c.Send(view.CartEmptyMessage, view.EmptyMenu)
}

func (h *OrderHandler) PromptAddressInput(c tele.Context) error {
	if len(temp.GetCurrentCustomer(c).Cart) == 0 {
		return c.Send(view.CartEmptyMessage, view.EmptyMenu)
	}

	temp.GetCurrentCustomer(c).IsRequiredToSendAddress = true

	err := c.Send(view.AddressMenuMessage, view.AddressMenu)
	if err != nil {
		return err
	}
	return c.Respond()
}

func (h *OrderHandler) CancelOrder(c tele.Context) error {
	temp.GetCurrentCustomer(c).IsRequiredToSendAddress = false

	return c.Send(view.MenuMessage, view.Menu)
}

func (h *OrderHandler) SetCashAsPaymentMethod(c tele.Context) error {
	temp.GetCurrentCustomer(c).PreferredPaymentMethod = enums.PaymentMethodCash

	return h.sendOrder(c)
}

func (h *OrderHandler) SetCardAsPaymentMethod(c tele.Context) error {
	temp.GetCurrentCustomer(c).PreferredPaymentMethod = enums.PaymentMethodCard

	return h.sendOrder(c)
}

func (h *OrderHandler) SendOrderList(c tele.Context) error {
	id := uint(c.Sender().ID)
	orderList, err := h.service.Order().GetOrderList(adapter.OrderFilter{
		CustomerID: &id,
	})

	if err != nil {
		return err
	}

	if len(orderList) == 0 {
		return c.Send(view.EmptyMessage, view.EmptyMenu)
	}

	var text string
	for _, order := range orderList {
		text += fmt.Sprintf("Заказ %d:\n%s\n%s\n", order.ID, order.Address, orderStateToString(order.State))
	}

	return c.Send(text, view.EmptyMenu)
}

func (h *OrderHandler) sendOrder(c tele.Context) error {
	order := model.Order{
		CustomerID: uint(c.Sender().ID),
		Address:    temp.GetCurrentCustomer(c).PreferredAddress,
		State:      enums.OrderStatePending,
	}

	_, err := h.service.Order().CreateOrder(order)
	if err != nil {
		return err
	}

	temp.GetCurrentCustomer(c).Cart = []*model.OrderItem{}

	c.Send(view.OrderMessage)
	return c.Send(view.MenuMessage, view.Menu)
}

func priceSum(products []*model.OrderItem) (sum int) {
	for _, product := range products {
		sum += product.Product.Price * product.Amount
	}
	return sum
}

func orderStateToString(state enums.OrderState) string {
	switch state {
	case enums.OrderStatePending:
		return "Ожидает оплаты"
	case enums.OrderStateCanceled:
		return "Отменен"
	case enums.OrderStateAccepted:
		return "Принят"
	case enums.OrderStateDeliveryInProgress:
		return "В пути"
	case enums.OrderStateDelivered:
		return "Доставлен"
	default:
		return "Неизвестно"
	}
}

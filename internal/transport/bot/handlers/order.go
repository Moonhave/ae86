package handlers

import (
	"ae86/internal/enums"
	"ae86/internal/model"
	"ae86/internal/transport/adapter"
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
	order, err := h.tryGetCurrentOrder(c)
	if err != nil {
		return c.Send(err.Error(), view.EmptyMenu)
	}

	orderItems, err := h.service.OrderItem().ListByOrderID(order.ID)
	if err != nil {
		return err
	}

	text := ""
	priceSum := 0
	for _, orderItem := range orderItems {
		product, err := h.service.Product().ByID(orderItem.ProductID)
		if err != nil {
			return err
		}
		text += fmt.Sprintf("%s\nЦена: %dx%d=%d тенге\n\n", product.Title,
			product.Price, orderItem.Amount, product.Price*orderItem.Amount)
		priceSum += product.Price * orderItem.Amount
	}
	text += "Сумма: " + fmt.Sprintf("%d", priceSum) + " тенге"
	return c.Send(text, view.CartMenu)
}

func (h *OrderHandler) ClearCart(c tele.Context) error {
	order, err := h.tryGetCurrentOrder(c)
	if err != nil {
		return c.Send(err.Error(), view.EmptyMenu)
	}

	err = h.service.Order().Delete(order.ID)
	if err != nil {
		return err
	}

	return c.Send(view.CartEmptyMessage, view.EmptyMenu)
}

func (h *OrderHandler) PromptAddressInput(c tele.Context) error {
	_, err := h.tryGetCurrentOrder(c)
	if err != nil {
		return c.Send(err.Error(), view.EmptyMenu)
	}

	c.Bot().Handle(tele.OnText, func(c tele.Context) error {
		return h.SetOrderAddress(c)
	})

	err = c.Send(view.AddressMenuMessage, view.AddressMenu)
	if err != nil {
		return err
	}
	return c.Respond()
}

func (h *OrderHandler) SetOrderAddress(c tele.Context) error {
	order, err := h.tryGetCurrentOrder(c)
	if err != nil {
		return c.Send(err.Error(), view.EmptyMenu)
	}

	order.Address = c.Message().Text

	if err = h.service.Order().Update(order.ID, order); err != nil {
		return err
	}

	c.Bot().Handle(tele.OnText, func(c tele.Context) error {
		return c.Send(view.UnknownCommandMessage, view.EmptyMenu)
	})

	return c.Send(view.PaymentMethodMenuMessage, view.PaymentMethodMenu)
}

func (h *OrderHandler) CancelOrder(c tele.Context) error {
	c.Bot().Handle(tele.OnText, func(c tele.Context) error {
		return c.Send(view.UnknownCommandMessage, view.EmptyMenu)
	})

	return c.Send(view.MenuMessage, view.Menu)
}

func (h *OrderHandler) SetCashAsPaymentMethod(c tele.Context) error {
	order, err := h.tryGetCurrentOrder(c)
	if err != nil {
		return c.Send(err.Error(), view.EmptyMenu)
	}

	order.PaymentMethod = enums.PaymentMethodCash

	if err = h.service.Order().Update(order.ID, order); err != nil {
		return err
	}

	return h.sendOrder(c)
}

func (h *OrderHandler) SetCardAsPaymentMethod(c tele.Context) error {
	order, err := h.tryGetCurrentOrder(c)
	if err != nil {
		return c.Send(err.Error(), view.EmptyMenu)
	}

	order.PaymentMethod = enums.PaymentMethodCard

	if err = h.service.Order().Update(order.ID, order); err != nil {
		return err
	}

	return h.sendOrder(c)
}

func (h *OrderHandler) SendOrderList(c tele.Context) error {
	customer, err := h.service.Customer().ByExternalID(uint(c.Sender().ID))
	if err != nil {
		return err
	}

	orderList, err := h.service.Order().ListBy(adapter.OrderFilter{
		CustomerID: &customer.ID,
	})

	if err != nil {
		return err
	}

	if len(orderList) == 0 {
		return c.Send(view.EmptyMessage, view.EmptyMenu)
	}

	var text string
	for _, order := range orderList {
		text += fmt.Sprintf("**Заказ №%d**:\nАдрес - %s\nСостояние - %s\n", order.ID, order.Address, orderStateToString(order.State))
	}

	return c.Send(text, view.EmptyMenu)
}

func (h *OrderHandler) sendOrder(c tele.Context) error {
	order, err := h.tryGetCurrentOrder(c)
	if err != nil {
		return c.Send(err.Error(), view.EmptyMenu)
	}

	order.State = enums.OrderStateAccepted

	if err = h.service.Order().Update(order.ID, order); err != nil {
		return err
	}

	c.Send(fmt.Sprintf("%s\nID Заказа: %d", view.OrderMessage, order.ID), view.EmptyMenu)
	return c.Send(view.MenuMessage, view.Menu)
}

func (h *OrderHandler) tryGetCurrentOrder(c tele.Context) (model.Order, error) {
	customer, err := h.service.Customer().ByExternalID(uint(c.Sender().ID))
	if err != nil {
		return model.Order{}, err
	}

	state := enums.OrderStatePending
	orders, err := h.service.Order().ListBy(adapter.OrderFilter{CustomerID: &customer.ID, State: &state})
	if err != nil {
		return model.Order{}, err
	}

	if len(orders) == 0 {
		return model.Order{}, fmt.Errorf(view.CartEmptyMessage)
	}

	return orders[0], nil
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

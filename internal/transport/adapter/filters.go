package adapter

import "ae86/internal/enums"

type ProductFilter struct {
	Title      *string
	MinPrice   *int
	MaxPrice   *int
	IsActive   *bool
	IsDeleted  *bool
	CategoryID *uint
}

type OrderFilter struct {
	Address       *string
	State         *enums.OrderState
	PaymentMethod *enums.PaymentMethod
	CustomerID    *uint
}

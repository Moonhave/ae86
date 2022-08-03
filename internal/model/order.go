package model

import (
	"ae86/internal/enums"
	"time"
)

type Order struct {
	ID                 uint                `gorm:"column:id;primaryKey"`
	Address            string              `gorm:"column:address;not null"`
	State              enums.OrderState    `gorm:"column:state;not null"`
	PaymentMethod      enums.PaymentMethod `gorm:"column:payment_method;not null"`
	CancellationReason string              `gorm:"column:cancellation_reason"`
	CustomerID         uint                `gorm:"column:customer_id"`
	CreatedAt          time.Time           `gorm:"column:created_at"`
	UpdatedAt          time.Time           `gorm:"column:updated_at"`

	Customer *Customer
}

func (Order) TableName() string {
	return "orders"
}

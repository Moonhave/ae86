package model

import "time"

type Store struct {
	ID               uint          `gorm:"column:id;primaryKey"`
	Title            string        `gorm:"column:title;not null"`
	Info             string        `gorm:"column:info"`
	Address          string        `gorm:"column:address;not null"`
	Image            string        `gorm:"column:image;not null"`
	AvgDeliveryTime  time.Duration `gorm:"column:avg_delivery_time;not null"`
	WorkingHourBegin time.Time     `gorm:"column:working_hour_begin;not null"`
	WorkingHourEnd   time.Time     `gorm:"column:working_hour_end;not null"`
	MinOrderPrice    int           `gorm:"column:min_order_price;check:min_order_price > 0;not null"`
	DeliveryPrice    int           `gorm:"column:delivery_price;check:delivery_price > 0;not null"`
	ContactPhone     string        `gorm:"column:contact_phone;not null"`
	ManagerID        uint          `gorm:"column:manager_id"`
	CreatedAt        time.Time     `gorm:"column:created_at"`
	UpdatedAt        time.Time     `gorm:"column:updated_at"`

	Manager *Manager
}

func (Store) TableName() string {
	return "store"
}

func (s Store) IsOpen() bool {
	n := time.Now()
	return s.WorkingHourBegin.Before(n) && s.WorkingHourEnd.After(n)
}

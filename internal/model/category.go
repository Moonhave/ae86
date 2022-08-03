package model

import "time"

type Category struct {
	ID        uint      `gorm:"column:id;primaryKey"`
	Title     string    `gorm:"column:title;not null"`
	StoreID   uint      `gorm:"column:store_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`

	Store *Store
}

func (Category) TableName() string {
	return "category"
}

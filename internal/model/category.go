package model

import "time"

type Category struct {
	ID        uint      `gorm:"column:id;primaryKey"`
	Title     string    `gorm:"column:title;not null"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (Category) TableName() string {
	return "category"
}

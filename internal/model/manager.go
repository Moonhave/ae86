package model

import "time"

type Manager struct {
	ID        uint      `gorm:"column:id;primaryKey"`
	Username  string    `gorm:"column:username;unique;not null"`
	Password  string    `gorm:"column:password;not null"`
	FirstName string    `gorm:"column:first_name;not null"`
	LastName  string    `gorm:"column:last_name;not null"`
	Phone     string    `gorm:"column:phone;not null"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (Manager) TableName() string {
	return "manager"
}

package models

import (
	"time"

	"gorm.io/gorm"
)


type Payment struct {
	ID                  uint      `gorm:"primaryKey"`
	UserID              uint      `gorm:"not null;index"`
	User                User      `gorm:"constraint:OnDelete:CASCADE;"`
	SubscriptionID      *uint     `gorm:"index"`
	Subscription        *Subscription
	Amount              float64   `gorm:"not null"`
	StripePaymentStatus string    `gorm:"size:50"`
	StripePaymentIntent string    `gorm:"size:100"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`
}

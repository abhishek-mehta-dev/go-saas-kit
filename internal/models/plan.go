package models

import (
	"time"

	"github.com/abhishek-mehta-dev/go-saas-kit.git/internal/enums"
	"gorm.io/gorm"
)


type Plan struct {
	ID          uint                `gorm:"primaryKey"`
	Name        enums.PlanType      `gorm:"type:int;not null"`
	Period      enums.BillingCycle  `gorm:"type:int;not null"`
	Price       float64             `gorm:"not null"`
	Description string              `gorm:"size:255"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt      `gorm:"index"`

	// Associations
	Subscriptions []Subscription
}

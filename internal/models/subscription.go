package models

import (
	"time"

	"github.com/abhishek-mehta-dev/go-saas-kit.git/internal/enums"
	"gorm.io/gorm"
)

type Subscription struct {
	ID              uint                         `gorm:"primaryKey"`
	UserID          uint                         `gorm:"not null;index"`
	User            User                         `gorm:"constraint:OnDelete:CASCADE;"`
	Status          enums.SubscriptionStatus     `gorm:"type:int;not null;default:0"`
	NextPaymentDate time.Time
	IsFreeTrial     bool                         `gorm:"default:false"`
	IsRenewCancelled bool                        `gorm:"default:false"`

	PlanID uint `gorm:"not null;index"`
	Plan   Plan `gorm:"constraint:OnDelete:CASCADE;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

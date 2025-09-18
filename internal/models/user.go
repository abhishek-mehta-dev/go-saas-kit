package models

import (
	"time"

	"github.com/abhishek-mehta-dev/go-saas-kit.git/internal/enums"
	"gorm.io/gorm"
)

type User struct {
	ID                  uint           `gorm:"primaryKey"`
	FirstName           string         `gorm:"size:100;not null"`
	LastName            string         `gorm:"size:100"`
	UserName            string         `gorm:"uniqueIndex;size:50;not null"`
	Email               string         `gorm:"uniqueIndex;size:100;not null"`
	Password            string         `gorm:"not null"`
	UserType            enums.UserRole `gorm:"type:varchar(20);not null;default:'member'"`
	StripeCustomerId    string         `gorm:"size:100"`
	StripePaymentMethodId string       `gorm:"size:100"`

	CardLast4            string `gorm:"size:4"`
    CardBrand            string `gorm:"size:50"`
		
	
	RefreshToken        string         `gorm:"size:255"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`

	// Associations
	Subscriptions []Subscription
	Payments      []Payment
}

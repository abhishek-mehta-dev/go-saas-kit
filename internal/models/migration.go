package models

import "time"

type Migration struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"unique;not null"` 
	AppliedAt time.Time `gorm:"autoCreateTime"`
}

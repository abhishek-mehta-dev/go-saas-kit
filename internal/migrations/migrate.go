package migrations

import (
	"log"
	"time"

	"github.com/abhishek-mehta-dev/go-saas-kit.git/internal/db"
	"github.com/abhishek-mehta-dev/go-saas-kit.git/internal/models"
)

// RunMigrations applies schema migrations in order
func RunMigrations() {
	// Step 1: Ensure migrations table exists
	if err := db.DB.AutoMigrate(&models.Migration{}); err != nil {
		log.Fatalf("Failed to migrate migrations table: %v", err)
	}

	// Step 2: Check if this migration batch already ran
	batchName := "initial_schema_v1" // Change when you add new fields/models

	var count int64
	db.DB.Model(&models.Migration{}).Where("name = ?", batchName).Count(&count)

	if count > 0 {
		log.Printf("Migration already applied: %s", batchName)
		return
	}

	// Step 3: Run all migrations together
	log.Printf("Running migration batch: %s", batchName)
	if err := db.DB.AutoMigrate(models.AllModels...); err != nil {
		log.Fatalf("Failed migration batch %s: %v", batchName, err)
	}

	// Step 4: Record applied migration
	db.DB.Create(&models.Migration{
		Name:      batchName,
		AppliedAt: time.Now(),
	})

	log.Printf("Migration batch applied: %s", batchName)
}

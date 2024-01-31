package database

import (
	"flag"
	"fmt"
	"log"

	"github.com/Imamsubekti26/Perpustakaan_Go/models"
)

func (db *dbInstance) Migrate() {
	migrator := db.Connection.Migrator()

	isMigrate := flag.Bool("migrate", false, "Run Gorm migration")
	isForce := flag.Bool("migrate:force", false, "Force Gorm migration (drop tables if exists)")

	flag.Parse()

	if *isMigrate || *isForce {

		modelList := []interface{}{
			&models.Users{},
			&models.Presences{},
			&models.Borrows{},
			&models.Books{},
			&models.Categories{},
		}

		for _, model := range modelList {
			if *isForce {
				if err := migrator.DropTable(model); err != nil {
					log.Fatalf("Failed to drop table %T: %s", model, err)
				}
			}

			if err := migrator.CreateTable(model); err != nil {
				log.Fatalf("Failed to migrate table %T: %s", model, err)
			}
		}

		fmt.Println("Database migration completed.")
	}
}

package database

import (
	"flag"
	"fmt"

	"github.com/Imamsubekti26/Perpustakaan_Go/models"
	xlogger "github.com/Imamsubekti26/Perpustakaan_Go/utils/XLogger"
)

func (db *dbInstance) Migrate() error {
	migrator := db.Connection.Migrator()

	isMigrate := flag.Bool("migrate", false, "Run Gorm migration")
	isForce := flag.Bool("migrate:force", false, "Force Gorm migration (drop tables if exists)")

	flag.Parse()

	if *isMigrate || *isForce {

		modelList := []interface{}{
			&models.Users{},
			&models.Presences{},
			&models.Books{},
			&models.Categories{},
			&models.Borrows{},
		}

		for _, model := range modelList {
			if *isForce {
				if err := migrator.DropTable(model); err != nil {
					return xlogger.Errorf("Failed to drop table %T: %s", model, err)
				}
			}
		}

		if err := db.Connection.AutoMigrate(modelList...); err != nil {
			return xlogger.Errorf("Failed to migrate : %s", err)
		}

		fmt.Println("Database migration completed.")
	}
	return nil
}

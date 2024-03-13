package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	newMigration := &gormigrate.Migration{
		ID: "202403071143_create_users_table",
		Migrate: func(tx *gorm.DB) error {
			sql := `CREATE TABLE IF NOT EXISTS users (
    					id INT AUTO_INCREMENT PRIMARY KEY,
    					first_name VARCHAR(150),
    					last_name VARCHAR(150),
    					email VARCHAR(70) UNIQUE,
    					hashed_password VARCHAR(64),
    					created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    					updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
					);`

			if err := tx.Exec(sql).Error; err != nil {
				return err
			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("users")
		},
	}

	MigrationsToExec = append(MigrationsToExec, newMigration)
}

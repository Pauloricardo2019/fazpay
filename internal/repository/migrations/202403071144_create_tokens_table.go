package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	newMigration := &gormigrate.Migration{
		ID: "202403071144_create_tokens_table",
		Migrate: func(tx *gorm.DB) error {
			sql := `CREATE TABLE IF NOT EXISTS tokens (
							id       bigint AUTO_INCREMENT PRIMARY KEY,
							user_id 	integer,
							value     varchar(36) unique,
							created_at timestamp default CURRENT_TIMESTAMP,
							expires_at timestamp default CURRENT_TIMESTAMP,
							foreign key (user_id) references users(id) ON DELETE CASCADE 
					);`

			if err := tx.Exec(sql).Error; err != nil {
				return err
			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("tokens")
		},
	}

	MigrationsToExec = append(MigrationsToExec, newMigration)
}

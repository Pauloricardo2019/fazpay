package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	newMigration := &gormigrate.Migration{
		ID: "202211122353_create_tokens_table",
		Migrate: func(tx *gorm.DB) error {
			sql := `create table "tokens" (
							"id"        serial primary key,
							"user_id" 	bigint,
							"value"     varchar(255),
							"created_at" timestamp default CURRENT_TIMESTAMP,
							"expires_at" timestamp default CURRENT_TIMESTAMP,
							foreign key ("user_id") references "users"("id") ON DELETE CASCADE 
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

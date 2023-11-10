package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	newMigration := &gormigrate.Migration{
		ID: "202211100541_create_users_table",
		Migrate: func(tx *gorm.DB) error {
			sql := `create table "users" (
    					"id" serial primary key,
    					"full_name" varchar(150),
                        "email" varchar(60),
                        "login" varchar(60),
                        "hashed_password" varchar(100),
						"last_login" timestamp default CURRENT_TIMESTAMP,
                        "created_at" timestamp default CURRENT_TIMESTAMP,
                        "updated_at" timestamp default CURRENT_TIMESTAMP
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

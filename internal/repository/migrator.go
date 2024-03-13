package repository

import (
	"context"
	"errors"
	"fmt"
	repositoryIntf "github.com/Pauloricardo2019/teste_fazpay/internal/repository/interface"
	"github.com/Pauloricardo2019/teste_fazpay/internal/repository/migrations"
	"github.com/go-gormigrate/gormigrate/v2"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
	"log"
)

type Migrator struct {
	*BaseRepository
}

func NewMigrator(db *gorm.DB) repositoryIntf.Migrator {
	baseRepo := NewBaseRepository(db)
	return &Migrator{
		baseRepo,
	}
}

// ExecuteMigrations execute the pending migrations.
func (m *Migrator) ExecuteMigrations(ctx context.Context) error {
	conn, err := m.GetConnection(ctx)
	if err != nil {
		return errors.New("error connection")
	}

	migrationsToExec := migrations.GetMigrationsToExec()
	migrator := gormigrate.New(conn, gormigrate.DefaultOptions, migrationsToExec)

	if err = migrator.Migrate(); err != nil {
		return errors.New(fmt.Sprintf("Could not migrate. %s", err.Error()))
	}

	log.Printf("Migration did run successfully")

	return nil
}

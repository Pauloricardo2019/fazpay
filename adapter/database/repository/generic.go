package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type GenericRepository[T interface{}] struct {
	*BaseRepository
	db *gorm.DB
}

func newGenericRepository[T interface{}](db *gorm.DB) *GenericRepository[T] {
	baseRepository := NewBaseRepository(db)
	return &GenericRepository[T]{
		BaseRepository: baseRepository,
		db:             db,
	}
}

func (g *GenericRepository[T]) Create(ctx context.Context, obj *T) (*T, error) {
	conn, err := g.GetConnection(ctx)
	if err != nil {
		return nil, errors.New("error connection")
	}

	err = conn.Create(obj).Error
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (g *GenericRepository[T]) CreateMany(ctx context.Context, objs []T) ([]T, error) {
	conn, err := g.GetConnection(ctx)
	if err != nil {
		return nil, errors.New("error connection")
	}

	err = conn.Create(objs).Error
	if err != nil {
		return nil, err
	}

	return objs, nil
}

// GetList requests the database to get a list of gyms.
func (g *GenericRepository[T]) GetList(ctx context.Context, limit int, offset int) ([]T, error) {
	conn, err := g.GetConnection(ctx)
	if err != nil {
		return nil, err
	}

	objs := make([]T, 0)

	err = conn.
		Limit(limit).
		Offset(offset).
		Find(&objs).
		Error

	if err != nil {
		return nil, err
	}

	return objs, nil
}

// GetCount requests the database to get all gyms.
func (g *GenericRepository[T]) GetCount(ctx context.Context) (int64, error) {
	conn, err := g.GetConnection(ctx)
	if err != nil {
		return 0, err
	}

	objs := make([]T, 0)
	var count int64
	err = conn.Find(&objs).Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

// GetById requests the database to get a gym.
func (g *GenericRepository[T]) GetById(ctx context.Context, id uint64) (bool, *T, error) {
	conn, err := g.GetConnection(ctx)
	if err != nil {
		return false, nil, err
	}

	obj := new(T)
	if err = conn.First(obj, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil, nil
		}
		return false, nil, err
	}

	return true, obj, nil
}

// Update requests the database to update a gym.
func (g *GenericRepository[T]) Update(ctx context.Context, obj *T) error {
	conn, err := g.GetConnection(ctx)
	if err != nil {
		return err
	}

	err = conn.Save(obj).Error
	if err != nil {
		return err
	}

	return nil
}

func (g *GenericRepository[T]) UpdateMany(ctx context.Context, objs []T) error {
	conn, err := g.GetConnection(ctx)
	if err != nil {
		return err
	}

	err = conn.Save(objs).Error
	if err != nil {
		return err
	}

	return nil
}

// Delete requests the database to delete a gym.
func (g *GenericRepository[T]) Delete(ctx context.Context, id uint64) error {
	conn, err := g.GetConnection(ctx)
	if err != nil {
		return err
	}

	obj := new(T)

	var count int64
	if err := conn.Model(obj).Where("id = ?", id).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		return gorm.ErrRecordNotFound
	}

	err = conn.Delete(obj, id).Error
	return err
}

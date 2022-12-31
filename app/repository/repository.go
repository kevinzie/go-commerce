package repository

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strconv"
)

//gorm generic repository

type repository[T any] struct {
	db *gorm.DB
}

func NewRepository[T any](db *gorm.DB) *repository[T] {
	return &repository[T]{
		db: db,
	}
}

func (r *repository[T]) Add(entity *T, ctx context.Context) error {
	tx := r.db.Begin().Debug()
	tx.WithContext(ctx).Create(&entity)
	if tx.Error != nil {
		tx.Rollback()
	}
	return tx.Commit().Error
}

func (r *repository[T]) AddAll(entity *[]T, ctx context.Context) error {
	return r.db.WithContext(ctx).Create(&entity).Error
}

func (r *repository[T]) GetById(id uuid.UUID, ctx context.Context) (*T, error) {
	var entity T
	err := r.db.WithContext(ctx).Model(&entity).Where("uuid = ? AND status = ?", id, "active").FirstOrInit(&entity).Error
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *repository[T]) Get(params *[]T, ctx context.Context) *T {
	var entity T
	r.db.WithContext(ctx).Where(&params).FirstOrInit(&entity)
	return &entity
}

func (r *repository[T]) GetAll(ctx context.Context, c *fiber.Ctx) (interface{}, error) {
	var entities []T

	err := r.db.Debug().WithContext(ctx).Scopes(Paginate(c)).Find(&entities).Order("created_at desc").Error
	////err := r.db.WithContext(ctx).Find(&entities).Order("created_at desc").Error
	if err != nil {
		return nil, err
	}
	return &entities, nil
}

func (r *repository[T]) GetByEmail(email string, ctx context.Context) (*T, error) {
	var entity *T
	err := r.db.WithContext(ctx).Where("email = ?", email).Where("status = ?", "active").Find(&entity).Error

	if err != nil {
		fmt.Println("error ga didsini", err)
		return nil, err
	}

	return entity, nil
	//if err != nil {
	//	return nil, err
	//}
	//return &entities, nil
}

func (r *repository[T]) Where(params *T, ctx context.Context) (*[]T, error) {
	var entities []T
	err := r.db.WithContext(ctx).Where(&params).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return &entities, nil
}

func (r *repository[T]) Update(id string, entity *T, ctx context.Context) error {
	return r.db.WithContext(ctx).Where("uuid = ?", id).Updates(&entity).Error
}

func (r repository[T]) UpdateAll(entities *[]T, ctx context.Context) error {
	return r.db.WithContext(ctx).Save(&entities).Error
}

func (r *repository[T]) Delete(id int, ctx context.Context) error {
	var entity T
	//return r.db.WithContext(ctx).Where("id = ?", id).FirstOrInit(&entity).UpdateColumn("deleted_at", time.Now()).Delete(&entity).Error
	return r.db.WithContext(ctx).Where("id = ?", id).First(&entity).UpdateColumn("status", "inactive").Delete(&entity).Error

}

func (r *repository[T]) SkipTake(skip int, take int, ctx context.Context) (*[]T, error) {
	var entities []T
	err := r.db.WithContext(ctx).Offset(skip).Limit(take).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return &entities, nil
}

func (r *repository[T]) Count(ctx context.Context) int64 {
	var entity T
	var count int64
	r.db.WithContext(ctx).Model(&entity).Count(&count)
	return count
}

func (r *repository[T]) CountWhere(params *T, ctx context.Context) int64 {
	var entity T
	var count int64
	r.db.WithContext(ctx).Model(&entity).Where(&params).Count(&count)
	return count
}

func Paginate(c *fiber.Ctx) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.Query("page"))
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(c.Query("limit"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

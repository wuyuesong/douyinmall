package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type Category struct {
	Base
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Products    []Product `json:"product" gorm:"many2many:product_category"`
}

func (c Category) TableName() string {
	return "category"
}

type CategoryQuery struct {
	ctx context.Context
	db  *gorm.DB
}

type PaginatedResult struct {
	Categories []Category
	TotalCount int64 // 总记录数
	TotalPages int64 // 总页数
}

func (c CategoryQuery) GetProductsByCategoryName(page int32, pageSize int32, name string) (result PaginatedResult, err error) {
	// 参数校验
	if page < 1 || pageSize < 1 {
		return result, fmt.Errorf("invalid page or pageSize")
	}

	// 查询总记录数
	err = c.db.Model(&Category{}).Where(&Category{Name: name}).Count(&result.TotalCount).Error
	if err != nil {
		return
	}

	// 计算总页数
	result.TotalPages = (result.TotalCount + int64(pageSize) - 1) / int64(pageSize)

	// 分页查询
	offset := (page - 1) * pageSize
	err = c.db.WithContext(c.ctx).
		Model(&Category{}).
		Where(&Category{Name: name}).
		Preload("Products", func(db *gorm.DB) *gorm.DB { // 关键修改点
			return db.Order("product.id ASC") // 按 product_id 降序
		}).
		Limit(int(pageSize)).
		Offset(int(offset)).
		Find(&result.Categories).Error
	return
}

func NewCategoryQuery(ctx context.Context, db *gorm.DB) *CategoryQuery {
	return &CategoryQuery{
		ctx: ctx,
		db:  db,
	}
}

package model

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Product struct {
	Base
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Picture     string     `json:"picture"`
	Price       float32    `json:"price"`
	Stock       uint32     `json:"stock"`
	Categories  []Category `json:"categories" gorm:"many2many:product_category"`
}

func (p Product) TableName() string {
	return "product"
}

type ProductQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func (p ProductQuery) GetById(ProductId int) (product Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).First(&product, ProductId).Error
	return
}

type CachedProductQuery struct {
	productQuery *ProductQuery
	cacheClient  *redis.Client
	prefix       string
}

func (c CachedProductQuery) GetById(productId int) (product Product, err error) {
	// fmt.Print("productQuery", productQuery)
	// fmt.Print("cacheClient", cacheClient)
	cacheKey := fmt.Sprintf("%s_%s_%d", c.prefix, "product_by_id", productId)
	cachedResult := c.cacheClient.Get(c.productQuery.ctx, cacheKey)

	err = func() error {
		err1 := cachedResult.Err()
		if err1 != nil {
			return err1
		}
		cachedResultByte, err2 := cachedResult.Bytes()
		if err2 != nil {
			return err2
		}
		err3 := json.Unmarshal(cachedResultByte, &product)
		if err3 != nil {
			return err3
		}
		return nil
	}()
	if err != nil {
		product, err = c.productQuery.GetById(productId)
		if err != nil {
			return Product{}, err
		}
		encoded, err := json.Marshal(product)
		if err != nil {
			return product, nil
		}
		_ = c.cacheClient.Set(c.productQuery.ctx, cacheKey, encoded, time.Hour)
	}
	return
}

func NewCachedProductQuery(pq *ProductQuery, cacheClient *redis.Client) CachedProductQuery {
	return CachedProductQuery{productQuery: pq, cacheClient: cacheClient, prefix: "douyin_mall"}
}

func (p ProductQuery) SearchProducts(q string) (products []*Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).Find(&products, "name like? or description like?",
		"%"+q+"%", "%"+q+"%",
	).Error
	return
}

type ProductAllResult struct {
	Products   []Product
	TotalCount int64 // 总记录数
	TotalPages int64 // 总页数
}

func (c ProductQuery) ListProductsAll(page int32, pageSize int32, name string) (result ProductAllResult, err error) {
	// 参数校验
	if page < 1 || pageSize < 1 {
		return result, fmt.Errorf("invalid page or pageSize")
	}

	// 查询总记录数
	err = c.db.Model(&Product{}).Count(&result.TotalCount).Error
	if err != nil {
		return
	}

	// 计算总页数
	result.TotalPages = (result.TotalCount + int64(pageSize) - 1) / int64(pageSize)

	// 分页查询，按id升序排列
	offset := (page - 1) * pageSize
	err = c.db.WithContext(c.ctx).
		Model(&Product{}).
		Order("id ASC"). // 新增排序条件
		Limit(int(pageSize)).
		Offset(int(offset)).
		Find(&result.Products).Error
	return
}

func (c *ProductQuery) AddProduct(item *Product, categoryId uint) (id int, err error) {
	// 开启事务
	tx := c.db.WithContext(c.ctx).Begin()
	if tx.Error != nil {
		return 0, tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			err = fmt.Errorf("transaction panic: %v", r)
		}
	}()

	// 1. 创建商品基础信息
	if err = tx.Create(item).Error; err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("创建商品失败: %v", err)
	}

	// 2. 验证分类存在性
	var category Category
	if err = tx.First(&category, categoryId).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, fmt.Errorf("分类ID %d 不存在", categoryId)
		}
		return 0, fmt.Errorf("分类查询失败: %v", err)
	}

	// 3. 建立关联
	if err = tx.Model(item).Association("Categories").Append(&category); err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("关联分类失败: %v", err)
	}

	// 提交事务
	if err = tx.Commit().Error; err != nil {
		return 0, fmt.Errorf("事务提交失败: %v", err)
	}

	return item.ID, nil
}

func NewProductQuery(ctx context.Context, db *gorm.DB) *ProductQuery {
	return &ProductQuery{
		ctx: ctx,
		db:  db,
	}
}

package model

import (
	"context"

	"gorm.io/gorm"
)

type Consignee struct {
	StreetAddress string
	City          string
	State         string
	Country       string
	ZipCode       string
}

const (
	OrderStatusPending   = 0 // 待支付
	OrderStatusPaid      = 1 // 支付成功
	OrderStatusCancelled = 2 // 取消支付
)

type Order struct {
	gorm.Model
	OrderId    string      `gorm:"type:varchar(100);uniqueIndex"`
	UserId     uint32      `gorm:"type:int(11)"`
	Consignee  Consignee   `gorm:"embedded"`
	Status     int         `gorm:"type:tinyint(1);default:0;index;comment:0-待支付 1-支付成功 2-取消支付"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderIdRefer;references:OrderId"`
}

func (Order) TableName() string {
	return "order"
}

func ListOrder(ctx context.Context, db *gorm.DB, userId uint32) ([]*Order, error) {
	var orders []*Order
	err := db.WithContext(ctx).Where("user_id = ?", userId).Preload("OrderItems").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

package entity

import "time"

type Product struct {
	ID        int64     `gorm:"primaryKey;column:id;autoIncrement"`
	Name      string    `gorm:"column:name"`
	Price     int64     `gorm:"column:price"`
	Stock     int64     `gorm:"column:stock"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (p *Product) TableName() string {
	return "products"
}
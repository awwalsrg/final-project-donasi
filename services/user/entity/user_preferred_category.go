package entity

import "time"

type UserPreferredCategory struct {
	ID         uint      `gorm:"primaryKey;column:id" json:"id"`
	UserID     uint      `gorm:"column:user_id;not null" json:"user_id"`
	CategoryID uint      `gorm:"column:category_id;not null" json:"category_id"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
}

func (UserPreferredCategory) TableName() string { return "user_preferred_categories" }

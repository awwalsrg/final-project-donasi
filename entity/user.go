package entity

import "time"

type User struct {
	ID                uint      `gorm:"primaryKey;column:id" json:"id"`
	Name              string    `gorm:"type:varchar(100);not null" json:"name"`
	Email             string    `gorm:"type:varchar(255);unique;not null" json:"email"`
	PasswordHash      string    `gorm:"type:varchar(255);not null" json:"-"`
	Role              string    `gorm:"type:varchar(50);not null" json:"role"`
	NotifyNewCampaign bool      `gorm:"default:false" json:"notify_new_campaign"`
	CreatedAt         time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (User) TableName() string { return "users" }

type UserPreferredCategory struct {
	ID         uint      `gorm:"primaryKey;column:id" json:"id"`
	UserID     uint      `gorm:"column:user_id" json:"user_id"`
	CategoryID uint      `gorm:"column:category_id" json:"category_id"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`

	User User `gorm:"foreignKey:UserID" json:"user"`
}

func (UserPreferredCategory) TableName() string { return "user_preferred_categories" }

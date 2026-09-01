package entity

import "time"

type Category struct {
	ID          uint   `gorm:"primaryKey;column:id" json:"id"`
	Name        string `gorm:"type:varchar(60);unique;not null" json:"name"`
	Description string `gorm:"type:varchar(255)" json:"description"`
}

func (Category) TableName() string { return "categories" }

type Campaign struct {
	ID              uint       `gorm:"primaryKey;column:id" json:"id"`
	Title           string     `gorm:"type:varchar(200);not null" json:"title"`
	Description     string     `gorm:"type:text;not null" json:"description"`
	CollectedAmount float64    `gorm:"type:numeric(15,2);not null;default:0" json:"collected_amount"`
	Deadline        time.Time  `gorm:"type:date;not null" json:"deadline"`
	Status          string     `gorm:"type:varchar(50);not null;default:'active'" json:"status"`
	CreatedByUserID uint       `gorm:"column:created_by_user_id;not null" json:"created_by_user_id"`
	CreatedAt       time.Time  `gorm:"column:created_at" json:"created_at"`
	DeletedAt       *time.Time `gorm:"column:deleted_at;index" json:"deleted_at"`

	Categories []Category `gorm:"many2many:campaign_categories;joinForeignKey:CampaignID;joinReferences:CategoryID" json:"categories"`
}

func (Campaign) TableName() string { return "campaigns" }

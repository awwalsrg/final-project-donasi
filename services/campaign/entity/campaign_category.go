package entity

type CampaignCategory struct {
	ID         uint `gorm:"primaryKey;column:id" json:"id"`
	CampaignID uint `gorm:"column:campaign_id;not null" json:"campaign_id"`
	CategoryID uint `gorm:"column:category_id;not null" json:"category_id"`
}

func (CampaignCategory) TableName() string { return "campaign_categories" }

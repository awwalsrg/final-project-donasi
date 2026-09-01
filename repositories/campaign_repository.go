package repositories

import (
	"donasi-app/entity"

	"gorm.io/gorm"
)

type CampaignRepository interface {
	FindAll() ([]entity.Campaign, error)
	FindByID(id uint) (*entity.Campaign, error)
	Create(campaign *entity.Campaign) error
}

type campaignRepository struct {
	db *gorm.DB
}

func NewCampaignRepository(db *gorm.DB) CampaignRepository {
	return &campaignRepository{db: db}
}

func (r *campaignRepository) FindAll() ([]entity.Campaign, error) {
	var campaigns []entity.Campaign
	// Preload dipakai kalau mau manggil relasi kategorinya sekaligus
	err := r.db.Preload("Categories").Find(&campaigns).Error
	return campaigns, err
}

func (r *campaignRepository) FindByID(id uint) (*entity.Campaign, error) {
	var campaign entity.Campaign
	err := r.db.Preload("Categories").First(&campaign, id).Error
	if err != nil {
		return nil, err
	}
	return &campaign, nil
}

func (r *campaignRepository) Create(campaign *entity.Campaign) error {
	return r.db.Create(campaign).Error
}

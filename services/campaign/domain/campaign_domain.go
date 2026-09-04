package domain

import (
	"context"
	"time"

	"campaign-service/entity"
)

type CampaignFilter struct {
	CategoryID   *uint
	Status       *string
	Limit        int
	Offset       int
	IncludeEnded bool
}

type CampaignInput struct {
	Title        string
	Description  string
	TargetAmount float64
	Deadline     time.Time
	CategoryIDs  []uint
}

type CampaignRepository interface {
	Create(ctx context.Context, c *entity.Campaign) error
	Update(ctx context.Context, c *entity.Campaign) error
	FindByID(ctx context.Context, id uint) (*entity.Campaign, error)
	FindAll(ctx context.Context, f CampaignFilter) ([]entity.Campaign, error)
	SoftDelete(ctx context.Context, id uint) error

	ReplaceCategories(ctx context.Context, campaignID uint, categoryIDs []uint) error

	IncrementCollected(ctx context.Context, id uint, amount float64) error
	SetCollected(ctx context.Context, id uint, amount float64) error
}

type CategoryRepository interface {
	FindAll(ctx context.Context) ([]entity.Category, error)
	FindByIDs(ctx context.Context, ids []uint) ([]entity.Category, error)
}

type TransactionClient interface {
	SumAccepted(ctx context.Context, campaignID uint) (float64, error)
	DonorIDs(ctx context.Context, campaignID uint) ([]uint, error)
}

type CampaignUsecase interface {
	Create(ctx context.Context, adminID uint, in CampaignInput) (*entity.Campaign, error)
	Update(ctx context.Context, id uint, in CampaignInput) (*entity.Campaign, error)
	Delete(ctx context.Context, id uint) error
	GetByID(ctx context.Context, id uint) (*entity.Campaign, error)
	List(ctx context.Context, f CampaignFilter) ([]entity.Campaign, error)

	AddCollected(ctx context.Context, id uint, amount float64) error
	Recalculate(ctx context.Context, id uint) (*entity.Campaign, error)
}

type CategoryUsecase interface {
	List(ctx context.Context) ([]entity.Category, error)
}

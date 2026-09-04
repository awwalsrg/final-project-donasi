package domain

import (
	"context"
	"time"

	"donation-service/entity"
)

type CreateTransactionInput struct {
	UserID      uint
	CampaignID  uint
	Name        string
	Amount      float64
	Message     string
	IsAnonymous bool
}

// PublicDonation is one row of a campaign's public donor list.
// DonorName is resolved through UserClient and replaced with a
// placeholder when the donor asked to stay anonymous.
type PublicDonation struct {
	DonorName string
	Amount    float64
	Message   string
	CreatedAt time.Time
}

type CampaignSnapshot struct {
	ID       uint
	Title    string
	Status   string
	IsActive bool
}

type TransactionRepository interface {
	Create(ctx context.Context, t *entity.Transaction) error
	FindByID(ctx context.Context, id uint) (*entity.Transaction, error)
	FindByUserID(ctx context.Context, userID uint, limit, offset int) ([]entity.Transaction, error)
	FindByStatus(ctx context.Context, status string, limit, offset int) ([]entity.Transaction, error)

	UpdateVerification(ctx context.Context, id uint, status string, verifiedBy uint) error

	SumAcceptedByCampaign(ctx context.Context, campaignID uint) (float64, error)
	FindDonorIDsByCampaign(ctx context.Context, campaignID uint) ([]uint, error)
	FindAcceptedByCampaign(ctx context.Context, campaignID uint, limit, offset int) ([]entity.Transaction, error)
}

type CampaignClient interface {
	GetCampaign(ctx context.Context, id uint) (*CampaignSnapshot, error)
	AddCollected(ctx context.Context, id uint, amount float64) error
}

// UserClient is the outbound port to user-service.
type UserClient interface {
	NamesByIDs(ctx context.Context, ids []uint) (map[uint]string, error)
}

type TransactionUsecase interface {
	Submit(ctx context.Context, in CreateTransactionInput) (*entity.Transaction, error)
	Accept(ctx context.Context, id, adminID uint) (*entity.Transaction, error)
	Reject(ctx context.Context, id, adminID uint) (*entity.Transaction, error)

	GetByID(ctx context.Context, id uint) (*entity.Transaction, error)
	ListByUser(ctx context.Context, userID uint, limit, offset int) ([]entity.Transaction, error)
	ListPending(ctx context.Context, limit, offset int) ([]entity.Transaction, error)

	SumAccepted(ctx context.Context, campaignID uint) (float64, error)
	DonorIDs(ctx context.Context, campaignID uint) ([]uint, error)
	ListCampaignDonors(ctx context.Context, campaignID uint, limit, offset int) ([]PublicDonation, error)
}

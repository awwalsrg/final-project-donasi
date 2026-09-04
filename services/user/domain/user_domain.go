package domain

import (
	"context"

	"user-service/entity"
)

type RegisterInput struct {
	Name        string
	Email       string
	Password    string
	CategoryIDs []uint
}

type UpdateProfileInput struct {
	Name              *string
	NotifyNewCampaign *bool
}

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	FindByID(ctx context.Context, id uint) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindAll(ctx context.Context, limit, offset int) ([]entity.User, error)

	ReplacePreferredCategories(ctx context.Context, userID uint, categoryIDs []uint) error
	FindPreferredCategoryIDs(ctx context.Context, userID uint) ([]uint, error)
	FindSubscribersByCategoryIDs(ctx context.Context, categoryIDs []uint) ([]entity.User, error)
}
type CategoryClient interface {
	ExistingCategoryIDs(ctx context.Context, ids []uint) ([]uint, error)
}

type UserUsecase interface {
	Register(ctx context.Context, in RegisterInput) (*entity.User, error)
	Login(ctx context.Context, email, password string) (token string, err error)
	GetByID(ctx context.Context, id uint) (*entity.User, error)
	List(ctx context.Context, limit, offset int) ([]entity.User, error)
	UpdateProfile(ctx context.Context, userID uint, in UpdateProfileInput) (*entity.User, error)

	GetPreferences(ctx context.Context, userID uint) ([]uint, error)
	SetPreferences(ctx context.Context, userID uint, categoryIDs []uint) error
	ListSubscribers(ctx context.Context, categoryIDs []uint) ([]entity.User, error)
}

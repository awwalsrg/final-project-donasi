package entity

const (
	RoleUser  = "user"
	RoleAdmin = "admin"
)

type User struct {
	ID                uint   `gorm:"primaryKey;column:id" json:"id"`
	Name              string `gorm:"column:name;not null" json:"name"`
	Email             string `gorm:"column:email;not null;unique" json:"email"`
	PasswordHash      string `gorm:"column:password_hash;not null" json:"-"`
	Role              string `gorm:"column:role;not null;default:'user'" json:"role"`
	NotifyNewCampaign bool   `gorm:"column:notify_new_campaign;not null;default:true" json:"notify_new_campaign"`
}

func (User) TableName() string { return "users" }

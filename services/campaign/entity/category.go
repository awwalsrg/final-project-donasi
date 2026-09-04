package entity

type Category struct {
	ID          uint   `gorm:"primaryKey;column:id" json:"id"`
	Name        string `gorm:"column:name;not null;unique" json:"name"`
	Description string `gorm:"column:description" json:"description"`
}

func (Category) TableName() string { return "categories" }

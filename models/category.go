package models

type Category struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"type:varchar(255)" json:"name"`
	Slug        string `gorm:"type:varchar(255)" json:"slug"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

package entity

import (
	"time"
)

// Todo entity
type Todo struct {
	ID          uint   `gorm:"primary_key;column:id;auto_increment:true"`
	Title       string `gorm:"column:title"`
	Description string `gorm:"column:description"`

	CreatedAt time.Time `gorm:"column:created_at;autocreatetime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autocupdatetime"`
}

package domain

import "time"

type Store struct {
	ID        string     `gorm:"type:uuid;default:gen_random_uuid();primaryKey;column:id"`
	Name      string     `gorm:"column:name;not null"`
	CreatedAt time.Time  `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt *time.Time `gorm:"column:updated_at;autoUpdateTime:false"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}

func (Store) TableName() string {
	return "stores"
}

package domain

import "time"

type Employee struct {
	ID           string     `gorm:"type:uuid;default:gen_random_uuid();primaryKey;column:id"`
	StoreID      string     `gorm:"type:uuid;not null;column:store_id"`
	Name         string     `gorm:"column:name;not null"`
	Email        string     `gorm:"column:email;not null"`
	PasswordHash string     `gorm:"column:password_hash;not null"`
	Role         string     `gorm:"column:role;not null"`
	IsActive     bool       `gorm:"column:is_active;not null;default:true"`
	CreatedAt    time.Time  `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    *time.Time `gorm:"column:updated_at;autoUpdateTime:false"`
	DeletedAt    *time.Time `gorm:"column:deleted_at"`
}

func (Employee) TableName() string {
	return "employees"
}

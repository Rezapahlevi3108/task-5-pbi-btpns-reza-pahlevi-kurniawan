package models

import "time"

type User struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"not null;unique" json:"username"`
	Email     string    `gorm:"not null;unique" json:"email"`
	Password  string    `gorm:"not null;size:255;min:6" json:"password"`
	Photos []Photo `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"photos"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
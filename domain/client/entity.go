package client

import "time"

type Client struct {
	ID        int    `gorm:"primary_key;auto_increment;not_null"`
	Fullname  string `gorm:"type:varchar(50);not null"`
	Email     string `gorm:"type:varchar(100);not null"`
	Address   string `gorm:"type:longtext;not null"`
	City      string `gorm:"type:varchar(50);not null"`
	ZipCode   string `gorm:"type:varchar(50);not null"`
	Company   string `gorm:"type:varchar(50);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

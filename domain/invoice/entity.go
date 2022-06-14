package invoice

import (
	"time"

	"github.com/kokolopo/capstone_alta/domain/client"
)

type Invoice struct {
	ID          int           `gorm:"primary_key;auto_increment;not_null"`
	ClientID    int           `gorm:"type:int(25);not null"`
	Client      client.Client `gorm:"foreignKey:ClientID;not null"`
	PaymentURL  string        `gorm:"type:varchar(100);not null"`
	InvoiceDate time.Time
	InvoiceDue  time.Time
}

type DetailInvoice struct {
	ID        int     `gorm:"primary_key;auto_increment;not_null"`
	InvoiceID int     `gorm:"type:int(25);not null"`
	Invoice   Invoice `gorm:"foreignKey:InvoiceID;not null"`
	ItemName  string  `gorm:"type:varchar(100);not null"`
	Price     int     `gorm:"type:int(100);not null"`
	Quantity  int     `gorm:"type:int(25);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

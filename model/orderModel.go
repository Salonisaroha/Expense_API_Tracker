package Models

import "time"

type Order struct {
	Id             int       `json:"Id" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;primary key"`
	CashierId      int       `json:"cashierId"`
	PaymentTypesId int       `json:"paymentTypesId"`
	TotalPrice     int       `json:"totalPrice"`
	TotalPaid      int       `json:"totalPaid"`
	TOtalReturn    int       `json:"totalReturn"`
	ReceiptId      string    `json:"receiptId"`
	IsDownload     int       `json:"is_download"`
	ProductId      string    `json:"productId"`
	Quantities     string    `json:"quantities"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

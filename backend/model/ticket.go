package model

import (
	"time"
)

type Ticket struct {
	Id 				string		`gorm:"type:uuid;primary_key" json:"id"`
	Customer_id 	string		`gorm:"type:uuid" json:"-"`
	Customer 		Customer	`gorm:"foreignkey:Id;association_foreignkey:customer_id" json:"customer"`
	Status 			string		`gorm:"type:requestStatus" json:"status"`
	Cargo_type		string		`gorm:"type:cargoType" json:"cargo_type"`
	Title			string		`gorm:"type:varchar(200)" json:"title"`
	Date_from		time.Time	`gorm:"type:DATE" json:"date_from"`
	Date_to			time.Time	`gorm:"type:DATE" json:"date_to"`
	Dest_from		string		`gorm:"type:uuid" json:"-"`
	Dest_from_s		Airport		`gorm:"foreignkey:Id;association_foreignkey:dest_from" json:"dest_from"`
	Dest_to			string		`gorm:"type:uuid" json:"-"`
	Dest_to_s		Airport		`gorm:"foreignkey:Id;association_foreignkey:dest_to" json:"dest_to"`
	Price			uint		`gorm:"type:bigint" json:"price"`
	Ticket_comment	string		`gorm:"type:varchar" json:"ticket_comment"`
}

func (Ticket) TableName() string {
    return "app_db_ticket"
}

package model

import (
	"time"
)

type Request struct {
	Id 					string		`gorm:"type:uuid;primary_key" json:"id"`
	Operator_id 		string		`gorm:"type:uuid" json:"-"`
	Operator 			Operator	`gorm:"foreignkey:Id;association_foreignkey:operator_id" json:"operator"`
	Status 				string		`gorm:"type:requestStatus" json:"status"`
	Price				uint		`gorm:"type:bigint" json:"price"`
	Pilot_id 			string		`gorm:"type:uuid" json:"-"`
	Pilot 				Pilot		`gorm:"foreignkey:Id;association_foreignkey:pilot_id" json:"pilot"`
	Required_license	string		`gorm:"type:licenceType" json:"required_license"`
	Required_visa       string     	`gorm:"type:visaType" json:"required_visa"`
	Deadline			time.Time	`gorm:"type:DATE" json:"deadline"`
	Request_type        uint        `gorm:"type:bigint" json:"request_type"`
	Request_comment		string		`gorm:"type:varchar" json:"request_comment"`
	Ticket_id			string		`gorm:"type:uuid" json:"-"`
	Ticket				Ticket		`gorm:"foreignkey:Id;association_foreignkey:ticket_id" json:"ticket"`
	Plane_id            string      `gorm:"type:uuid" json:"-"`
	Plane				Plane		`gorm:"foreignkey:Id;association_foreignkey:plane_id" json:"plane"`
}

func (Request) TableName() string {
    return "app_db_pilot_request"
}

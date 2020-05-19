package model

type Plane struct {
	Id 					string		`gorm:"type:uuid;primary_key" json:"id"`
	Name 				string		`gorm:"type:varchar(200)" json:"name"`
	Registration_prefix string  	`gorm:"type:varchar(7)" json:"registration_prefix"`
	Registration_id     string  	`gorm:"type:varchar(30)" json:"registration_id"`
	Plane_type     		string  	`gorm:"type:varchar(50)" json:"plane_type"`
	Current_location	string		`gorm:"type:uuid" json:"-"`
	Current_location_s	Airport		`gorm:"foreignkey:Id;association_foreignkey:current_location" json:"current_location"`
	//Operators 			[]Operator	`gorm:"many2many:app_db_operator_plane_bridge;association_foreignkey:plane_id;foreignkey:Id"`
}

func (Plane) TableName() string {
    return "app_db_plane"
}

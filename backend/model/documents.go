package model

type License struct {
	Id 					string		`gorm:"type:uuid;primary_key" json:"id"`
	Name 				string		`gorm:"type:varchar(200)" json:"name"`
	License_type 		string  	`gorm:"type:licenceType" json:"license_type"`
	Image     			uint  		`gorm:"type:OID" json:"image"`
	Image_size     		uint  		`gorm:"type:bigint" json:"image_size"`
	Is_active			bool		`gorm:"type:boolean" json:"is_active"`
	Pilot_id			string		`gorm:"type:uuid" json:"-"`
	Pilot				Pilot		`gorm:"foreignkey:Id;association_foreignkey:Pilot_id" json:"pilot"`
}

func (License) TableName() string {
    return "app_db_license"
}

type Visa struct {
	Id 					string		`gorm:"type:uuid;primary_key" json:"id"`
	Name 				string		`gorm:"type:varchar(200)" json:"name"`
	Visa_type 			string  	`gorm:"type:visaType" json:"visa_type"`
	Image     			uint  		`gorm:"type:OID" json:"image"`
	Image_size     		uint  		`gorm:"type:bigint" json:"image_size"`
	Is_active			bool		`gorm:"type:boolean" json:"is_active"`
	Pilot_id			string		`gorm:"type:uuid" json:"-"`
	Pilot				Pilot		`gorm:"foreignkey:Id;association_foreignkey:Pilot_id" json:"pilot"`
}

func (Visa) TableName() string {
    return "app_db_visa"
}

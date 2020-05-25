package model

type Country struct {
	Id 		uint	`gorm:"type:smallint;primary_key" json:"-"`
	Name 	string	`gorm:"type:varchar(50)" json:"name"`
	Iso 	string	`gorm:"type:varchar(2)" json:"iso"`
}

func (Country) TableName() string {
	return "app_db_country"
}

type City struct {
	Id 			uint	`gorm:"type:bigint;primary_key" json:"-"`
	Name 		string	`gorm:"type:varchar(50)" json:"name"`
	Latitude 	float32	`gorm:"type:float" json:"latitude"`
	Longitude 	float32	`gorm:"type:float" json:"longitude"`
	Country_id	uint	`gorm:"type:smallint" json:"-"`
	Country 	Country	`gorm:"foreignkey:Id;association_foreignkey:Country_id" json:"country"`
}

func (City) TableName() string {
	return "app_db_city"
}

type Airport struct {
	Id          string	`gorm:"type:uuid;primary_key" json:"-"`
	Type_		string	`gorm:"type:varchar(100);column:type" json:"type"`
    Name        string  `gorm:"type:varchar(200)" json:"name"`
    Latitude    float32	`gorm:"type:float" json:"latitude"`
    Longitude   float32 `gorm:"type:float" json:"longitude"`
	City_id     uint    `gorm:"type:bigint" json:"-"`
	City		City	`gorm:"foreignkey:Id;association_foreignkey:City_id" json:"city"`
}

func (Airport) TableName() string {
	return "app_db_airport"
}

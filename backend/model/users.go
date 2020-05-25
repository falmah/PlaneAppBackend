package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
	_ "fmt"
)

type User struct {
	Id			string		`gorm:"type:UUID;primary_key;default:CURRENT_TIMESTAMP" json:"id"`
	Name 		string 		`gorm:"type:varchar(200)" json:"name"`
	Surname 	string 		`gorm:"type:varchar(200)" json:"surname"`
	Phone 		string 		`gorm:"type:varchar(15)" json:"phone"`
	Email 		string 		`gorm:"type:varchar(50)" json:"email"`
	Created_at 	time.Time 	`gorm:"type:TIMESTAMP,not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	Role		string		`gorm:"type:userType" json:"role"`
	Password 	string		`gorm:"type:varchar(100)" json:"-"`
}

func (User) TableName() string {
	return "app_db_user"
}

func createUser(db *gorm.DB, user User) string {
	db.Create(&user)
	return user.Id
}

type Customer struct {
	Id 		string	`gorm:"type:UUID;primary_key;default:CURRENT_TIMESTAMP" json:"id"`
	User_id string  `gorm:"type:UUID" json:"-"`
	User 	User 	`gorm:"foreignkey:Id;association_foreignkey:user_id" json:"user"`
}

func (Customer) TableName() string {
    return "app_db_customer"
}

func (c *Customer) BeforeCreate(db *gorm.DB) {
	c.User_id = createUser(db, c.User)
}

type Operator struct {
	Id 				string	`gorm:"type:UUID;primary_key" json:"id"`
	Company_name 	string	`gorm:"type:varchar(200)" json:"company_name"`
	City_id			uint	`gorm:"type:bigint" json:"-"`
	City			City	`gorm:"foreignkey:Id;association_foreignkey:city_id" json:"city"`
	User_id 		string  `gorm:"type:UUID" json:"-"`
	User 			User 	`gorm:"foreignkey:Id;association_foreignkey:user_id" json:"user"`
}

func (Operator) TableName() string {
    return "app_db_operator"
}

func (o *Operator) BeforeCreate(db *gorm.DB) {	
	o.User_id = createUser(db, o.User)
}

type Pilot struct {
	Id 					string  `gorm:"type:UUID;primary_key" json:"id"`
	Busy 				bool	`gorm:"type:boolean" json:"busy"`
	Current_location	uint	`gorm:"type:bigint" json:"-"`
	City 				City	`gorm:"foreignkey:Id;association_foreignkey:current_location" json:"city"`
	User_id 			string  `gorm:"type:UUID" json:"-"`
	User 				User 	`gorm:"foreignkey:Id;association_foreignkey:user_id" json:"user"`
}

func (Pilot) TableName() string {
    return "app_db_pilot"
}

func (p *Pilot) BeforeCreate(db *gorm.DB) {
	p.User_id = createUser(db, p.User)
}

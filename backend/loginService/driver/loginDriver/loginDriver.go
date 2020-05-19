package loginDriver

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"encoding/json"
	"fmt"
	"model"
	log "github.com/sirupsen/logrus"
)

const connection_str = "user=docker password=docker host=172.18.0.2 dbname=app_db sslmode=disable"

func openConnection() *gorm.DB {
	db, err := gorm.Open("postgres", connection_str)
	if err != nil {
		db.Close()
		log.Error("cannot connect to the database")
		return nil
	}
	return db
}

func RegisterCustomer(cus *model.Customer) {
	out, _ := json.Marshal(cus)
	log.Infof("Customer creation: %s", string(out))
	
	db := openConnection()
	if db == nil {
		log.Error("Customer registration failed")
		return
	}

	if err := db.Create(cus).Error; err != nil {
		log.Errorf("cannot create customer %s", err.Error())
	}

	db.Close()
	return
}

func getCountry(id uint) model.Country {
	country := model.Country{}

	db := openConnection()
	if db == nil {
		log.Error("Customer registration failed")
		return country
	}

	if err := db.First(&country, model.Country{Id: id}).Error; err != nil {
		log.Errorf("cannot find country. id %u: %s", id, err.Error())
		return country
	}

	db.Close()

	return country
}

func GetCity(name string) model.City {
	city := model.City{}

	db, err := gorm.Open("postgres", connection_str)
	if err != nil {
		panic("cannot connect to the database")
	}
	defer db.Close()

	if err := db.First(&city, model.City{Name: name}).Error; err != nil {
		fmt.Println("cannot find city")
		return city
	}
	db.Close()

	city.Country = getCountry(city.Country_id)

	return city
}

func getCityById(id uint) model.City {
	city := model.City{}

	db, err := gorm.Open("postgres", connection_str)
	if err != nil {
		panic("cannot connect to the database")
	}
	defer db.Close()

	if err := db.First(&city, model.City{Id: id}).Error; err != nil {
		fmt.Println("cannot find city")
		return city
	}
	db.Close()

	city.Country = getCountry(city.Country_id)

	return city
}

func GetUser(email string, password string) model.User {
	user := model.User{}

	db, err := gorm.Open("postgres", connection_str)
	if err != nil {
		panic("cannot connect to the database")
	}
	defer db.Close()
	
	if err := db.First(&user, model.User{Email: email, Password: password}).Error; err != nil {
		fmt.Println("cannot find user")
		return user
	}
	db.Close()

	fmt.Println(user)

	return user
}

func GetCustomer(user model.User) model.Customer {
	customer := model.Customer{}
	
	db, err := gorm.Open("postgres", connection_str)
	if err != nil {
		panic("cannot connect to the database")
	}
	defer db.Close()

	if err := db.First(&customer, model.Customer{User_id: user.Id}).Error; err != nil {
		fmt.Println("cannot find customer")
		return customer
	}

	customer.User = user

	db.Close()

	return customer
}

func RegisterOperator(op *model.Operator) {

	db, err := gorm.Open("postgres", connection_str)
	if err != nil {
		panic("cannot connect to the database")
	}
	defer db.Close()

	db.Create(op)

	db.Close()
	return
}

func GetOperator(user model.User) model.Operator {
	operator := model.Operator{}
	
	db, err := gorm.Open("postgres", connection_str)
	if err != nil {
		panic("cannot connect to the database")
	}
	defer db.Close()

	if err := db.First(&operator, model.Operator{User_id: user.Id}).Error; err != nil {
		fmt.Println("cannot find operator")
		return operator
	}

	operator.User = user
	operator.City = getCityById(operator.City_id)

	db.Close()

	return operator
}

func RegisterPilot(pi *model.Pilot) {

	db, err := gorm.Open("postgres", connection_str)
	if err != nil {
		panic("cannot connect to the database")
	}
	defer db.Close()

	db.Create(pi)

	db.Close()
	return
}

func GetPilot(user model.User) model.Pilot {
	pilot := model.Pilot{}
	
	db, err := gorm.Open("postgres", connection_str)
	if err != nil {
		panic("cannot connect to the database")
	}
	defer db.Close()

	if err := db.First(&pilot, model.Pilot{User_id: user.Id}).Error; err != nil {
		fmt.Println("cannot find operator")
		return pilot
	}

	pilot.User = user
	pilot.City = getCityById(pilot.Current_location)

	db.Close()

	return pilot
}
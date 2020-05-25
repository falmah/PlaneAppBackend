package customerDriver

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "encoding/json"
	_ "fmt"
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

func checkCustomer(id string) bool {
	db := openConnection();
	if db == nil {
		log.Error("Customer check failed")
		return false
	}

	var cus model.Customer
	if db.First(&cus, "id = ?", id).RecordNotFound() {
		return false
	}
		
	return true
}

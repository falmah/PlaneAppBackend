package operatorDriver

import (
	_"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "encoding/json"
	"fmt"
	"model"
	log "github.com/sirupsen/logrus"
)

func checkOperator(id string) bool {
	db := openConnection();
	if db == nil {
		log.Error("Operator check failed")
		return false
	}

	var op model.Operator
	if db.First(&op, "id = ?", id).RecordNotFound() {
		return false
	}

	return true
}

func GetTickets(operator string) []model.Ticket {
	tickets := []model.Ticket{}

	if !checkOperator(operator) {
		return tickets
	}

	db := openConnection()
	if db == nil {
		log.Error("get ticket failed")
		return tickets
	}
	defer db.Close()

	if err := db.Find(&tickets).Error; err != nil {
		fmt.Println("cannot find tickets")
		return tickets
	}

	for i := range tickets {
		db.Preload("Customer.User").Preload("Dest_from_s").Preload("Dest_to_s").First(&tickets[i])
	}
	
	db.Close()

	return tickets
}

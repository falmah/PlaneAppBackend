package operatorDriver

import (
	_"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "encoding/json"
	"fmt"
	"model"
	log "github.com/sirupsen/logrus"
)

func CreateRequest(id string, request *model.Request) {
	if !checkOperator(id) {
		log.Errorf("Operator id %s does not exist", id)
		return 
	}
	db := openConnection()
	if db == nil {
		log.Error("create request failed")
		return
	}

	db.Create(request)

	db.Close()
	return
}

func GetRequest(operator string, request string) model.Request {
	request_o := model.Request{}

	if !checkOperator(operator) {
		log.Errorf("Customer id %s does not exist", operator)
		return request_o
	}

	db := openConnection()
	if db == nil {
		log.Error("get request failed")
		return request_o
	}
	defer db.Close()

	if err := db.First(&request_o, model.Request{Id: request}).Error; err != nil {
		fmt.Println("cannot find request")
		return request_o
	}

	db.Preload("Operator.User").
		Preload("Operator.City.Country").
		Preload("Pilot.User").
		Preload("Pilot.City.Country").
		Preload("Ticket.Dest_from_s").
		Preload("Ticket.Dest_to_s").
		Preload("Ticket.Customer.User").
		Preload("Plane.Current_location_s").
		First(&request_o)

	db.Close()

	return request_o
}

func UpdateRequest(operator string, request *model.Request) {
	if !checkOperator(operator) {
		log.Errorf("Operator id %s does not exist", operator)
		return 
	}

	db := openConnection()
	if db == nil {
		log.Error("update request failed")
		return
	}
	defer db.Close()
	fmt.Println(request)
	db.Save(request)
}

func DeleteRequest(operator string, request string) {
	request_o := GetRequest(operator, request)
	request_o.Status = "closed"
	UpdateRequest(operator, &request_o)
}

func GetRequests(operator string) []model.Request {
	requests := []model.Request{}

	db := openConnection()
	if db == nil {
		log.Error("get requests failed")
		return requests
	}
	defer db.Close()

	if err := db.Find(&requests).Error; err != nil {
		fmt.Println("cannot find requests")
		return requests
	}

	for i := range requests {
			db.Preload("Operator.User").
			Preload("Operator.City.Country").
			Preload("Pilot.User").
			Preload("Pilot.City.Country").
			Preload("Ticket.Dest_from_s").
			Preload("Ticket.Dest_to_s").
			Preload("Ticket.Customer.User").
			Preload("Plane.Current_location_s").
			First(&requests[i])							
	}
	
	db.Close()

	return requests
}
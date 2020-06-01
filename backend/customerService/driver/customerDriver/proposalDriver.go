package customerDriver

import (
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "encoding/json"
	_ "fmt"
	"model"
	log "github.com/sirupsen/logrus"
)

func GetProposals(id string) []model.Request {
	requests := []model.Request{}

	db := openConnection()
	if db == nil {
		log.Error("get proposals failed")
		return requests
	}
	defer db.Close()

	if err := db.Debug().Table("app_db_pilot_request").Select("app_db_pilot_request.*").
		Joins("LEFT JOIN app_db_ticket ON app_db_pilot_request.ticket_id = app_db_ticket.id").
		Where("app_db_pilot_request.status = 'pending' AND app_db_ticket.customer_id = ?", id).
		Find(&requests).Error; err != nil {
			log.Error("cannot find proposals")
		return requests
	}

	for i := range requests {
		db.Preload("Operator.User").
		Preload("Operator.City.Country").
		Preload("Pilot.User").
		Preload("Pilot.City.Country").
		Preload("Ticket.Dest_from_s.City.Country").
		Preload("Ticket.Dest_to_s.City.Country").
		Preload("Ticket.Customer.User").
		Preload("Plane.Current_location_s.City.Country").
		First(&requests[i])
	}
	
	db.Close()

	return requests
}

func ChangeProposalStatus(customer string, status string, proposal string) {
	db := openConnection()
	if db == nil {
		log.Error("get requests failed")
		return 
	}
	defer db.Close()

	var req model.Request
	if err := db.First(&req, model.Request{Id: proposal}).Error; err != nil {
		log.Error("cannot find request")
		return 
	}

	if status == "pending" {
		var tic model.Ticket
		db.Preload("Ticket").First(&req)
		if err := db.First(&tic, model.Ticket{Id: req.Ticket.Id}).Error; err != nil {
			log.Error("cannot find ticket")
			return 
		}
		tic.Status = status
		db.Save(&tic)
	}else if status == "rejected" {
		req.Status = status
		db.Save(&req)
	}
}

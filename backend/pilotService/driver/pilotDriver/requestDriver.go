package pilotDriver

import (
	_"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "encoding/json"
	"model"
	log "github.com/sirupsen/logrus"
)

func GetRequests(pilot string) []model.Request {
	requests := []model.Request{}

	db := openConnection()
	if db == nil {
		log.Error("get requests failed")
		return requests
	}
	defer db.Close()

	if err := db.Where("status != 'closed' AND status != 'rejected' AND pilot_id = ?", pilot).
						Find(&requests).Error; err != nil {
		log.Error("cannot find requests")
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

func ChangeRequestStatus(pilot string, status string, request string) {
	db := openConnection()
	if db == nil {
		log.Error("get requests failed")
		return 
	}
	defer db.Close()

	var req model.Request
	if err := db.First(&req, model.Request{Id: request}).Error; err != nil {
		log.Error("cannot find request")
		return 
	}

	req.Status = status
	db.Save(&req)
}
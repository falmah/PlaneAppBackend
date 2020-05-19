package customerDriver

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "encoding/json"
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

func CreateTicket(ticket *model.Ticket, id string) {
	if !checkCustomer(id) {
		log.Errorf("Customer id %s does not exist", id)
		return 
	}

	db := openConnection()
	if db == nil {
		log.Error("create ticket failed")
		return
	}

	db.Debug().Create(ticket)

	db.Close()
	return
}
/*{
    "status": "open",
    "cargo_type": "passenger",
    "title": "lalal",
    "date_from": "2020-03-01",
    "date_to": "2020-04-01",
    "dest_from": "ffe2232e-e0e9-41ce-b447-570ce8949671",
    "dest_to": "448ea469-d63c-4572-8dae-b11f21e2885c",
    "price": 1200,
    "ticket_comment": "qwerasdfxzvc"
}*/
func GetTicket(customer string, ticket string) model.Ticket {
	ticket_o := model.Ticket{}

	if !checkCustomer(customer) {
		log.Errorf("Customer id %s does not exist", customer)
		return ticket_o
	}

	db := openConnection()
	if db == nil {
		log.Error("get ticket failed")
		return ticket_o
	}
	defer db.Close()

	if err := db.First(&ticket_o, model.Ticket{Id: ticket}).Error; err != nil {
		fmt.Println("cannot find ticket")
		return ticket_o
	}

	db.Preload("Customer.User").Preload("Dest_from_s").Preload("Dest_to_s").First(&ticket_o)

	db.Close()

	return ticket_o
}

func UpdateTicket(customer string, ticket *model.Ticket) {
	if !checkCustomer(customer) {
		log.Errorf("Customer id %s does not exist", customer)
		return 
	}

	db := openConnection()
	if db == nil {
		log.Error("get ticket failed")
		return
	}
	defer db.Close()
	fmt.Println(ticket)
	db.Save(ticket)
}

func DeleteTicket(customer string, ticket string) {
	ticket_o := GetTicket(customer, ticket)
	ticket_o.Status = "closed"
	UpdateTicket(customer, &ticket_o)
}

func GetTickets(customer string) []model.Ticket {
	tickets := []model.Ticket{}

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
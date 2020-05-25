package customerDriver

import (
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "encoding/json"
	"fmt"
	"model"
	log "github.com/sirupsen/logrus"
)

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

	db.Create(ticket)

	db.Close()
	return
}

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

	db.Preload("Customer.User").Preload("Dest_from_s.City.Country").Preload("Dest_to_s.City.Country").First(&ticket_o)

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

	if err := db.Debug().Where("status != 'closed' AND customer_id = ?", customer).
						Find(&tickets).Error; err != nil {
		fmt.Println("cannot find tickets")
		return tickets
	}

	for i := range tickets {
		db.Preload("Customer.User").Preload("Dest_from_s.City.Country").Preload("Dest_to_s.City.Country").First(&tickets[i])
	}
	
	db.Close()

	return tickets
}
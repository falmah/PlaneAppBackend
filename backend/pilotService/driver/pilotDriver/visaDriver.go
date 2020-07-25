package pilotDriver

import (
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "encoding/json"
	_"fmt"
	"model"
	log "github.com/sirupsen/logrus"
)


func CreateVisa(id string, visa *model.Visa) {
	if !checkPilot(id) {
		log.Errorf("Pilot id %s does not exist", id)
		return 
	}
	db := openConnection()
	if db == nil {
		log.Error("create visa failed")
		return
	}
	visa.Image = CreateOid()
	visa.Image_size = 0
	db.Create(visa)
	db.Preload("Pilot.City.Country").First(visa)
	db.Close()
	return
}

func GetVisas(id string) []model.Visa {
	visas := []model.Visa{}

	db := openConnection()
	if db == nil {
		log.Error("get requests failed")
		return visas
	}
	defer db.Close()

	if err := db.Where("pilot_id = ?", id).Find(&visas).Error; err != nil {
		log.Error("cannot find visas")
		return visas
	}
	log.Info(visas)
	for i := range visas {
		db.Preload("Pilot.City.Country").First(&visas[i])
	}
	
	db.Close()

	return visas
}

func GetVisa(id string) model.Visa {

	var visa_o model.Visa
	db := openConnection()
	if db == nil {
		log.Error("get visa failed")
		return visa_o
	}
	defer db.Close()

	if err := db.First(&visa_o, model.Visa{Id: id}).Error; err != nil {
		log.Error("cannot find license")
		return visa_o
	}

	return visa_o
}

func UpdateVisaImageSize(id string, s int) {
	db := openConnection()
	if db == nil {
		log.Error("create request failed")
		return
	}
	defer db.Close()
	var vis model.Visa
	db.Where("id = ?", id).First(&vis)
	vis.Image_size = uint(s)
	db.Save(&vis)
}
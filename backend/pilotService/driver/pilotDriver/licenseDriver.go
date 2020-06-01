package pilotDriver

import (
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "encoding/json"
	_"fmt"
	"model"
	log "github.com/sirupsen/logrus"
)


func CreateLicense(id string, license *model.License) {
	if !checkPilot(id) {
		log.Errorf("Operator id %s does not exist", id)
		return 
	}
	db := openConnection()
	if db == nil {
		log.Error("create request failed")
		return
	}
	license.Image = CreateOid()
	license.Image_size = 0
	db.Create(license)
	db.Preload("Pilot.City.Country").First(license)
	db.Close()
	return
}

func GetLicenses(id string) []model.License {
	licenses := []model.License{}

	db := openConnection()
	if db == nil {
		log.Error("get requests failed")
		return licenses
	}
	defer db.Close()

	if err := db.Where("pilot_id = ?", id).Find(&licenses).Error; err != nil {
		log.Error("cannot find licenses")
		return licenses
	}
	log.Info(licenses)
	for i := range licenses {
		db.Preload("Pilot.City.Country").First(&licenses[i])
	}
	
	db.Close()

	return licenses
}

func GetLicense(id string) model.License {

	var license_o model.License
	db := openConnection()
	if db == nil {
		log.Error("create request failed")
		return  license_o
	}
	defer db.Close()

	if err := db.First(&license_o, model.License{Id: id}).Error; err != nil {
		log.Error("cannot find license")
		return license_o
	}

	return license_o
}

func UpdateImageSize(id string, s int) {
	db := openConnection()
	if db == nil {
		log.Error("create request failed")
		return
	}
	defer db.Close()
	var lis model.License
	db.Where("id = ?", id).First(&lis)
	lis.Image_size = uint(s)
	db.Save(&lis)
}
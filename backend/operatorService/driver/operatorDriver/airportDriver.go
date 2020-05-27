package operatorDriver

import (
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "encoding/json"
	_ "fmt"
	"model"
	log "github.com/sirupsen/logrus"
)

func GetAirportId(name string) string {
	airport := model.Airport{}

	db := openConnection()
	if db == nil {
		log.Error("get airport id failed")
		return ""
	}
	defer db.Close()

	if err := db.First(&airport, model.Airport{Name: name}).Error; err != nil {
		log.Error("cannot find airport")
		return ""
	}

	db.Close()

	return airport.Id
}
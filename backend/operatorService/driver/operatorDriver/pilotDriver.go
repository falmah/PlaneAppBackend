package operatorDriver

import (
 	_"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "encoding/json"
	"fmt"
	"model"
	log "github.com/sirupsen/logrus"
)

func GePilots(operator string) []model.Pilot {
	pilots := []model.Pilot{}

	db := openConnection()
	if db == nil {
		log.Error("get pilots failed")
		return pilots
	}
	defer db.Close()

	if err := db.Find(&pilots).Error; err != nil {
		fmt.Println("cannot find pilots")
		return pilots
	}

	for i := range pilots {
		db.Preload("User").Preload("City").First(&pilots[i])
	}
	
	db.Close()

	return pilots
}


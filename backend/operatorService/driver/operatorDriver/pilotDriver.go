package operatorDriver

import (
 	_"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "encoding/json"
	"fmt"
	"model"
	log "github.com/sirupsen/logrus"
)

func GetUser(name string) model.User {
	user := model.User{}

	db := openConnection()
	if db == nil {
		log.Error("get user failed")
		return user
	}
	defer db.Close()
	
	if err := db.First(&user, model.User{Name: name}).Error; err != nil {
		fmt.Println("cannot find user")
		return user
	}
	db.Close()

	fmt.Println(user)

	return user
}

func GetPilot(name string) string {
	pilot := model.Pilot{}
	
	pilot.User = GetUser(name)

	db := openConnection()
	if db == nil {
		log.Error("get pilot failed")
		return ""
	}
	defer db.Close()

	if err := db.First(&pilot, model.Pilot{User_id: pilot.User.Id}).Error; err != nil {
		fmt.Println("cannot find pilot")
		return ""
	}

	db.Close()

	return pilot.Id
}

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


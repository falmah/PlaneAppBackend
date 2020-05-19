package operatorDriver

import (
 	_"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "encoding/json"
	"fmt"
	"database/sql"
	"model"
	log "github.com/sirupsen/logrus"
)

func CreatePlane(operator string, plane *model.Plane) {

	if !checkOperator(operator) {
		log.Errorf("operator id %s does not exist", operator)
		return
	}
	fmt.Println("there")
	db := openConnection()
	if db == nil {
		log.Error("create plane failed")
		return
	}
	defer db.Close()

	db.Create(plane)
	
	db.Exec("INSERT INTO app_db_operator_plane_bridge (operator_id, plane_id) VALUES (?, ?)",operator, plane.Id)
	db.Close()
	return
}

func GetPlane(operator string, plane string) model.Plane {
	plane_o := model.Plane{}

	if !checkOperator(operator) {
		log.Errorf("Operator id %s does not exist", operator)
		return plane_o
	}

	db := openConnection()
	if db == nil {
		log.Error("get ticket failed")
		return plane_o
	}
	defer db.Close()

	if err := db.First(&plane_o, model.Plane{Id: plane}).Error; err != nil {
		fmt.Println("cannot find plane")
		return plane_o
	}

	db.Preload("Current_location_s").First(&plane_o)

	db.Close()

	return plane_o
} 

func GetPlanes(operator string) []model.Plane {

	var plane_ids []string
	var planes []model.Plane

	if !checkOperator(operator) {
		log.Errorf("operator id %s does not exist", operator)
		return planes
	}

	db, err := sql.Open("postgres", connection_str)
	if db == nil {
		log.Error("get planes failed")
		return planes
	}
	defer db.Close()
	
	rows, err := db.Query("SELECT plane_id FROM app_db_operator_plane_bridge WHERE operator_id = $1", operator)
	if err != nil {
		log.Error(err)
		return planes
	}
	defer rows.Close()

	for rows.Next() {
		var plane_id string
		err := rows.Scan(&plane_id)
		if err != nil {
			log.Error(err)
			return planes
		}
		plane_ids = append(plane_ids, plane_id)
	}
	db.Close()

	dbs := openConnection()
	if dbs == nil {
		log.Error("get planes failed")
		return planes
	}
	defer dbs.Close()

	for _ ,v := range plane_ids {
		var p model.Plane
		dbs.Preload("Current_location_s").First(&p, model.Plane{Id: v})
		planes = append(planes, p)
	}

	dbs.Close()
	return planes
}

func UpdatePlane(operator string, plane *model.Plane) {
	if !checkOperator(operator) {
		log.Errorf("Customer id %s does not exist", operator)
		return 
	}

	db := openConnection()
	if db == nil {
		log.Error("get ticket failed")
		return
	}
	defer db.Close()
	fmt.Println(plane)
	db.Save(plane)
}
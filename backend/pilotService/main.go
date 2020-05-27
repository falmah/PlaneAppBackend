package main

import (
	_"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/postgres"
	"app"
	log "github.com/sirupsen/logrus"
	)

func main () {

	log.SetFormatter(&log.TextFormatter{FullTimestamp: true,})
	log.Info("Login service create")
	app := &app.App{}
	app.Initialize()
	app.Run("0.0.0.0:3000")

}

package quark

import (
	"os"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

func (a *App) SetupDB() {

	connection := os.Getenv("DATABASE")
	gorm, err := gorm.Open("mysql", connection)
	if err != nil {
		log.Fatal("DB: Failed to connect database: ", err)
	}

    gorm.LogMode(true)
	a.DB = gorm
}

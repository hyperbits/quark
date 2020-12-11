package quark

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) LoadEnvironment() {
	// Load env settings
	err := godotenv.Load()
	if err != nil {
		log.Info("Could not find .env file, make sure you set global environment variables")
	}
}

func (a *App) Initialize() {
	a.LoadEnvironment()
	a.SetupLogger()
	a.SetupDB()
	a.SetupRouter()
}

func (a *App) Serve() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	address := fmt.Sprintf(":%s", port)

	log.Print("Serving on ", address)
	log.Fatal(http.ListenAndServe(address, a.Cors()))
}

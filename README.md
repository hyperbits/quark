# Quark

A go framework for building RESTful API's

## Setup

```
package main

import (
	"net/http"
	"github.com/mattfroese/quark"
	"github.com/mattfroese/quark/controller"
	"github.com/mattfroese/quark/response"
)

func main() {
	a := quark.App{}

    // Load environment, db and router
	a.Initialize()

    // Migrate Gorm models
	migrateDB(&a)

    // Setup routes
	routes(&a)

	a.Serve()
}

type Controller struct {
	App *quark.App
}

func (c *Controller) HealthCheck() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response.RespondWithJSON(w, http.StatusOK, map[string]string{"version": "1.0.0"})
	})
}

func routes(a *quark.App) {
	c := Controller{App: a}
	a.Get("/health", c.HealthCheck())
}

func migrateDB(a *quark.App) {
	a.DB.AutoMigrate(&models.User{})
}

```
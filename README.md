# Quark

A go framework for building RESTful API's

## Setup

```
package main

import (
	"net/http"

	"github.com/hyperbits/quark"
	"github.com/hyperbits/quark/response"
)

func main() {
	a := quark.App{}

	// Load environment, db and router
	a.LoadEnvironment()
	a.SetupLogger()
	//a.SetupDB()
	a.SetupRouter()

	// Migrate Gorm models
	// migrateDB(&a)

	// Setup routes
	routes(&a)

	a.Serve()
}

type Controller struct {
	App *quark.App
}

func (c *Controller) HelloWorld() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response.RespondWithJSON(w, http.StatusOK, map[string]string{"Hello": "World!"})
	})
}

func routes(a *quark.App) {
	c := Controller{App: a}
	a.Get("/", c.HelloWorld())
}

```

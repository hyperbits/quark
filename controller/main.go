package controller

import (
	"net/http"
	"quark/quark"
	"quark/response"
)

type Controller struct {
	App *quark.App
}

func (c *Controller) HealthCheck() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response.RespondWithJSON(w, http.StatusOK, map[string]string{"version": "1.0.0"})
	})
}

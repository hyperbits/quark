package quark

import (
	"net/http"
	"os"
	"strings"

	"github.com/rs/cors"
)

func (a *App) Cors() http.Handler {
	var origins = []string{""}
	if os.Getenv("CORS_ALLOWED_ORIGINS") != "" {
		origins = strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ",")
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   origins,
		AllowCredentials: true,
		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application
		},
		AllowedMethods: []string{
			"POST", "GET", "OPTIONS", "PUT", "DELETE",
		},
		Debug: os.Getenv("ENV") == "Dev",
	})

	handler := c.Handler(a.Router)

	return handler
}

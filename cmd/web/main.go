package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/salindae25/go-booking/pkg/config"
	"github.com/salindae25/go-booking/pkg/handlers"
	"github.com/salindae25/go-booking/pkg/renders"
)

var portNumber = ":8081"

var (
	app     config.AppConfig
	session *scs.SessionManager
)

func main() {
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := renders.CreateCachedTemplate()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false
	renders.NewTemplates(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	srv := http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}

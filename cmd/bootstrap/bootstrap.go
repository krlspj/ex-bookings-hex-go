package bootstrap

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/krlspj/ex-bookings-hex-go/internal/config"
	"github.com/krlspj/ex-bookings-hex-go/internal/handlers"
	mid "github.com/krlspj/ex-bookings-hex-go/internal/middleware"
	render_service "github.com/krlspj/ex-bookings-hex-go/internal/render/service"
	"github.com/krlspj/ex-bookings-hex-go/internal/server"
)

func Run() error {
	var app config.AppConfig

	// Change this to true when in production
	app.InProduction = false

	session := scs.New()
	session.Lifetime = 24 * time.Hour
	// Persiste de cookie after browser window is closed
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session
	app.UseCache = false

	rs := render_service.NewTemplates(&app)
	as := mid.NewMiddleware(&app)

	tc, err := rs.CreateTemplateCache()
	if err != nil {
		log.Fatal("[ERROR] Cannot create template cache:", err)
	}
	app.TemplateCache = tc

	hr := handlers.NewHanldersRepo(&app, rs)

	ctx := context.TODO()
	s := server.NewServer(ctx, "localhost", 60002, hr, as)

	return s.Run(ctx)
}

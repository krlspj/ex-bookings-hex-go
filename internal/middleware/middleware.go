package mid

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
	"github.com/krlspj/ex-bookings-hex-go/internal/config"
)

type MiddlewareService interface {
	WritetoConsole(next http.Handler) http.Handler
	NoSrurf(next http.Handler) http.Handler
	SessionLoad(next http.Handler) http.Handler
}

type DefMiddlewareService struct {
	app *config.AppConfig
}

func NewMiddleware(a *config.AppConfig) *DefMiddlewareService {
	return &DefMiddlewareService{
		app: a,
	}
}

func (s *DefMiddlewareService) WritetoConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page by:", r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

// NoSurf adds CSRF protection to all POST requests
func (s *DefMiddlewareService) NoSrurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   s.app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad loads and saves the session on every request
func (s *DefMiddlewareService) SessionLoad(next http.Handler) http.Handler {
	return s.app.Session.LoadAndSave(next)
}

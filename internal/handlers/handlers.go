package handlers

import (
	"log"
	"net/http"

	"github.com/krlspj/ex-bookings-hex-go/internal/config"
	"github.com/krlspj/ex-bookings-hex-go/internal/render/domain"
	render_service "github.com/krlspj/ex-bookings-hex-go/internal/render/service"
)

type HandlerRepo struct {
	app *config.AppConfig
	rs  render_service.RenderService
}

func NewHanldersRepo(a *config.AppConfig, render render_service.RenderService) *HandlerRepo {
	return &HandlerRepo{
		app: a,
		rs:  render,
	}
}

// Home is the home page handler
func (h *HandlerRepo) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	h.app.Session.Put(r.Context(), "remote_ip", remoteIP)

	h.rs.RenderTemplate(w, "home.page.tmpl", &domain.TemplateData{})
}

// About is the about page handler
func (h *HandlerRepo) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "hello again..."

	remoteIP := h.app.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	h.rs.RenderTemplate(w, "about.page.tmpl", &domain.TemplateData{
		StringMap: stringMap,
	})
}

func (h *HandlerRepo) Index(w http.ResponseWriter, r *http.Request) {
	err := h.rs.RenderHtml(w, "index.html")
	if err != nil {
		log.Println("[ERROR] rendering template index.html:", err.Error())
	}

}

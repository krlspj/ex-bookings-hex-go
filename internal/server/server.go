package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/krlspj/ex-bookings-hex-go/internal/handlers"
	mid "github.com/krlspj/ex-bookings-hex-go/internal/middleware"
)

type Server struct {
	httpAddr string
	mux      http.Handler
	h        *handlers.HandlerRepo
	ms       mid.MiddlewareService

	//mux      *pat.PatternServeMux
	//mux      *http.ServeMux
}

func NewServer(ctx context.Context, host string, port uint,
	hr *handlers.HandlerRepo,
	m mid.MiddlewareService) Server { //(context.Context, Server) {
	//func New(ctx context.Context, host string, port uint, router http.Handler, hr *handlers.HandlerRepo) Server { //(context.Context, Server) {
	srv := Server{
		httpAddr: fmt.Sprintf(host + ":" + fmt.Sprint(port)),
		//mux:      http.NewServeMux(),
		//mux: pat.New(),

		// Handlers
		h:  hr,
		ms: m,
	}

	//return serverContext(ctx), srv
	return srv

}

func (s *Server) Run(ctx context.Context) error {
	log.Printf("Listening on %s\n", s.httpAddr)

	s.mux = s.registerRoutes()

	return http.ListenAndServe(s.httpAddr, s.mux)
}

// registerRoutes add routes and middlewares
func (s *Server) registerRoutes() http.Handler {
	//mux := pat.New()
	//mux.Get("/", http.HandlerFunc(s.h.Home))
	//mux.Get("/about", http.HandlerFunc(s.h.About))
	//return mux
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(s.ms.WritetoConsole)
	mux.Use(s.ms.NoSrurf)
	mux.Use(s.ms.SessionLoad)

	// routes
	mux.Get("/", s.h.Home)
	mux.Get("/about", s.h.About)
	mux.Get("/index", s.h.Index)

	// file server for static files
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}

//func (s *Server) registerRoutes() {
//	fmt.Println("registring routes...")
//	s.mux.HandleFunc("/", s.h.Home)
//	s.mux.HandleFunc("/about", s.h.About)
//}

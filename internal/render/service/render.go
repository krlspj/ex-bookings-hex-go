package render_service

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/krlspj/ex-bookings-hex-go/internal/config"
	"github.com/krlspj/ex-bookings-hex-go/internal/render/domain"
)

//var functions_0 = template.FuncMap{}

type RenderService interface {
	RenderTemplate(w http.ResponseWriter, tmpl string, td *domain.TemplateData)
	CreateTemplateCache() (map[string]*template.Template, error)
}

type tmplService struct {
	app       *config.AppConfig
	functions template.FuncMap
}

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) *tmplService {
	return &tmplService{
		app:       a,
		functions: template.FuncMap{},
	}
}

func (s *tmplService) RenderTemplate(w http.ResponseWriter, tmpl string, td *domain.TemplateData) {
	var tc map[string]*template.Template
	if s.app.UseCache {
		tc = s.app.TemplateCache
	} else {
		tc, _ = s.CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Template not found")
	}

	buff := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buff, td)

	_, err := buff.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser:", err.Error())
	}

	//	parsedTemplate, err := template.ParseFiles("./templates/" + tmpl)
	//
	//	if err != nil {
	//		log.Println("[ERROR] [renderTemplate] ParseFiles error:", err.Error())
	//		return
	//	}
	//	err = parsedTemplate.Execute(w, nil)
	//	if err != nil {
	//		log.Println("[ERROR] [renderTemplate] Execute error:", err.Error())
	//		return
	//	}

}

// CreateTemplateCache creates a template cache as a map
func (s *tmplService) CreateTemplateCache() (map[string]*template.Template, error) {
	//myCache := map[string]*template.Template{}
	myCache := make(map[string]*template.Template, 0)

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(s.functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}

	return myCache, nil
}

/// Old render template function
//func RenderTemplate(w http.ResponseWriter, tmpl string) {
//	parsedTemplate, err := template.ParseFiles("./templates/" + tmpl)
//	if err != nil {
//		log.Println("[ERROR] [renderTemplate] ParseFiles error:", err.Error())
//		return
//	}
//	err = parsedTemplate.Execute(w, nil)
//	if err != nil {
//		log.Println("[ERROR] [renderTemplate] Execute error:", err.Error())
//		return
//	}
//
//}

func AddDefaultData(td *domain.TemplateData) *domain.TemplateData {

	return td
}

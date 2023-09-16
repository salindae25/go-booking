package renders

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/salindae25/go-booking/pkg/config"
	"github.com/salindae25/go-booking/pkg/models"
)

var app *config.AppConfig

// NewTemplates set application config inside render package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate render the template based on the template name
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var templateSet map[string]*template.Template
	if app.UseCache {
		templateSet = app.TemplateCache
	} else {
		var err error
		templateSet, err = CreateCachedTemplate()
		if err != nil {
			log.Println(err)
			return
		}
	}
	templateFileName := fmt.Sprintf("%v.page.tpl", tmpl)
	t, ok := templateSet[templateFileName]
	if !ok {
		log.Println("can not find the template", templateFileName)
		return
	}
	buf := new(bytes.Buffer)
	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

// CreateCachedTemplate create a template cache with layout attach to page template
func CreateCachedTemplate() (map[string]*template.Template, error) {
	tempCach := map[string]*template.Template{}
	templatePages, err := filepath.Glob("./templates/*.page.tpl")
	if err != nil {
		return tempCach, err
	}

	templateLayouts, err := filepath.Glob("./templates/*.layout.tpl")
	if err != nil {
		return tempCach, err
	}
	for _, page := range templatePages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return tempCach, err
		}
		if len(templateLayouts) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tpl")
			if err != nil {
				return tempCach, err
			}

		}

		tempCach[name] = ts
	}
	return tempCach, nil
}

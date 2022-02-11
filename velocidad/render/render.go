package render

import (
	"fmt"
	"github.com/CloudyKit/jet/v6"
	"html/template"
	"net/http"
	"strings"
)

type Render struct {
	Renderer   string
	RootPath   string
	Secure     bool
	Port       string
	ServerName string
	JetViews   *jet.Set
}

type TemplateData struct {
	IsAuthenticated bool
	IntMap          map[string]int
	StringMap       map[string]string
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	Port            string
	ServerName      string
	Secure          bool
}

func (v *Render) Page(w http.ResponseWriter, r http.Request, view string, variables, data interface{}) error {
	switch strings.ToLower(v.Renderer) {
	case "go":
		return v.GoPage(w, r, view, data)
	case "jet":

	}
	return nil
}

func (v *Render) GoPage(w http.ResponseWriter, r http.Request, view string, data interface{}) error {
	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/views/%s.page.tmpl", v.RootPath, view))
	if err != nil {
		return err
	}

	td := &TemplateData{}
	if data != nil {
		td = data.(*TemplateData)
	}

	err = tmpl.Execute(w, &td)
	if err != nil {
		return err
	}

	return nil
}

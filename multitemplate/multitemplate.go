package multitemplate

// This part was taken from https://github.com/gin-contrib/multitemplate/blob/master/multitemplate.go

import (
	"fmt"
	"html/template"
)

type Multitemplate map[string]*template.Template

func New() Multitemplate {
	return make(Multitemplate)
}

// Adds a new template
func (m Multitemplate) Add(name string, tmpl *template.Template) {
	if tmpl == nil {
		panic("template can not be nil")
	}
	if len(name) == 0 {
		panic("template name cannot be empty")
	}
	if _, ok := m[name]; ok {
		panic(fmt.Sprintf("template %s already exists", name))
	}
	m[name] = tmpl
}

// AddFromFiles supply add template from files
func (m Multitemplate) AddFromFiles(name string, files ...string) *template.Template {
	tmpl := template.Must(template.ParseFiles(files...))
	m.Add(name, tmpl)
	return tmpl
}

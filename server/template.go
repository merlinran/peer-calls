package server

import (
	"fmt"
	"html/template"
	"path/filepath"
	"strings"
)

type Templates map[string]*template.Template

func (t Templates) Get(name string) (tpl *template.Template, ok bool) {
	tpl, ok = t[name]

	return
}

func ParseTemplates(rootPath string) Templates {
	templates := Templates{}
	partials, _ := filepath.Glob(fmt.Sprintf("%s/_*.html", rootPath))
	all, _ := filepath.Glob(fmt.Sprintf("%s/*.html", rootPath))
	fileWithPartials := make([]string, len(partials)+1)
	copy(fileWithPartials[1:], partials)
	for _, filename := range all {
		fileWithPartials[0] = filename
		basename := filepath.Base(filename)
		if strings.HasPrefix(basename, "_") { // exclude partials
			continue
		}
		t, err := template.ParseFiles(fileWithPartials...)
		template.Must(t, err)
		templates[basename] = t

	}

	return templates
}

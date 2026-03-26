package lucide

import (
	"embed"
	"html/template"
	"io/fs"
)

//go:embed *.html *.templ
var templates embed.FS

// AddIcons combines the lucide icons with the given template.
func AddIcons(tpls *template.Template) error {
	_, err := tpls.ParseFS(templates, "*.html")
	return err
}

// GetHTMLFile returns the contents of the given html icon file.
func GetHTMLFile(name string) ([]byte, error) {
	return fs.ReadFile(templates, name+".html")
}

// GetTemplFile returns the contents of the given templ icon file.
func GetTemplFile(name string) ([]byte, error) {
	return fs.ReadFile(templates, name+".templ")
}

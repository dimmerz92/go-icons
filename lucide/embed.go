package lucide

import (
	"embed"
	"html/template"
)

//go:embed *.html *.templ
var Templates embed.FS

// AddIcons combines the lucide icons with the given template.
func AddIcons(tpls *template.Template) error {
	_, err := tpls.ParseFS(Templates, "*.html")
	return err
}

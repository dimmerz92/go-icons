package internal

import (
	"embed"
	"io/fs"

	"github.com/dimmerz92/go-icons/lucide"
)

// key=library, value[0]=icons path in cloned repo, value[1]=path in current repo
var SupportedLibraries = map[string][]string{
	"lucide": {"./lucide-repo/icons", "./lucide"},
}

var SupportedFormats = map[string]struct{}{
	"html":  {},
	"templ": {},
}

var IconEmbeds = map[string]embed.FS{
	"lucide": lucide.Templates,
}

// GetHTMLFile returns the contents of the given html icon file.
func GetHTMLFile(library, icon string) ([]byte, error) {
	return fs.ReadFile(IconEmbeds[library], icon+".html")
}

// GetTemplFile returns the contents of the given templ icon file.
func GetTemplFile(library, icon string) ([]byte, error) {
	return fs.ReadFile(IconEmbeds[library], icon+".templ")
}

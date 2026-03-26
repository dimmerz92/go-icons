package internal

import (
	"embed"
	"io/fs"

	"github.com/dimmerz92/go-icons/lucide"
	simpleicons "github.com/dimmerz92/go-icons/simple-icons"
)

// key=library, value[0]=icons path in cloned repo, value[1]=path in current repo
var SupportedLibraries = map[string][]string{
	"lucide":       {"./lucide-repo/icons", "./lucide"},
	"simple-icons": {"./simple-icons-repo/icons", "./simple-icons"},
}

var SupportedFormats = map[string]struct{}{
	"html":  {},
	"templ": {},
}

var IconEmbeds = map[string]embed.FS{
	"lucide":       lucide.Templates,
	"simple-icons": simpleicons.Templates,
}

// GetHTMLFile returns the contents of the given html icon file.
func GetHTMLFile(library, icon string) ([]byte, error) {
	return fs.ReadFile(IconEmbeds[library], icon+".html")
}

// GetTemplFile returns the contents of the given templ icon file.
func GetTemplFile(library, icon string) ([]byte, error) {
	return fs.ReadFile(IconEmbeds[library], icon+".templ")
}

package internal

import (
	"embed"
	"io/fs"

	"github.com/dimmerz92/go-icons/ionicons"
	"github.com/dimmerz92/go-icons/lucide"
	"github.com/dimmerz92/go-icons/material"
	radixicons "github.com/dimmerz92/go-icons/radix-icons"
	simpleicons "github.com/dimmerz92/go-icons/simple-icons"
)

// key=library, value[0]=icons path in cloned repo, value[1]=path in current repo
var SupportedLibraries = map[string][]string{
	"lucide":       {"./lucide-repo/icons", "./lucide"},
	"simple-icons": {"./simple-icons-repo/icons", "./simple-icons"},
	"radix-icons":  {"./radix-icons-repo/packages/radix-icons/icons", "./radix-icons"},
	"ionicons":     {"./ionicons-repo/src/svg", "./ionicons"},
	"material":     {"./material-repo/packages/mui-icons-material/material-icons", "./material"},
}

var SupportedFormats = map[string]struct{}{
	"html":  {},
	"templ": {},
}

var IconEmbeds = map[string]embed.FS{
	"lucide":       lucide.Templates,
	"simple-icons": simpleicons.Templates,
	"radix-icons":  radixicons.Templates,
	"ionicons":     ionicons.Templates,
	"material":     material.Templates,
}

// GetHTMLFile returns the contents of the given html icon file.
func GetHTMLFile(library, icon string) ([]byte, error) {
	return fs.ReadFile(IconEmbeds[library], icon+".html")
}

// GetTemplFile returns the contents of the given templ icon file.
func GetTemplFile(library, icon string) ([]byte, error) {
	return fs.ReadFile(IconEmbeds[library], icon+".templ")
}

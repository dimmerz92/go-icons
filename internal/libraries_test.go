package internal_test

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/dimmerz92/go-icons/internal"
)

func TestGetHTMLFile(t *testing.T) {
	for library := range internal.SupportedLibraries {
		t.Run(library, func(t *testing.T) {
			icons, err := internal.IconEmbeds[library].ReadDir(".")
			if err != nil {
				t.Fatalf("failed to list icons: %v", err)
			}

			var icon string
			for _, file := range icons {
				if filepath.Ext(file.Name()) == ".html" {
					icon = strings.ReplaceAll(file.Name(), ".html", "")
				}
			}

			data, err := internal.GetHTMLFile(library, icon)
			if err != nil {
				t.Fatalf("failed to get html file: %v", err)
			}

			if len(data) == 0 {
				t.Fatal("returned empty data")
			}
		})
	}
}

func TestGetTemplFile(t *testing.T) {
	for library := range internal.SupportedLibraries {
		t.Run(library, func(t *testing.T) {
			icons, err := internal.IconEmbeds[library].ReadDir(".")
			if err != nil {
				t.Fatalf("failed to list icons: %v", err)
			}

			var icon string
			for _, file := range icons {
				if filepath.Ext(file.Name()) == ".templ" {
					icon = strings.ReplaceAll(file.Name(), ".templ", "")
				}
			}

			data, err := internal.GetTemplFile(library, icon)
			if err != nil {
				t.Fatalf("failed to get templ file: %v", err)
			}

			if len(data) == 0 {
				t.Fatal("returned empty data")
			}
		})
	}
}

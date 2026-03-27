package internal_test

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/dimmerz92/go-icons/internal"
)

func TestFileSet(t *testing.T) {
	dir := t.TempDir()

	t.Run("test icon names", func(t *testing.T) {
		files := []string{
			"icon1.svg",
			"icon2.svg",
			"not_icon.txt",
			".hidden.svg",
			"icon3.SVG",
		}

		for _, file := range files {
			err := os.WriteFile(filepath.Join(dir, file), []byte("test"), 0600)
			if err != nil {
				t.Fatalf("failed to create file %s: %v", file, err)
			}
		}

		want := map[string]struct{}{
			"icon1":   {},
			"icon2":   {},
			".hidden": {},
			"icon3":   {},
		}

		got := internal.FileSet(dir, ".svg")
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %#v; want %#v", got, want)
		}
	})

	t.Run("test nested directories", func(t *testing.T) {
		err := os.Mkdir(filepath.Join(dir, "subdir"), 0755)
		if err != nil {
			t.Fatalf("failed to create subdir: %v", err)
		}
	})

	t.Run("test invalid path", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("IconSet did not panic on invalid path")
			}
		}()

		_ = internal.FileSet("/invalid/path/to/icons", ".svg")
	})
}

func TestFileSetDiff(t *testing.T) {
	setA := map[string]struct{}{
		"item1": {},
		"item2": {},
		"item3": {},
		"item4": {},
	}

	setB := map[string]struct{}{
		"item1": {},
		"item4": {},
	}

	want := map[string]struct{}{
		"item2": {},
		"item3": {},
	}

	got := internal.DiffFileSet(setA, setB)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v; want %#v", got, want)
	}
}

func TestKebabToPascal(t *testing.T) {
	tests := map[string]string{
		"i-am-a-test-string": "IAmATestString",
		"imnotakebab":        "Imnotakebab",
		"":                   "",
	}

	for kebab, pascal := range tests {
		got, err := internal.KebabToPascal(kebab)
		if err != nil {
			t.Fatal(err)
		}

		if got != pascal {
			t.Fatalf("got %s, wanted %s", got, pascal)
		}
	}
}

func TestToHTML(t *testing.T) {
	test := `<svg>
  <g></g>
</svg>
`
	want := `{{ define "my-icon" }}
<svg
  {{ range $value := . }}
    {{ $value }}
  {{ end }}
>
  <g></g>
</svg>
{{ end }}`

	tmp := t.TempDir()

	err := internal.ToHTML("my-icon", test, tmp)
	if err != nil {
		t.Fatalf("failed to write html icon: %v", err)
	}

	got, err := os.ReadFile(filepath.Join(tmp, "my-icon.html"))
	if err != nil {
		t.Fatalf("failed to read html icon: %v", err)
	}

	if string(got) != want {
		t.Fatalf("wanted\n%s\n\ngot\n%s", want, string(got))
	}
}

func TestToTempl(t *testing.T) {
	test := `<svg>
  <g></g>
</svg>
`
	want := `package lucide

templ MyIcon(attrs ...templ.Attributes) {
	<svg
		if len(attrs) > 0 {
			{ attrs[0]... }
		}
	>
		<g></g>
</svg>
}`

	tmp := t.TempDir()

	err := internal.ToTempl("lucide", "my-icon", test, tmp)
	if err != nil {
		t.Fatalf("failed to write templ icon: %v", err)
	}

	got, err := os.ReadFile(filepath.Join(tmp, "my-icon.templ"))
	if err != nil {
		t.Fatalf("failed to read templ icon: %v", err)
	}

	if string(got) != want {
		t.Fatalf("wanted\n%s\n\ngot\n%s", want, string(got))
	}
}

func TestNonAlphaPrefixer(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected string
	}{
		{name: "empty string"},
		{name: "no prefix", value: "test", expected: "test"},
		{name: "prefix from numeric", value: "0test", expected: "T0test"},
		{name: "prefix from symbol", value: "-test", expected: "T-test"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := internal.NonAlphaPrefixer(test.value); got != test.expected {
				t.Errorf("expected %s, got %s", test.expected, got)
			}
		})
	}
}

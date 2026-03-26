package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const HTML_TEMPLATE = `{{ define "%s" }}
%s
  {{ range $value := . }}
    {{ $value }}
  {{ end }}
>
  %s
{{ end }}`

const TEMPL_TEMPLATE = `package icons

templ %s(attrs ...templ.Attributes) {
	%s
		if len(attrs) > 0 {
			{ attrs[0]... }
		}
	>
		%s
}`

// FileSet returns a set of names of files of type 'ext' from the specified 'path' with 'ext' trimmed.
func FileSet(path, ext string) map[string]struct{} {
	entries, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	icons := make(map[string]struct{})
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		entryExt := strings.ToLower(filepath.Ext(entry.Name()))
		if entryExt == strings.ToLower(ext) {
			name := strings.TrimSuffix(strings.ToLower(entry.Name()), entryExt)
			icons[name] = struct{}{}
		}
	}

	return icons
}

// DiffSet returns the difference set of `setA` minus `setB`.
func DiffFileSet(setA, setB map[string]struct{}) map[string]struct{} {
	diff := make(map[string]struct{})
	for file := range setA {
		if _, ok := setB[file]; !ok {
			diff[file] = struct{}{}
		}
	}

	return diff
}

// KebabToPascal converts a kebab case string to a pascal case string.
func KebabToPascal(v string) (string, error) {
	var b strings.Builder
	caser := cases.Title(language.English)

	for part := range strings.SplitSeq(v, "-") {
		_, err := b.WriteString(caser.String(part))
		if err != nil {
			return "", err
		}
	}

	return b.String(), nil
}

// ToHTML embeds the 'svg' into a html template and saves it as 'name' to the 'outputPath'.
func ToHTML(name, svg, outputPath string) error {
	svgParts := strings.SplitN(svg, ">", 2)
	if len(svgParts) != 2 {
		return fmt.Errorf("malformed svg file")
	}

	template := fmt.Appendf([]byte{}, HTML_TEMPLATE, name, strings.TrimSpace(svgParts[0]), strings.TrimSpace(svgParts[1]))

	return os.WriteFile(filepath.Join(outputPath, name+".html"), template, 0600)
}

// ToTempl embeds the 'svg' into a templ template and saves it as 'name' to the 'outputPath'.
func ToTempl(name, svg, outputPath string) error {
	svgParts := strings.SplitN(svg, ">", 2)
	if len(svgParts) != 2 {
		return fmt.Errorf("malformed svg file")
	}

	fname, err := KebabToPascal(name)
	if err != nil {
		return err
	}

	template := fmt.Appendf([]byte{}, TEMPL_TEMPLATE, fname, strings.TrimSpace(svgParts[0]), strings.TrimSpace(svgParts[1]))

	return os.WriteFile(filepath.Join(outputPath, name+".templ"), template, 0600)
}

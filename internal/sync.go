package internal

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

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

// NonAlphaPrefixer prefixes the given string with T if it starts with a non alphanumeric character.
func NonAlphaPrefixer(v string) string {
	if len(v) < 1 || unicode.IsLetter(rune(v[0])) {
		return v
	}

	v = "T" + v

	return v
}

// ToHTML embeds the 'svg' into a html template and saves it as 'name' to the 'outputPath'.
func ToHTML(name string, svg []byte, outputPath string) error {
	comment, attrs, inner, err := parseSVG(svg)
	if err != nil {
		panic(err)
	}

	var b strings.Builder
	fmt.Fprintf(&b, "{{ define \"%s\" }}\n", name)
	if comment != "" {
		fmt.Fprintf(&b, "<!--\n%s\n-->\n", comment)
	}
	b.WriteString("<svg")

	for _, a := range attrs {
		fmt.Fprintf(&b, " %s=\"%s\"", a.Name.Local, escapeAttr(a.Value))
	}

	b.WriteString("\n\t{{ range $value := . }}\n")
	b.WriteString("\t\t{{ $value }}\n")
	b.WriteString("\t{{ end }}\n")
	b.WriteString(">\n\t")
	b.WriteString(strings.TrimSpace(inner))
	b.WriteString("\n</svg>\n")
	b.WriteString("{{ end }}\n")

	return os.WriteFile(filepath.Join(outputPath, name+".html"), []byte(b.String()), 0644)
}

// ToTempl embeds the 'svg' into a templ template and saves it as 'name' to the 'outputPath'.
func ToTempl(library, name string, svg []byte, outputPath string) error {
	comment, attrs, inner, err := parseSVG(svg)
	if err != nil {
		panic(err)
	}

	fname, err := KebabToPascal(NonAlphaPrefixer(name))
	if err != nil {
		return err
	}

	var b strings.Builder
	fmt.Fprintf(&b, "package %s\n\n", library)
	if comment != "" {
		for line := range strings.SplitSeq(comment, "\n") {
			b.WriteString("// " + line + "\n")
		}
	}
	fmt.Fprintf(&b, "templ %s(attrs templ.Attributes) {\n\t<svg", fname)

	for _, a := range attrs {
		fmt.Fprintf(&b, " %s=\"%s\"", a.Name.Local, escapeAttr(a.Value))
	}

	b.WriteString("\n\t\tif len(attrs) > 0 {\n\t\t\t{ attrs... }\n\t\t}\n\t>\n\t\t")
	b.WriteString(strings.TrimSpace(inner))
	b.WriteString("\n\t</svg>\n}\n")

	return os.WriteFile(filepath.Join(outputPath, name+".templ"), []byte(b.String()), 0644)
}

func escapeAttr(s string) string {
	var b strings.Builder
	xml.EscapeText(&b, []byte(s)) // re-escape, since Unmarshal already decoded entities
	return b.String()
}

func parseSVG(data []byte) (comment string, attrs []xml.Attr, inner string, err error) {
	dec := xml.NewDecoder(bytes.NewReader(data))
	var comments []string

	for {
		tok, terr := dec.Token()
		if terr != nil {
			if terr == io.EOF {
				err = fmt.Errorf("no root element found")
			} else {
				err = terr
			}
			return
		}

		switch t := tok.(type) {
		case xml.Comment:
			// t is []byte containing just the text between <!-- and -->
			comments = append(comments, strings.TrimSpace(string(t)))

		case xml.StartElement:
			if t.Name.Local != "svg" {
				err = fmt.Errorf("root element is <%s>, not <svg>", t.Name.Local)
				return
			}
			attrs = t.Attr
			comment = strings.Join(comments, "\n")

			// hand the already-consumed start element to DecodeElement
			// so it can grab the rest (children) as raw innerxml
			var body struct {
				Inner string `xml:",innerxml"`
			}
			if derr := dec.DecodeElement(&body, &t); derr != nil {
				err = derr
				return
			}
			inner = body.Inner
			return
		}
		// xml.ProcInst, xml.Directive, xml.CharData (whitespace): ignore
	}
}

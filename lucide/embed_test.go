package lucide_test

import (
	"html/template"
	"strings"
	"testing"

	"github.com/dimmerz92/go-icons/lucide"
)

func TestGetHTMLFile(t *testing.T) {
	name := "a-arrow-up"
	data, err := lucide.GetHTMLFile(name)
	if err != nil {
		t.Fatal(err)
	}

	if len(data) == 0 {
		t.Fatal("returned empty data")
	}
}

func TestAddIcons(t *testing.T) {
	html := `{{ define "icon" }}<svg>...</svg>{{ end }}`
	tpls := template.New("base")

	_, err := tpls.Parse(html)
	if err != nil {
		t.Fatalf("failed to parse: %v", err)
	}

	err = lucide.AddIcons(tpls)
	if err != nil {
		t.Fatalf("failed to add icons: %v", err)
	}

	t.Run("test existence", func(t *testing.T) {
		if tpls.Lookup("a-arrow-down") == nil {
			t.Fatal("expected icon to be present, got none")
		}
	})

	t.Run("test without args", func(t *testing.T) {
		want := `
<svg
  xmlns="http://www.w3.org/2000/svg"
  width="24"
  height="24"
  viewBox="0 0 24 24"
  fill="none"
  stroke="currentColor"
  stroke-width="2"
  stroke-linecap="round"
  stroke-linejoin="round"
  
>
  <path d="m14 12 4 4 4-4" />
  <path d="M18 16V7" />
  <path d="m2 16 4.039-9.69a.5.5 0 0 1 .923 0L11 16" />
  <path d="M3.304 13h6.392" />
</svg>
`
		var b strings.Builder
		err := tpls.ExecuteTemplate(&b, "a-arrow-down", nil)
		if err != nil {
			t.Fatalf("failed to execute: %v", err)
		}

		got := b.String()
		if got != want {
			t.Fatalf("got\n%s\n\nwanted\n%s", got, want)
		}
	})

	t.Run("test with args", func(t *testing.T) {
		want := `
<svg
  xmlns="http://www.w3.org/2000/svg"
  width="24"
  height="24"
  viewBox="0 0 24 24"
  fill="none"
  stroke="currentColor"
  stroke-width="2"
  stroke-linecap="round"
  stroke-linejoin="round"
  
    id="my-icon"
  
>
  <path d="m14 12 4 4 4-4" />
  <path d="M18 16V7" />
  <path d="m2 16 4.039-9.69a.5.5 0 0 1 .923 0L11 16" />
  <path d="M3.304 13h6.392" />
</svg>
`
		var b strings.Builder
		err := tpls.ExecuteTemplate(&b, "a-arrow-down", []template.HTMLAttr{`id="my-icon"`})
		if err != nil {
			t.Fatalf("failed to execute: %v", err)
		}

		got := b.String()
		if got != want {
			t.Fatalf("got\n%s\n\nwanted\n%s", got, want)
		}
	})

}

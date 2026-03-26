package main

import (
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"slices"

	"github.com/dimmerz92/go-icons/internal"
)

// key=library, value[0]=icons path in cloned repo, value[1]=path in current repo
var supportedLibraries = map[string][]string{
	"lucide": {"./lucide-repo/icons", "./lucide"},
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("error: usage: sync <library>")
		os.Exit(1)
	}

	library := os.Args[1]
	paths, ok := supportedLibraries[library]
	if !ok {
		fmt.Printf(
			"error: unsupported library [%s], use any of: %v\n",
			library, slices.Collect(maps.Keys(supportedLibraries)),
		)
		os.Exit(1)
	}

	svgIcons := internal.FileSet(filepath.Join(paths[0]), ".svg")
	htmlWant := internal.DiffFileSet(svgIcons, internal.FileSet(paths[1], ".html"))
	templWant := internal.DiffFileSet(svgIcons, internal.FileSet(paths[1], ".templ"))

	for svg := range svgIcons {
		_, inHtml := htmlWant[svg]
		_, inTempl := templWant[svg]

		var file []byte
		var err error

		if inHtml || inTempl {
			file, err = os.ReadFile(filepath.Join(paths[0], svg+".svg"))
			if err != nil {
				panic(err)
			}
		}

		if inHtml {
			err = internal.ToHTML(svg, string(file), paths[1])
			if err != nil {
				panic(err)
			}
		}

		if inTempl {
			err = internal.ToTempl(svg, string(file), paths[1])
			if err != nil {
				panic(err)
			}
		}
	}

	fmt.Printf("Finished: %d html files & %d templ files", len(htmlWant), len(templWant))
}

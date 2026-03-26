package main

import (
	"flag"
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/dimmerz92/go-icons/internal"
	"github.com/fatih/color"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Println(help)
		os.Exit(1)
	}

	parts := strings.Split(os.Args[1], ":")
	if len(parts) != 2 {
		fmt.Println(help)
		os.Exit(1)
	}

	library := parts[0]
	if _, ok := internal.SupportedLibraries[library]; !ok {
		color.Red(
			"error: unsupported library [%s], use any of %v",
			library, slices.Collect(maps.Keys(internal.SupportedLibraries)),
		)
		os.Exit(1)
	}

	format := strings.ToLower(parts[1])
	if _, ok := internal.SupportedFormats[format]; !ok {
		color.Red(
			"error: unsupported format [%s], use any of %v",
			format, slices.Collect(maps.Keys(internal.SupportedFormats)),
		)
		os.Exit(1)
	}

	icon := os.Args[2]

	var (
		file []byte
		err  error
	)

	switch format {
	case "html":
		file, err = internal.GetHTMLFile(library, icon)

	case "templ":
		file, err = internal.GetTemplFile(library, icon)
	}

	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}

	f := flag.NewFlagSet("templ", flag.ContinueOnError)
	output := f.String("out", ".", "the output directory")
	f.Parse(args[3:])

	err = os.WriteFile(filepath.Join(*output, icon+"."+format), file, 0755)
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}
}

var help = fmt.Sprint(
	fmt.Sprintf("\n%s\n\tgo-icons %s %s\n\n",
		color.YellowString("USAGE:"),
		color.BlueString("<LIBRARY>:<FORMAT> <ICON>"),
		color.MagentaString("[OPTIONS]"),
	),
	fmt.Sprintf("%s\n\t%s\n\n",
		color.YellowString("LIBRARIES:"),
		color.BlueString("%v", slices.Collect(maps.Keys(internal.SupportedLibraries))),
	),
	fmt.Sprintf("%s\n\t%s\n\n",
		color.YellowString("FORMATS:"),
		color.BlueString("%v", slices.Collect(maps.Keys(internal.SupportedFormats))),
	),
	color.YellowString("OPTIONS:\n"),
	fmt.Sprintf("\t%s: directory to save icon to. Default: [.]\n",
		color.MagentaString("-out"),
	),
)

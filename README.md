# Go Icons

Go icons provides an aggregation of icon libraries for golang developers using either html/template or [a-h/templ](https://github.com/a-h/templ).

Simply use the command line utility to install individual icons into your project, or use them directly as `templ.Component` templates.

## Icon Libraries

| Icon Library | Version | License |
| - | - | - |
| [Ionicons](https://github.com/ionic-team/ionicons) | ![Ionicons version](https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fraw.githubusercontent.com%2Fdimmerz92%2Fgo-icons%2Frefs%2Fheads%2Fmaster%2Fversions.json&query=%24.ionicons&style=flat-square&label=) | [MIT](https://github.com/ionic-team/ionicons/blob/main/LICENSE) |
| [Lucide Icons](https://github.com/lucide-icons/lucide) | ![Lucide Icons version](https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fraw.githubusercontent.com%2Fdimmerz92%2Fgo-icons%2Frefs%2Fheads%2Fmaster%2Fversions.json&query=%24.lucide&style=flat-square&label=) | [ISC](https://github.com/lucide-icons/lucide/blob/main/LICENSE) |
| [Material Icons](https://github.com/mui/material-ui) | ![Material Icons version](https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fraw.githubusercontent.com%2Fdimmerz92%2Fgo-icons%2Frefs%2Fheads%2Fmaster%2Fversions.json&query=%24.material&style=flat-square&label=) | [MIT](https://github.com/mui/material-ui/blob/master/LICENSE) |
| [Radix Icons](https://github.com/radix-ui/icons) | ![Radix Icons version](https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fraw.githubusercontent.com%2Fdimmerz92%2Fgo-icons%2Frefs%2Fheads%2Fmaster%2Fversions.json&query=%24.radix&style=flat-square&label=) | [MIT](https://github.com/radix-ui/icons/blob/main/LICENSE) |
| [Simple Icons](https://github.com/simple-icons/simple-icons) | ![Simple Icons version](https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fraw.githubusercontent.com%2Fdimmerz92%2Fgo-icons%2Frefs%2Fheads%2Fmaster%2Fversions.json&query=%24.simpleIcons&style=flat-square&label=) | [CC0 1.0](https://github.com/simple-icons/simple-icons/blob/develop/LICENSE.md) |

## Installation & Usage

Go Icons can be used in two different ways:

- Icons imported as a package of templ icons
- A command line utility to generate html or templ icons directly into your project

### Packages

> [!NOTE]
> If an icon you want to use starts with a number, it will be prefixed with a capital T.
>
> E.g,
>
> ```templ
> @material.T10_24px()
> ```

#### Ionicons

**Install**

```bash
go get github.com/dimmerz92/go-icons/ionicons@latest
```

**Usage**

```templ
@ionicons.Accessibility() // no argument
@ionicons.Accessibility(templ.Attributes{}) // templ.Attributes argument
```

#### Lucide Icons

**Install**

```bash
go get github.com/dimmerz92/go-icons/lucide@latest
```

**Usage**

```templ
@lucide.AArrowDown() // no argument
@lucide.AArrowDown(templ.Attributes{}) // templ.Attributes argument
```

#### Material Icons

**Install**

```bash
go get github.com/dimmerz92/go-icons/material@latest
```

**Usage**

```templ
@material.T10_24px() // no argument
@material.T10_24px(templ.Attributes{}) // templ.Attributes argument
```

#### Radix Icons

**Install**

```bash
go get github.com/dimmerz92/go-icons/radix-icons@latest
```

**Usage**

```templ
@radixicons.Accessibility() // no argument
@radixicons.Accessibility(templ.Attributes{}) // templ.Attributes argument
```

#### Simple Icons

**Install**

```bash
go get github.com/dimmerz92/go-icons/simple-icons@latest
```

**Usage**

```templ
@simpleicons.T1001trackLists() // no argument
@simpleicons.T1001trackLists(templ.Attributes{}) // templ.Attributes argument
```

### Command Line Utility

**Install**

```bash
go install github.com/dimmerz92/go-icons/cmd/go-icons@latest
```

**Usage**

> [!NOTE]
> The command follows a simple model of `go-lucide <LIBRARY>:<FORMAT> <ICON NAME> [OPTIONS]`
>
> LIBRARIES: [simple-icons radix-icons ionicons material lucide]
>
> FORMATS: [html templ]
>
> ICON NAMES: depending on the library, some are kebab case, snake case, or not separated.
>
> This is an understandable painpoint, and search functionality will be on the future todos.
>
> OPTIONS: \
>   -out: directory to save icon to. Default: [.]

```bash
go-icons lucide:html a-arrow-down -out templates/icons
```

**HTML Icon Usage**

The HTML icons are prepared to accept optional attributes on the `<svg>` tags.

This can easily be done by passing them the `template.HTMLAttr` data structure.

```go
iconData := []template.HTMLAttr{
    `class="some-class another-class"`,
    `style="height: 2rem; width: 2rem"`,
}
err := tpls.ExecuteTemplate(w, "a-arrow-down", iconData)
```

```html
<!DOCTYPE html>
<html>
    <head>
	<title>My Page</title>
    </head>
    <body>
        <p>some text</p>
        <!-- expecting data -->
        {{ template "worm" . }}
        <!-- not expecting data -->
        {{ template "fish" }}
    </body>
</html>
```

## LICENSE

This project is provided under the [MIT License](./LICENSE)

All icons are provided under their respective licenses listed at the top of this document.

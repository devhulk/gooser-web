module github.com/devhulk/gooser-web

go 1.21.4

replace github.com/devhulk/gooser-web/components v0.0.0 => ./components

replace github.com/devhulk/gooser-web/handlers v0.0.0 => ./handlers

replace github.com/devhulk/gooser-web/services v0.0.0 => ./services

require github.com/a-h/templ v0.2.513

require (
	github.com/fatih/color v1.16.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	golang.org/x/sys v0.14.0 // indirect
)

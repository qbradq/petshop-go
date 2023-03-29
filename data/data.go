package data

import "embed"

// StaticFS is the static file system for the web server.
//
//go:embed static
var StaticFS embed.FS

// TemplateFS is the template file system for the web server.
//
//go:embed templates
var TemplateFS embed.FS

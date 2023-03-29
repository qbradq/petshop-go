package data

import "embed"

// StaticFS is the static file system for the web server.
//
//go:embed static
var StaticFS embed.FS

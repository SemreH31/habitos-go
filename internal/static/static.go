package static

import (
	"embed"
)

//go:embed all:web
var content embed.FS

// FS expone el sub-dir “web” como http.FileSystem
func FS() embed.FS {
	return content
}

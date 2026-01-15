package static

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed all:web
var content embed.FS

// FS expone el sub-dir “web” como http.FileSystem
func FS() http.FileSystem {
	sub, _ := fs.Sub(content, "web")
	return http.FS(sub)
}

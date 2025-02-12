package main

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed html/index.html
var indexHTML embed.FS

//go:embed css/styles.css
var stylesCSS embed.FS

func handler(fsys embed.FS, subfolder string) http.Handler {
	html, _ := fs.Sub(fs.FS(fsys), subfolder)

	return http.FileServer(http.FS(html))
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", handler(indexHTML, "html"))
	mux.Handle("/styles.css", handler(stylesCSS, "css"))
	http.ListenAndServe(":8000", mux)
}

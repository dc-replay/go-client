package main

import (
	"flag"
	"github.com/dc-replay/go-client/ipcmux"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var builtinMimeTypesLower = map[string]string{
	".css":  "text/css; charset=utf-8",
	".gif":  "image/gif",
	".htm":  "text/html; charset=utf-8",
	".html": "text/html; charset=utf-8",
	".jpg":  "image/jpeg",
	".js":   "application/javascript",
	".wasm": "application/wasm",
	".pdf":  "application/pdf",
	".png":  "image/png",
	".svg":  "image/svg+xml",
	".xml":  "text/xml; charset=utf-8",
}

func Mime(ext string) string {
	if v, ok := builtinMimeTypesLower[ext]; ok {
		return v
	}
	return mime.TypeByExtension(ext)
}

func ServeParse() {
	addr := flag.String("ipc", "samplestatic", "IPC Addr")
	root := flag.String("root", ".", "Root dir")
	flag.Parse()
	Serve(*addr, *root)
}

func Serve(addr string, root string) {

	ipcmux.SetName(addr)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		path := strings.Split(r.URL.Path, "?")[0]
		if path == "" {
			path = "index.html"
		}
		ext := filepath.Ext(path)

		bs, err := os.ReadFile(filepath.Join(root, path))

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", Mime(ext))

		w.Write(bs)
	})
	ipcmux.ServeDefault()
}

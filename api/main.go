package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// https://blog.jetbrains.com/go/2021/06/09/how-to-use-go-embed-in-go-1-16/

//go:embed all:build/*
var build embed.FS

func main() {
	contentRoot, _ := fs.Sub(build, "build")
	fs := http.FileServer(http.FS(contentRoot))
	//fs := http.FileServer(http.Dir("./build"))
	//fs := http.FileServer(http.FS(build))
	http.Handle("/", Adapter("/", fs))

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Adapter adds html suffix is the request lacks an extension
func Adapter(prefix string, h http.Handler) http.Handler {
	if prefix == "" {
		return h
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := strings.TrimPrefix(r.URL.Path, prefix)
		rp := strings.TrimPrefix(r.URL.RawPath, prefix)
		if len(p) < len(r.URL.Path) && (r.URL.RawPath == "" || len(rp) < len(r.URL.RawPath)) {

			// append html
			if len(p) > 0 && !strings.Contains(p, ".") {
				p = p + ".html"
			}

			r2 := new(http.Request)
			*r2 = *r
			r2.URL = new(url.URL)
			*r2.URL = *r.URL
			r2.URL.Path = p
			r2.URL.RawPath = rp
			h.ServeHTTP(w, r2)
		} else {
			http.NotFound(w, r)
		}
	})
}

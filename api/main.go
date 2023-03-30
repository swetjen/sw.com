package main

import (
	"embed"
	"log"
	"net/http"
	"text/template"
)

// https://blog.jetbrains.com/go/2021/06/09/how-to-use-go-embed-in-go-1-16/

var (
	//go:embed build
	res   embed.FS
	pages = map[string]string{
		"/": "build/index.html",
	}
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		page, ok := pages[r.URL.Path]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		tpl, err := template.ParseFS(res, page)
		if err != nil {
			log.Printf("page %s not found in pages cache...", r.RequestURI)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		data := map[string]interface{}{
			"userAgent": r.UserAgent(),
		}

		if err := tpl.Execute(w, data); err != nil {
			return
		}

	})
	http.FileServer(http.FS(res))

	log.Println("Server started...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return err
	}

	return nil
}

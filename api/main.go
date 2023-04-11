package main

import (
	"context"
	"embed"
	"errors"
	"flag"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/sw.com/api/db"
	"go.mongodb.org/mongo-driver/mongo"
	"io/fs"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// https://blog.jetbrains.com/go/2021/06/09/how-to-use-go-embed-in-go-1-16/

//go:embed all:build/*
var build embed.FS

// Version is set during build.
var Version = "dev"

type Server struct {
	config *Config
	db     *mongo.Database
	router *chi.Mux
}

func run(args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.Usage = func() {
		fmt.Println(args[0] + " " + Version + ` usage: api [flags] [arguments]`)
		fmt.Println(`flags:`)
		flags.PrintDefaults()
	}

	var configFile = flags.String("config", "", "yaml configuration file to start the server with")
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}
	if *configFile == "" {
		flags.PrintDefaults()
		return errors.New("missing config.yaml")
	}

	file, err := os.ReadFile(*configFile)
	if err != nil {
		return err
	}

	config, err := NewConfig(file)
	if err != nil {
		return err
	}

	database := db.GetDatabase(config.DbUri, config.DbName)

	s := &Server{
		config: config,
		db:     database,
		router: chi.NewRouter(),
	}

	s.TestsDb()

	////
	s.setupRouter()

	contentRoot, _ := fs.Sub(build, "build")
	fs := http.FileServer(http.FS(contentRoot))
	s.router.Handle("/", Adapter("/", fs))

	log.Print("Listening on :3000...")
	err = http.ListenAndServe(":3000", s.router)
	if err != nil {
		log.Fatal(err)
	}

	return nil

}

func main() {
	if err := run(os.Args); err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
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

func (s *Server) TestsDb() {

	type testdoc struct {
		Label     string
		TSCreated int64
	}

	col := s.db.Collection("test")

	docsTodo := make([]interface{}, 1)
	docsTodo[0] = testdoc{
		Label:     "Startup Test",
		TSCreated: time.Now().Unix(),
	}

	_, err := col.InsertMany(context.TODO(), docsTodo)
	if err != nil {
		panic(err)
	}

	return

}

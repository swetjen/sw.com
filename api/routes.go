package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sw.com/api/db"
	"io"
	"net/http"
	"strings"
	"time"
)

func (s *Server) setupRouter() {
	// a good base middleware stack
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)

	// cors
	s.router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{fmt.Sprintf("https://%v", s.config.Host), fmt.Sprintf("http://%v", s.config.Host)},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Set timeout on request control context
	s.router.Use(middleware.Timeout(60 * time.Second))

	s.router.Post("/api/v1/contact", s.handleContactForm)
}

func respondErr(w http.ResponseWriter, r *http.Request, status int, args ...interface{}) {
	respond(w, r, status, map[string]interface{}{"error": map[string]interface{}{"message": fmt.Sprint(args...)}})
}

func respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	w.WriteHeader(status)
	if data != nil {
		Encode(w, r, status, data)
	}
}

// Encode writes the response.
func Encode(w http.ResponseWriter, r *http.Request, status int, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return errors.Wrap(err, "encode json")
	}
	var out io.Writer = w
	//if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
	//	w.Header().Set("Content-Encoding", "gzip")
	//	gzw := gzip.NewWriter(w)
	//	out = gzw
	//	defer gzw.Close()
	//}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if _, err := out.Write(b); err != nil {
		return err
	}
	return nil
}

func (s *Server) handleContactForm(w http.ResponseWriter, r *http.Request) {
	type ContactFormRequest struct {
		Name    string `json:"name,required"`
		Email   string `json:"email"`
		Message string `json:"message"`
		Guid    string `bson:"guid"`
	}

	type ContactFormResponse struct {
		Status string `json:"status"`
		Guid   string `json:"guid" bson:"guid"`
	}

	var form ContactFormRequest

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		respondErr(w, r, http.StatusUnprocessableEntity, "Unprocessable request body.")
	}

	form.Guid = strings.Replace(uuid.New().String(), "-", "", -1)

	fmt.Printf("%v\n", form)

	col := s.db.Collection(db.Contacts)
	_, err := col.InsertOne(r.Context(), form)

	if err != nil {
		respondErr(w, r, http.StatusUnprocessableEntity, "Unable to insert document")
	}

	respond(w, r, http.StatusOK, ContactFormResponse{
		Status: "Okay",
		Guid:   form.Guid,
	})

}

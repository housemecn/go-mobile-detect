package main

import (
	"fmt"
	"log"
	"net/http"

	md "github.com/housemecn/go-mobile-detect"
)

// Handler .
type Handler struct{}

// Mobile .
func (h *Handler) Mobile(w http.ResponseWriter, r *http.Request, m *md.MobileDetect) {
	fmt.Fprint(w, "Hello, this is mobile")
}

// Tablet .
func (h *Handler) Tablet(w http.ResponseWriter, r *http.Request, m *md.MobileDetect) {
	fmt.Fprint(w, "Hello, this is tablet")
}

// Desktop .
func (h *Handler) Desktop(w http.ResponseWriter, r *http.Request, m *md.MobileDetect) {
	fmt.Fprint(w, "Hello, this is desktop", m.MobileGrade())
}

func main() {
	log.Println("Starting local server http://localhost:10001/check (cmd+click to open from terminal)")
	mux := http.NewServeMux()
	h := &Handler{}
	mux.Handle("/", md.Handler(h, nil))
	http.ListenAndServe(":10001", mux)
}

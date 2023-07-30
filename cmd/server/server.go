package main

import (
	"fmt"
	"net/http"
)

type server struct {
	port int
	mux  *http.ServeMux
}

func NewServer(mux *http.ServeMux, port int) *server {
	return &server{mux: mux, port: port}
}

func (s *server) Start() {
	fmt.Printf("server running on port: %v\n", s.port)
	http.ListenAndServe(fmt.Sprintf(":%v", s.port), s.mux)
}

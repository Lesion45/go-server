package main

import (
	"go-server/v1/config"
	"log"
	"net/http"
)

type Server struct {
	Addr string
}

func (s *Server) Run() {
	hub := newHub()
	go hub.run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(s.Addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func NewServer(addr string) *Server {
	return &Server{
		Addr: addr,
	}
}

func main() {
	cfg := config.New()
	server := NewServer(cfg.Addr)
	server.Run()
}

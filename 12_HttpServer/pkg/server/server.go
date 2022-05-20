package server

import (
	"httpServer/pkg/calendar"
	"httpServer/pkg/handler"
	"log"
	"net/http"
)

type IHandler interface {
	Hand(w http.ResponseWriter, r *http.Request)
	Pattern() string
}

type Server struct {
	events  *calendar.Events
	routing *http.ServeMux
}

func (s *Server) HandlerIt(handler IHandler) {
	s.routing.HandleFunc(handler.Pattern(), handler.Hand)
}

func NewServer() *Server {
	return &Server{calendar.NewEvents(), http.NewServeMux()}
}

func (s *Server) Run() {
	srv := &http.Server{
		Addr:    "localhost:8080",
		Handler: s.routing,
	}
	create := &handler.CreateEvent{Events: s.events}
	s.HandlerIt(create)
	log.Fatal(srv.ListenAndServe())
}

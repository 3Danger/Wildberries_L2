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
	s.HandlerIt(&handler.CreateEvent{Events: s.events})
	s.HandlerIt(&handler.UpdateEvent{Events: s.events})
	s.HandlerIt(&handler.DeleteEvent{Events: s.events})
	s.HandlerIt(&handler.EventsForDay{Events: s.events})
	s.HandlerIt(&handler.EventsForWeek{Events: s.events})
	s.HandlerIt(&handler.EventsForMonth{Events: s.events})
	log.Fatal(srv.ListenAndServe())
}
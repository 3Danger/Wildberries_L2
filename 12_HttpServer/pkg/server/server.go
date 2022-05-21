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
	events  *calendar.EventsManager
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
	s.HandlerIt(&handler.CreateEvent{EventsManager: s.events})
	s.HandlerIt(&handler.UpdateEvent{EventsManager: s.events})
	s.HandlerIt(&handler.DeleteEvent{EventsManager: s.events})
	s.HandlerIt(&handler.EventsForDay{EventsManager: s.events})
	s.HandlerIt(&handler.EventsForWeek{EventsManager: s.events})
	s.HandlerIt(&handler.EventsForMonth{EventsManager: s.events})
	log.Fatal(srv.ListenAndServe())
}

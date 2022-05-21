package server

import (
	"httpServer/pkg/calendar"
	"log"
	"net/http"
	"time"
)

type Server struct {
	events  *calendar.EventsManager
	routing *http.ServeMux
}

func NewServer() *Server {
	return &Server{calendar.NewEvents(), http.NewServeMux()}
}

func MiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		log.Printf("%s, %s, %s\n", r.Method, r.URL, time.Since(start))
	}
}

func (s *Server) Run() {
	srv := &http.Server{
		Addr:    "localhost:8080",
		Handler: s.routing,
	}
	s.routing.HandleFunc("/create_event", MiddleWare(s.CreateEvent))
	s.routing.HandleFunc("/update_event", MiddleWare(s.UpdateEvent))
	s.routing.HandleFunc("/delete_event", MiddleWare(s.DeleteEvent))
	s.routing.HandleFunc("/events_for_day", MiddleWare(s.EventsForDay))
	s.routing.HandleFunc("/events_for_week", MiddleWare(s.EventsForWeek))
	s.routing.HandleFunc("/events_for_month", MiddleWare(s.EventsForMonth))
	log.Println("start server")
	log.Fatal(srv.ListenAndServe())
}

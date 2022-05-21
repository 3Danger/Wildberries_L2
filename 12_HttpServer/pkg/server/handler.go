package server

import (
	"encoding/json"
	uid "github.com/google/uuid"
	"net/http"
	"time"
)

func (s *Server) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if ValidateQuery(w, r, http.MethodPost) {
		ev := EventModel{}
		if err := json.NewDecoder(r.Body).Decode(&ev); err != nil {
			jsonResponse(true, w, http.StatusServiceUnavailable, err.Error())
		} else if !ev.Validate() {
			jsonResponse(true, w, http.StatusServiceUnavailable, "user_id is empty or date wasn't formatted")
		} else {
			ev.ID = uid.New().String()
			s.events.Add(ev.ToEvent())
			jsonResponse(false, w, http.StatusOK, "created")
		}
	}
}

func (s *Server) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if ValidateQuery(w, r, http.MethodPost) {
		ev := EventModel{}
		if ok := json.NewDecoder(r.Body).Decode(&ev); ok != nil {
			jsonResponse(true, w, http.StatusServiceUnavailable, ok.Error())
		} else if ok := s.events.UpdateEvent(ev.ToEvent()); ok != nil {
			jsonResponse(true, w, http.StatusServiceUnavailable, ok.Error())
		} else {
			jsonResponse(false, w, http.StatusOK, "updated")
		}
	}
}

func (s *Server) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if ValidateQuery(w, r, http.MethodPost) {
		ev := EventModel{}
		if err := json.NewDecoder(r.Body).Decode(&ev); err != nil || ev.ID == "" {
			jsonResponse(true, w, http.StatusServiceUnavailable, "not parameter ID or isn't correct")
			return
		}
		if ok := s.events.DeleteEvent(ev.ID); !ok {
			jsonResponse(true, w, http.StatusServiceUnavailable, "not found")
			return
		}
		jsonResponse(false, w, http.StatusOK, "deleted")
	}
}

func (s *Server) EventsForDay(w http.ResponseWriter, r *http.Request) {
	if ValidateQuery(w, r, http.MethodGet, "user_id", "date") {
		if date, ok := time.Parse("2006-01-02", r.URL.Query().Get("date")); ok != nil {
			jsonResponse(true, w, http.StatusServiceUnavailable, ok.Error())
		} else {
			userId := r.URL.Query().Get("user_id")
			evs := s.events.GetEvents(userId, date, date.AddDate(0, 0, 1))
			jsonResponse(false, w, http.StatusOK, evs)
		}
	}
}

func (s *Server) EventsForWeek(w http.ResponseWriter, r *http.Request) {
	if ValidateQuery(w, r, http.MethodGet, "user_id", "date") {
		if date, ok := time.Parse("2006-01-02", r.URL.Query().Get("date")); ok != nil {
			jsonResponse(true, w, http.StatusServiceUnavailable, ok.Error())
		} else {
			userId := r.URL.Query().Get("user_id")
			evs := s.events.GetEvents(userId, date, date.AddDate(0, 0, 7))
			jsonResponse(false, w, http.StatusOK, evs)
		}
	}
}

func (s *Server) EventsForMonth(w http.ResponseWriter, r *http.Request) {
	if ValidateQuery(w, r, http.MethodGet, "user_id", "date") {
		if date, ok := time.Parse("2006-01-02", r.URL.Query().Get("date")); ok != nil {
			jsonResponse(true, w, http.StatusBadRequest, ok.Error())
		} else {
			userId := r.URL.Query().Get("user_id")
			evs := s.events.GetEvents(userId, date, date.AddDate(0, 1, 0))
			jsonResponse(false, w, http.StatusOK, evs)
		}
	}
}

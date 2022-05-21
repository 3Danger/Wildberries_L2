package handler

import (
	"encoding/json"
	"fmt"
	uid "github.com/google/uuid"
	"httpServer/pkg/calendar"
	"net/http"
	"time"
)

// Result - сообщение о результате
type Result struct {
	Msg interface{} `json:"result"`
}

type Error struct {
	Msg string `json:"error"`
}

//http ответ
func jsonResponse(err bool, w http.ResponseWriter, code int, msg interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	if err {
		errResp := Error{msg.(string)}
		if err := json.NewEncoder(w).Encode(errResp); err != nil {
			http.Error(w, "responseJson error", http.StatusInternalServerError)
		}
	} else {
		resResp := Result{msg}
		if err := json.NewEncoder(w).Encode(resResp); err != nil {
			http.Error(w, "responseJson error", http.StatusInternalServerError)
		}
	}
}

/*
   POST /create_event
   POST /update_event
   POST /delete_event
   GET /events_for_day
   GET /events_for_week
   GET /events_for_month
*/

// CreateEvent POST
type CreateEvent struct {
	*calendar.EventsManager
}

// Pattern return this http map
func (CreateEvent) Pattern() string { return "/create_event" }

// Hand processing query
func (c *CreateEvent) Hand(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		jsonResponse(true, w, http.StatusMethodNotAllowed, fmt.Sprintf("bad %v method", r.Method))
		return
	}
	ev := EventModel{}
	if err := json.NewDecoder(r.Body).Decode(&ev); err != nil {
		jsonResponse(true, w, http.StatusBadRequest, err.Error())
	} else if !ev.Validate() {
		jsonResponse(true, w, http.StatusBadRequest, "user_id is empty or date wasn't formatted")
	} else {
		ev.ID = uid.New().String()
		c.Add(ev.ToEvent())
		jsonResponse(false, w, http.StatusOK, "created")
	}
}

// UpdateEvent POST
type UpdateEvent struct {
	*calendar.EventsManager
}

func (UpdateEvent) Pattern() string { return "/update_event" }
func (e UpdateEvent) Hand(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		jsonResponse(true, w, http.StatusMethodNotAllowed, fmt.Sprintf("bad %v method", r.Method))
	}
	ev := EventModel{}
	if ok := json.NewDecoder(r.Body).Decode(&ev); ok != nil {
		jsonResponse(true, w, http.StatusBadRequest, ok.Error())
	} else if ok := e.UpdateEvent(ev.ToEvent()); ok != nil {
		jsonResponse(true, w, http.StatusBadRequest, ok.Error())
	} else {
		jsonResponse(false, w, http.StatusOK, "updated")
	}
}

// DeleteEvent POST
type DeleteEvent struct {
	*calendar.EventsManager
}

func (DeleteEvent) Pattern() string { return "/delete_event" }

func (e *DeleteEvent) Hand(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		jsonResponse(true, w, http.StatusMethodNotAllowed, fmt.Sprintf("bad %v method", r.Method))
		return
	}
	ev := EventModel{}
	if err := json.NewDecoder(r.Body).Decode(&ev); err != nil || ev.ID == "" {
		jsonResponse(true, w, http.StatusBadRequest, "not parameter ID or isn't correct")
		return
	}
	if ok := e.DeleteEvent(ev.ID); !ok {
		jsonResponse(true, w, http.StatusNotFound, "not found")
		return
	}
	jsonResponse(false, w, http.StatusOK, "deleted")
}

// EventsForDay GET
type EventsForDay struct {
	*calendar.EventsManager
}

func (EventsForDay) Pattern() string { return "/events_for_day" }
func (d *EventsForDay) Hand(w http.ResponseWriter, r *http.Request) {
	if ValidateQuery(w, r, "GET", "user_id", "date") {
		if date, ok := time.Parse("2006-01-02", r.URL.Query().Get("date")); ok != nil {
			jsonResponse(true, w, http.StatusBadRequest, ok.Error())
		} else {
			userId := r.URL.Query().Get("user_id")
			evs := d.EventsManager.GetEvents(userId, date, date.AddDate(0, 0, 1))
			jsonResponse(false, w, http.StatusOK, evs)
		}
	}
}

// EventsForWeek GET
type EventsForWeek struct {
	*calendar.EventsManager
}

func (EventsForWeek) Pattern() string { return "/events_for_week" }
func (w2 EventsForWeek) Hand(w http.ResponseWriter, r *http.Request) {
	if ValidateQuery(w, r, "GET", "user_id", "date") {
		if date, ok := time.Parse("2006-01-02", r.URL.Query().Get("date")); ok != nil {
			jsonResponse(true, w, http.StatusBadRequest, ok.Error())
		} else {
			userId := r.URL.Query().Get("user_id")
			evs := w2.EventsManager.GetEvents(userId, date, date.AddDate(0, 0, 7))
			jsonResponse(false, w, http.StatusOK, evs)
		}
	}
}

// EventsForMonth GET
type EventsForMonth struct {
	*calendar.EventsManager
}

func (m EventsForMonth) Hand(w http.ResponseWriter, r *http.Request) {
	if ValidateQuery(w, r, "GET", "user_id", "date") {
		if date, ok := time.Parse("2006-01-02", r.URL.Query().Get("date")); ok != nil {
			jsonResponse(true, w, http.StatusBadRequest, ok.Error())
		} else {
			userId := r.URL.Query().Get("user_id")
			evs := m.EventsManager.GetEvents(userId, date, date.AddDate(0, 1, 0))
			jsonResponse(false, w, http.StatusOK, evs)
		}
	}

}

func (EventsForMonth) Pattern() string { return "/events_for_month" }

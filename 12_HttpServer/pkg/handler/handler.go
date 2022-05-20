package handler

import (
	"encoding/json"
	"fmt"
	uid "github.com/google/uuid"
	"httpServer/pkg/calendar"
	"net/http"
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
		resResp := Result{
			Msg: msg,
		}
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
	*calendar.Events
}

// Pattern return this http map
func (CreateEvent) Pattern() string { return "/create_event" }

// Hand processing query
func (c *CreateEvent) Hand(w http.ResponseWriter, r *http.Request) {
	fmt.Println("crea")
	if r.Method != "POST" {
		jsonResponse(true, w, http.StatusMethodNotAllowed, fmt.Sprintf("bad %v method", r.Method))
		return
	}
	ev := EventModel{}
	if err := json.NewDecoder(r.Body).Decode(&ev); err != nil {
		jsonResponse(true, w, http.StatusBadRequest, err.Error())
		return
	}
	if !ev.Validate() {
		jsonResponse(true, w, http.StatusBadRequest, "user_id is empty or date wasn't formatted")
		return
	}
	ev.ID = uid.New().String()
	c.Add(ev.ToEvent())
	jsonResponse(false, w, http.StatusOK, "created")
}

// UpdateEvent POST
type UpdateEvent struct {
	*calendar.Events
}

func (e UpdateEvent) Hand(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (UpdateEvent) Pattern() string { return "/update_event" }

// DeleteEvent POST
type DeleteEvent struct {
	*calendar.Events
}

func (e DeleteEvent) Hand(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (DeleteEvent) Pattern() string { return "/delete_event" }

// EventsForDay GET
type EventsForDay struct {
	*calendar.Events
}

func (d EventsForDay) Hand(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (EventsForDay) Pattern() string { return "/events_for_day" }

// EventsForWeek GET
type EventsForWeek struct {
	*calendar.Events
}

func (w2 EventsForWeek) Hand(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (EventsForWeek) Pattern() string { return "/events_for_week" }

// EventsForMonth GET
type EventsForMonth struct {
	*calendar.Events
}

func (m EventsForMonth) Hand(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (EventsForMonth) Pattern() string { return "/events_for_month" }

package calendar

import (
	"sync"
	"time"
)

type Event struct {
	ID          string    `json:"id"`
	Date        time.Time `json:"date"`
	UserID      string    `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

func NewEvent() *Event {
	return new(Event)
}

type Events struct {
	*sync.RWMutex
	data map[string]*Event
}

func NewEvents() *Events {
	return &Events{&sync.RWMutex{}, make(map[string]*Event)}
}

func (es *Events) Add(e *Event) {
	es.Lock()
	es.data[e.ID] = e
	es.Unlock()
}

func (es *Events) DeleteEvent(id string) (ok bool) {
	es.Lock()
	if _, ok = es.data[id]; ok {
		delete(es.data, id)
	}
	es.Unlock()
	return ok
}

func (es *Events) GetEvent(id string) *Event {
	es.RLock()
	defer es.RUnlock()
	return es.data[id]
}

package calendar

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type EventsManager struct {
	*sync.RWMutex
	data map[string]*Event
}

func NewEvents() *EventsManager {
	return &EventsManager{&sync.RWMutex{}, make(map[string]*Event)}
}

func (es *EventsManager) Add(e *Event) {
	es.Lock()
	es.data[e.ID] = e
	es.Unlock()
}

func (es *EventsManager) DeleteEvent(id string) (ok bool) {
	es.Lock()
	if _, ok = es.data[id]; ok {
		delete(es.data, id)
	}
	es.Unlock()
	return ok
}

func (es *EventsManager) GetEvent(id string) *Event {
	es.RLock()
	defer es.RUnlock()
	return es.data[id]
}

func (es *EventsManager) UpdateEvent(e *Event) (ok error) {
	curEv := es.GetEvent(e.ID)
	if curEv == nil {
		return errors.New(fmt.Sprintf("event id with: %s - not found", e.ID))
	}
	if e.Title != "" {
		curEv.Lock()
		curEv.Title = e.Title
		curEv.Unlock()
	}
	if !e.Date.IsZero() {
		curEv.Lock()
		curEv.Date = e.Date
		curEv.Unlock()
	}
	if e.Description != "" {
		curEv.Lock()
		curEv.Description = e.Description
		curEv.Unlock()
	}
	if e.UserID != "" {
		curEv.Lock()
		curEv.UserID = e.UserID
		curEv.Unlock()
	}
	return nil
}

func (es *EventsManager) GetEvents(userId string, start, end time.Time) (ev []*Event) {
	es.RLock()
	for _, v := range es.data {
		if v.UserID == userId {
			if (v.Date == start || v.Date.After(start)) && v.Date.Before(end) {
				ev = append(ev, v)
			}
		}
	}
	es.RUnlock()
	return ev
}

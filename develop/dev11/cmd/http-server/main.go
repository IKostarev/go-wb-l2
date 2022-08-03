package main

import (
	"http-server-cust/internal/handlers"
	"http-server-cust/internal/middleware"
	"log"
	"net/http"
)

type myMux struct {
	*http.ServeMux
	structSlash bool
}

func (m *myMux) StructSlash(val bool) {
	m.structSlash = true
}

func main() {
	r := myMux{
		http.NewServeMux(),
		false,
	}
	r.StructSlash(true)

	EventCreateHandler := http.HandlerFunc(handlers.EventCreateHandler)
	EventUpdateHandler := http.HandlerFunc(handlers.EventUpdateHandler)
	EventDeleteHandler := http.HandlerFunc(handlers.EventDeleteHandler)
	EventsAllHandler := http.HandlerFunc(handlers.EventsAllHandler)
	EventsDayHandler := http.HandlerFunc(handlers.EventsDayHandler)
	EventsWeekHandler := http.HandlerFunc(handlers.EventsWeekHandler)
	EventsMonthHandler := http.HandlerFunc(handlers.EventsMonthHandler)

	r.Handle("create_event", middleware.MiddlewareJSONCheck(middleware.MiddlewareLog(EventCreateHandler)))
	r.Handle("update_event", middleware.MiddlewareJSONCheck(middleware.MiddlewareLog(EventUpdateHandler)))
	r.Handle("delete_event", middleware.MiddlewareLog(EventDeleteHandler))
	r.Handle("all_events", middleware.MiddlewareLog(EventsAllHandler))
	r.Handle("day_event", middleware.MiddlewareLog(EventsDayHandler))
	r.Handle("week_event", middleware.MiddlewareLog(EventsWeekHandler))
	r.Handle("month_event", middleware.MiddlewareLog(EventsMonthHandler))

	log.Fatal(http.ListenAndServe(":8080", r))
}

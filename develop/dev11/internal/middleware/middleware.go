package middleware

import (
	"http-server-cust/internal/helper"
	"http-server-cust/pkg/database"
	"http-server-cust/pkg/service"
	"log"
	"net/http"
)

func MiddlewareJSONCheck(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("415 - ожидаются данные в формате application/json"))
			return
		}
		handler.ServeHTTP(w, r)
	})
}

func MiddlewareLog(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		service.Contexts[r.Context()] = make(map[string]interface{})
		var lr database.LoggerRequest
		lr.Request = r.RequestURI
		lr.Client = r.RemoteAddr
		handler.ServeHTTP(w, r)
		res, ok := service.Contexts[r.Context()]["data"].(string)
		if ok {
			lr.Result = res
			log.Println(lr)
		}
		res, ok = service.Contexts[r.Context()]["err"].(string)
		if ok {
			helper.ProcessError(w, r)
		}
		delete(service.Contexts, r.Context())
	})
}

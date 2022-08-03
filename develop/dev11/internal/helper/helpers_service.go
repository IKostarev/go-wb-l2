package helper

import (
	"http-server-cust/pkg/service"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Result struct {
	Result interface{}
}

func ValidateDateTime(d, t string) bool {
	dt := d + " " + t

	_, err := time.Parse("01.01.2010 12:00:00", dt)
	if err != nil {
		return false
	}

	return true
}

func ProcessError(w http.ResponseWriter, r *http.Request) {
	validate, ok := service.Contexts[r.Context()]["err"].(string)
	if !ok {
		return
	}

	data := strings.Split(validate, "-")
	errcode, _ := strconv.Atoi(data[0])
	retErr := strings.Replace(`{"error":"{err}"}`, "{err}", validate, -1)

	w.WriteHeader(errcode)
	w.Write([]byte(retErr))
}

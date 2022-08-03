package handlers

import (
	"encoding/json"
	"http-server-cust/internal/helper"
	"http-server-cust/pkg/database"
	"http-server-cust/pkg/service"
	"io/ioutil"
	"net/http"
	"strconv"
)

func EventsDayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		service.Contexts[r.Context()]["err"] = "405 - Метод должен быть GET"
		service.Contexts[r.Context()]["data"] = "r.Method != \"GET\""
		return
	}

	day := r.URL.Query().Get("day")
	if helper.ValidateDateTime(day, "00:00:00") == false {
		service.Contexts[r.Context()]["err"] = "400-Некорректно указан день.\nПравильный формат даты: `ДД.ММ.ГГГГ`"
		service.Contexts[r.Context()]["data"] = "ValidateDateTime(day,\"00:00:00\")==false"
		return
	}

	service.EventsGetDay(day, r)
	var Data []database.Event
	Data, ok := service.Contexts[r.Context()]["events"].([]database.Event)
	if ok == false {
		return
	}
	if len(Data) == 0 {
		service.Contexts[r.Context()]["err"] = "200-нет записей"
		service.Contexts[r.Context()]["data"] = "нет записей"
		return
	}
	lines, err := json.MarshalIndent(helper.Result{Data}, "", "\t")
	if err != nil {
		service.Contexts[r.Context()]["err"] = "500-неопознанная ошибка"
		service.Contexts[r.Context()]["data"] = "json.MarshalIndent(Data,\"\",\"\\t\")"
		return
	}

	w.Write(lines)
}

func EventsWeekHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		service.Contexts[r.Context()]["err"] = "405 - Метод должен быть GET"
		service.Contexts[r.Context()]["data"] = "r.Method!=\"GET\""
		return
	}
	day := r.URL.Query().Get("day")
	if helper.ValidateDateTime(day, "00:00:00") == false {
		service.Contexts[r.Context()]["err"] = "400-Некорректно указан день.\nПравильный формат даты: `ДД.ММ.ГГГГ`"
		service.Contexts[r.Context()]["data"] = "ValidateDateTime(day,\"00:00:00\")==false"
		return
	}

	service.EventsGetWeek(day, r)
	var Data []database.Event
	Data, ok := service.Contexts[r.Context()]["events"].([]database.Event)
	if ok == false {
		return
	}

	if len(Data) == 0 {
		service.Contexts[r.Context()]["err"] = "200-нет записей"
		service.Contexts[r.Context()]["data"] = "нет записей"
		return
	}

	lines, err := json.MarshalIndent(helper.Result{Data}, "", "\t")
	if err != nil {
		service.Contexts[r.Context()]["err"] = "500-неопознанная ошибка"
		service.Contexts[r.Context()]["data"] = "json.MarshalIndent(Data,\"\",\"\\t\")"
		return
	}

	w.Write(lines)
}

func EventsMonthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		service.Contexts[r.Context()]["err"] = "405 - Метод должен быть GET"
		service.Contexts[r.Context()]["data"] = "r.Method!=\"GET\""
		return
	}
	day := r.URL.Query().Get("day")
	if helper.ValidateDateTime(day, "00:00:00") == false {
		service.Contexts[r.Context()]["err"] = "400-Некорректно указан день.\nПравильный формат даты: `ДД.ММ.ГГГГ`"
		service.Contexts[r.Context()]["data"] = "ValidateDateTime(day,\"00:00:00\")==false"
		return
	}
	service.EventsGetMonth(day, r)
	var Data []database.Event
	Data, ok := service.Contexts[r.Context()]["events"].([]database.Event)
	if ok == false {
		return
	}
	if len(Data) == 0 {
		service.Contexts[r.Context()]["err"] = "200-нет записей"
		service.Contexts[r.Context()]["data"] = "нет записей"
		return
	}
	lines, err := json.MarshalIndent(helper.Result{Data}, "", "\t")
	if err != nil {
		service.Contexts[r.Context()]["err"] = "500-неопознанная ошибка"
		service.Contexts[r.Context()]["data"] = "json.MarshalIndent(Data,\"\",\"\\t\")"
		return
	}
	w.Write(lines)
}

func EventsAllHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		service.Contexts[r.Context()]["err"] = "405 - Метод должен быть GET"
		service.Contexts[r.Context()]["data"] = "r.Method!=\"GET\""
		return
	}
	service.EventGetAll(r)
	var Data []database.Event
	Data, ok := service.Contexts[r.Context()]["events"].([]database.Event)
	if ok == false {
		return
	}
	if len(Data) == 0 {
		service.Contexts[r.Context()]["err"] = "200-нет записей"
		service.Contexts[r.Context()]["data"] = "нет записей"
		return
	}
	lines, err := json.MarshalIndent(helper.Result{Data}, "", "\t")
	if err != nil {
		service.Contexts[r.Context()]["err"] = "500-неопознанная ошибка"
		service.Contexts[r.Context()]["data"] = "json.MarshalIndent(Data,\"\",\"\\t\")"
		return
	}
	w.Write(lines)
}

func EventCreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		service.Contexts[r.Context()]["err"] = "405 - Метод должен быть POST"
		service.Contexts[r.Context()]["data"] = "r.Method!=\"POST\""
		return
	}
	var ev database.Event
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		service.Contexts[r.Context()]["err"] = "500-неопознанная ошибка"
		service.Contexts[r.Context()]["data"] = "ioutil.ReadAll(r.Body)"
		return
	}
	err = json.Unmarshal(body, &ev)
	if err != nil {
		service.Contexts[r.Context()]["err"] = "500-неопознанная ошибка"
		service.Contexts[r.Context()]["data"] = "json.Unmarshal(body,&ev)"
		return
	}
	if ev.Time == "" {
		ev.Time = "00:00:00"
	}
	if helper.ValidateDateTime(ev.Date, ev.Time) == false {
		service.Contexts[r.Context()]["err"] = "400-Некорректно указаны дата и время.\nПравильный формат даты: `ДД.ММ.ГГГГ`\nПравильный формат времени: `чч:мм:сс`"
		service.Contexts[r.Context()]["data"] = "ValidateDateTime(ev.Date,ev.Time)==false"
		return
	}
	service.EventCreate(ev, r)
}

func EventDeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		service.Contexts[r.Context()]["err"] = "405 - Метод должен быть POST"
		service.Contexts[r.Context()]["data"] = "r.Method!=\"POST\""
		return
	}
	id := r.URL.Query().Get("id")
	val, err := strconv.Atoi(id)
	if err != nil {
		service.Contexts[r.Context()]["data"] = "id:=r.URL.Query().Get(\"id\")\n\t_,err:=strconv.Atoi(id)"
		service.Contexts[r.Context()]["err"] = "400-Некорректный ввод. id=" + id + ". id должен быть целочисленным положителньым числом"
		return
	}
	service.EventDelete(val, r)
}

func EventUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		service.Contexts[r.Context()]["err"] = "405 - Метод должен быть POST"
		service.Contexts[r.Context()]["data"] = "r.Method!=\"POST\""
		return
	}
	var ev database.Event
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		service.Contexts[r.Context()]["err"] = "500-неопознанная ошибка"
		service.Contexts[r.Context()]["data"] = "ioutil.ReadAll(r.Body)"
		return
	}
	err = json.Unmarshal(body, &ev)
	if err != nil {
		service.Contexts[r.Context()]["err"] = "400-некорректный ввод"
		service.Contexts[r.Context()]["data"] = "json.Unmarshal(body,&ev)"
		return
	}
	if helper.ValidateDateTime(ev.Date, ev.Time) == false {
		service.Contexts[r.Context()]["err"] = "400-Некорректно указаны дата и время.\nПравильный формат даты: `ДД.ММ.ГГГГ`\nПравильный формат времени: `чч:мм:сс`"
		service.Contexts[r.Context()]["data"] = "ValidateDateTime(ev.Date,ev.Time)==false"
		return
	}
	service.EventUpdate(ev, r)
}

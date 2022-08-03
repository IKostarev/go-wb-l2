package service

import (
	"context"
	"fmt"
	"http-server-cust/pkg/database"
	"log"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	db       *sqlx.DB
	Contexts map[context.Context]map[string]interface{}
)

var schema = `
CREATE TABLE IF NOT EXISTS events (
id INTEGER PRIMARY KEY,
date TEXT,
name TEXT,
time TEXT
);`

func NewDb() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user = test_user dbname = test_dbname sslomde = disable")
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных - %s", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Ошибка пинга - %s", err)
	}

	db.Exec(schema)
	return db
}

func EventsGetDay(day string, r *http.Request) {
	query := `
	SELECT * FROM events
	WHERE date=?
	`

	rows, err := db.Query(query, day)
	if err != nil {
		Contexts[r.Context()]["data"] = fmt.Sprintf("Не смогли получить все записи из-за ошибки %v ", err.Error())
		Contexts[r.Context()]["err"] = "503-Ошибка бизнес-логики"
		return
	}

	var events []database.Event
	var Event database.Event

	for rows.Next() {
		err := rows.Scan(&Event.Id, &Event.Date, &Event.Name, &Event.Time)
		if err != nil {
			Contexts[r.Context()]["data"] = fmt.Sprintf("Не смогли считать строку с результатом из-за ошибки %v ", err.Error())
			Contexts[r.Context()]["err"] = "503-Ошибка бизнес-логики"
			return
		}
		events = append(events, Event)
	}

	Contexts[r.Context()]["events"] = events
}

func EventsGetWeek(day string, r *http.Request) {
	var events []database.Event

	week := calcWeek(day)

	for _, d := range week {
		query := `
		SELECT * FROM events
		WHERE date=?
		`
		rows, err := db.Query(query, d)
		if err != nil {
			Contexts[r.Context()]["data"] = fmt.Sprintf("Не смогли получить все записи из-за ошибки %v ", err.Error())
			Contexts[r.Context()]["err"] = "503-Ошибка бизнес-логики"
			return
		}

		var Event database.Event

		for rows.Next() {
			err := rows.Scan(&Event.Id, &Event.Date, &Event.Name, &Event.Time)
			if err != nil {
				Contexts[r.Context()]["data"] = fmt.Sprintf("Не смогли считать строку с результатом из-за ошибки %v ", err.Error())
				Contexts[r.Context()]["err"] = "503-Ошибка бизнес-логики"
				return
			}

			events = append(events, Event)
		}
	}

	Contexts[r.Context()]["events"] = events
}

func EventsGetMonth(day string, r *http.Request) {
	month := calcMonth(day)
	var events []database.Event
	for _, d := range month {
		querry := `
		SELECT * FROM events 
		WHERE date=?
		`

		rows, err := db.Query(querry, d)
		if err != nil {
			Contexts[r.Context()]["data"] = fmt.Sprintf("Не смогли получить все записи из-за ошибки %v ", err.Error())
			Contexts[r.Context()]["err"] = "503-Ошибка бизнес-логики"
			return
		}

		var Event database.Event

		for rows.Next() {
			err := rows.Scan(&Event.Id, &Event.Date, &Event.Name, &Event.Time)
			if err != nil {
				Contexts[r.Context()]["data"] = fmt.Sprintf("Не смогли считать строку с результатом из-за ошибки %v ", err.Error())
				Contexts[r.Context()]["err"] = "503-Ошибка бизнес-логики"
				return
			}

			events = append(events, Event)
		}
	}

	Contexts[r.Context()]["events"] = events
}

func EventCreate(event database.Event, r *http.Request) {
	querry := `
	INSERT INTO events (date, name,time)
	VALUES($1,$2,$3)`
	res, err := db.Exec(querry, event.Date, event.Name, event.Time)
	if err != nil {
		Contexts[r.Context()]["data"] = fmt.Sprintf("Не смоли создать событие под названием %v из-за ошибки:%v", event.Name, err.Error())
		Contexts[r.Context()]["err"] = "503-Ошибка бизнес-логики"
		return
	}
	id, _ := res.LastInsertId()
	Contexts[r.Context()]["data"] = fmt.Sprintf("Создано событие: %v под названием %v", id, event.Name)
}

func EventUpdate(event database.Event, r *http.Request) {
	if !exists(event.Id) {
		Contexts[r.Context()]["data"] = "события с id " + strconv.Itoa(event.Id) + " не существует"
		Contexts[r.Context()]["err"] = "400-события с id " + strconv.Itoa(event.Id) + " не существует"
		return
	}
	querry := `
	UPDATE events 
	SET date=?,name=?,time=?
	WHERE id=?
	`
	_, err := db.Exec(querry, event.Date, event.Name, event.Time, event.Id)
	if err != nil {
		Contexts[r.Context()]["data"] = fmt.Sprintf("Не смоли обновить событие %v из-за ошибки:%v", event.Id, err.Error())
		Contexts[r.Context()]["err"] = "503-Ошибка бизнес-логики"
		return
	}
	Contexts[r.Context()]["data"] = fmt.Sprintf("Обновлено событие: %v под названием %v", event.Id, event.Name)
}

func EventDelete(id int, r *http.Request) {
	if !exists(id) {
		Contexts[r.Context()]["data"] = "события с id " + strconv.Itoa(id) + " не существует"
		Contexts[r.Context()]["err"] = "400-события с id " + strconv.Itoa(id) + " не существует"
		return
	}
	querry := `
	DELETE FROM events 
	WHERE id=?
	`
	_, err := db.Exec(querry, id)
	if err != nil {
		Contexts[r.Context()]["data"] = fmt.Sprintf("Не смогли удалить событие %v из-за ошибки %v ", id, err.Error())
		Contexts[r.Context()]["err"] = "503-Ошибка бизнес-логики"
		return
	}
	Contexts[r.Context()]["data"] = fmt.Sprintf("Удалено событие: %v ", id)
}

func EventGetAll(r *http.Request) {
	querry := `
	SELECT * FROM events 
	`
	rows, err := db.Query(querry)
	if err != nil {
		Contexts[r.Context()]["data"] = fmt.Sprintf("Не смогли получить все записи из-за ошибки %v ", err.Error())
		Contexts[r.Context()]["err"] = "503-Ошибка бизнес-логики"
		return
	}
	var events []database.Event
	var Event database.Event
	for rows.Next() {
		err := rows.Scan(&Event.Id, &Event.Date, &Event.Name, &Event.Time)
		if err != nil {
			Contexts[r.Context()]["data"] = fmt.Sprintf("Не смогли считать строку с результатом из-за ошибки %v ", err.Error())
			Contexts[r.Context()]["err"] = "503-Ошибка бизнес-логики"
			return
		}
		events = append(events, Event)
	}
	Contexts[r.Context()]["events"] = events
}

func initContexts() {
	Contexts = make(map[context.Context]map[string]interface{})
}

func init() {
	db = NewDb()
	initContexts()
}

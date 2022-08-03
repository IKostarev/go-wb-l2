package service

import (
	"fmt"
	"log"
	"time"
)

func exists(id int) bool {
	querry := `
	SELECT COUNT(*) FROM events
	WHERE id=?
	`
	row := db.QueryRow(querry, id)
	var count int
	err := row.Scan(count)
	if err != nil {
		log.Fatal(err)
	}
	if count == 0 {
		return false
	}
	return true
}

func calcWeek(day string) []string {
	weekdays := make(map[string]time.Duration)
	weekdays["Monday"] = 0
	weekdays["Tuesday"] = 1
	weekdays["Wednesday"] = 2
	weekdays["Thursday"] = 3
	weekdays["Friday"] = 4
	weekdays["Saturday"] = 5
	weekdays["Sunday"] = 6
	date, _ := time.Parse("02.01.2006", day)
	date = date.Add(-weekdays[date.Weekday().String()] * 24 * time.Hour)
	var res []string
	for i := 0; i < 7; i++ {
		tmp := date.Format("02.01.2006")
		res = append(res, tmp)
		date = date.Add(24 * time.Hour)
	}
	return res
}

func calcMonth(day string) []string {
	firstDay := "01." + day[3:]
	fmt.Println(firstDay)
	date, _ := time.Parse("02.01.2006", firstDay)
	month := date.Month()
	var res []string
	for date.Month() == month {
		tmp := date.Format("02.01.2006")
		res = append(res, tmp)
		date = date.Add(24 * time.Hour)
	}
	return res
}

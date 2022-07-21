/*
Цепочка обязанностей — это поведенческий паттерн проектирования,
который позволяет передавать запросы последовательно по цепочке обработчиков.
Каждый последующий обработчик решает, может ли он обработать запрос сам и стоит ли передавать запрос дальше по цепи.

плюс - уменьшает зависимость между клиентом и обработчиками. Каждый обработчик сам выполняет свою логику независимо.
     - реализует принцип единственной обязанности.
	 - реализует принцип открытости и закрытости.
минус - запрос может остаться необработанным.
*/

package main

import "fmt"

type Service interface { // интерфейс сервиса
	Execute(*Data)
	SetNext(Service)
}

type Data struct { // структура данных
	GetSource    bool
	UpdateSource bool
}

type Device struct { // структура девайса
	Name string
	Next Service
}

func (d *Device) Execute(data *Data) { // реализация метода проверки данных
	if data.GetSource {
		fmt.Printf("Данные девайса %s уже переданы!!!\n", d.Name)
		d.Next.Execute(data)
		return
	}

	fmt.Printf("Данные для девайса %s\n", d.Name)
	data.GetSource = true
	d.Next.Execute(data)
}

func (d *Device) SetNext(serv Service) { // реализация метода передачи данных в сервис
	d.Next = serv
}

type UpdateDataService struct { // структура обновления данных сервиса
	Name string
	Next Service
}

func (u *UpdateDataService) Execute(data *Data) { // реализация метода обновления данных
	if data.UpdateSource {
		fmt.Printf("Данные девайса %s уже обновлены!!!\n", u.Name)
		u.Next.Execute(data)
		return
	}

	fmt.Printf("Обновленные данные для девайса %s\n", u.Name)
	data.UpdateSource = true
	u.Next.Execute(data)
}

func (u *UpdateDataService) SetNext(serv Service) { // реализация метода передачи обновленных данных в сервис
	u.Next = serv
}

type SaveService struct {
	Next Service
}

func (s *SaveService) Execute(data *Data) { // реализация метода сохранения данных
	if !data.UpdateSource {
		fmt.Println("Данные девайса не обновлены!!!\n")
		return
	}

	fmt.Println("Данные для девайса сохранены!\n")
}

func (s *SaveService) SetNext(serv Service) { // реализация метода сохранения данных в сервис
	s.Next = serv
}

func main() {
	device := &Device{Name: "Device_1"}

	updService := &UpdateDataService{Name: "Update Device_1"}

	saveService := &SaveService{}

	device.SetNext(updService)
	updService.SetNext(saveService)

	data := &Data{}
	device.Execute(data)
}

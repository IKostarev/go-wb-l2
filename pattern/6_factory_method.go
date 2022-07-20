/*  
Фабричный метод (англ. Factory Method) — порождающий шаблон проектирования,
предоставляющий подклассам (дочерним классам) интерфейс для создания экземпляров некоторого класса.
В момент создания наследники могут определить, какой класс создавать.
Иными словами, данный шаблон делегирует создание объектов наследникам родительского класса.

Плюс - Создание объектов независимо от их типов и сложности процесса создания.
Минусы - Даже для одного объекта необходимо создать соответствующую фабрику, что увеличивает код.
       - "Божественный" интерфейс от которого сложно будет отказаться в будущем!!!
*/

package main

import "fmt"

const (
	ServerType           = "server"
	PersonalComputerType = "personal_comp"
	NotebookType         = "notebook"
)

type Computer interface { // интерфейс компьютера
	GetType() string
	PrintDetails()
}

func New(typeName string) Computer {  // фабричный метод - общая реализация создания компьютера
	switch typeName {				  // в зависимости от переданного типа создаем конкретный компьютер
		case ServerType:
			return NewServer()
		case PersonalComputerType:
			return NewPersonalComputer()
		case NotebookType:
			return NewNotebook()
		default:
			fmt.Printf("Введенный тип: %s не найден!!!\n", typeName)
			return nil
	}
}

type Server struct { // структура сервер
	Type string
	Core int
	Memory int
	Monitor bool
}

func NewServer() Computer { // конструктор сервера
	return Server{
		Type: ServerType,
		Core: 16,
		Memory: 256,
		Monitor: false,
	}
}

func (s Server) GetType() string { // реализация структуры сервера
	return s.Type
}

func (s Server) PrintDetails() { // реализация структуры сервера
	fmt.Printf("Type - %s, core - %d, memory - %d, monitor - %v\n", s.Type, s.Core, s.Memory, s.Monitor)
}

type PersonalComputer struct { // структура персонального компьютера
	Type string
	Core int
	Memory int
	Monitor bool
}

func NewPersonalComputer() Computer { // конструктор персонального компьютера
	return PersonalComputer{
		Type: PersonalComputerType,
		Core: 8,
		Memory: 32,
		Monitor: true,
	}
}

func (p PersonalComputer) GetType() string { // реализация структуры персонального компьютера
	return p.Type
}

func (p PersonalComputer) PrintDetails() { // реализация структуры персонального компьютера
	fmt.Printf("Type - %s, core - %d, memory - %d, monitor - %v\n", p.Type, p.Core, p.Memory, p.Monitor)
}

type Notebook struct{ // структура ноутбука
	Type string
	Core int
	Memory int
	Monitor bool
}

func NewNotebook() Computer { // конструктор ноутбука
	return Notebook{
		Type: NotebookType,
		Core: 4,
		Memory: 8,
		Monitor: true,
	}
}

func (n Notebook) GetType() string { // реализация структуры ноутбука
	return n.Type
}

func (n Notebook) PrintDetails() { // реализация структуры ноутбука
	fmt.Printf("Type - %s, core - %d, memory - %d, monitor - %v\n", n.Type, n.Core, n.Memory, n.Monitor)
}

var types = []string{
	PersonalComputerType,
	NotebookType,
	ServerType,
	"phone",
}

func main() {
	for _, typeName := range types {
		comp := New(typeName)

		if comp == nil {
			continue
		}
		comp.PrintDetails()
	}
}

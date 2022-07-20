/*
Строитель (англ. Builder) — порождающий шаблон проектирования предоставляет способ создания составного объекта.

Плюс - позволяет создавать пошагово общий продукт который зависит от маленьких частей.
Минус - усложняет код из-за ведения новых структур либо интерфейсов.
      - А также клиент привязан к конкретному объекту строителя т.к. в интерфейсе может не быть какого-то метода и тогда есть необходимость добавлять его.
*/

package main

import "fmt"

const (
	DnsCollectorType   = "DNS"
	HonorCollectorType = "HONOR"
)

type Collector interface { // интерфейс сборки компьютера
	SetCore()              // добавление ядер
	SetBrand()             // добавление бренда
	SetMemory()            // добавление памяти
	SetMonitor()           // добавление монитора
	SetGraphicCard()       // добавление видеокарты
	GetComputer() Computer // возвращаем готовый компьютер
}

type Computer struct {
	Core        int
	Brand       string
	Monitor     string
	Memory      int
	GraphicCard int
}

func (c *Computer) Printer() { // печатаем компьютер
	fmt.Printf("Brand - %s, core - %v, memory - %v, graphic - %v, monitor - %s\n", c.Brand, c.Core, c.Memory, c.GraphicCard, c.Monitor)
}

func GetCollector(collectorType string) Collector {
	switch collectorType {
	case DnsCollectorType:
		return &DnsCollector{}
	case HonorCollectorType:
		return &HonorCollector{}
	default:
		fmt.Printf("Введенный тип: %s не найден!!!\n", collectorType)
		return nil
	}
}

type DnsCollector struct { // структура компьютера ДНС
	Core        int
	Brand       string
	Monitor     string
	Memory      int
	GraphicCard int
}

//----------РЕАЛИЗАЦИЯ ИНТЕРФЕЙСА ДЛЯ ДНС КОМПЬЮТЕРА----------
func (d *DnsCollector) SetCore() {
	d.Core = 4
}

func (d *DnsCollector) SetBrand() {
	d.Brand = "DNS"
}

func (d *DnsCollector) SetMemory() {
	d.Memory = 4
}

func (d *DnsCollector) SetMonitor() {
	d.Monitor = "DNS_DEFAUTL_MONITOR"
}

func (d *DnsCollector) SetGraphicCard() {
	d.GraphicCard = 1
}

func (d *DnsCollector) GetComputer() Computer { // конструктор базового компьютера DNS
	return Computer{
		Core:        d.Core,
		Brand:       d.Brand,
		Memory:      d.Memory,
		Monitor:     d.Monitor,
		GraphicCard: d.GraphicCard,
	}
}

type HonorCollector struct { // структура компьютера HONOR
	Core        int
	Brand       string
	Monitor     string
	Memory      int
	GraphicCard int
}

//----------РЕАЛИЗАЦИЯ ИНТЕРФЕЙСА ДЛЯ HONOR КОМПЬЮТЕРА----------
func (h *HonorCollector) SetCore() {
	h.Core = 8
}

func (h *HonorCollector) SetBrand() {
	h.Brand = "HONOR"
}

func (h *HonorCollector) SetMemory() {
	h.Memory = 8
}

func (h *HonorCollector) SetMonitor() {
	h.Monitor = "HONOR_DEFAUTL_MONITOR"
}

func (h *HonorCollector) SetGraphicCard() {
	h.GraphicCard = 2
}

func (h *HonorCollector) GetComputer() Computer { // конструктор базового компьютера HONOR
	return Computer{
		Core:        h.Core,
		Brand:       h.Brand,
		Memory:      h.Memory,
		Monitor:     h.Monitor,
		GraphicCard: h.GraphicCard,
	}
}

type Factory struct {
	Collector Collector
}

func NewFactory(collector Collector) *Factory { // конструктор фабрики
	return &Factory{
		Collector: collector,
	}
}

func (f *Factory) SetCollector(collector Collector) { // меняет поведение завода в зависимости от того, что мы передаем
	f.Collector = collector
}

func (f *Factory) CreateComputer() Computer { // основная функция по "строительству" компьютера
	f.Collector.SetCore()
	f.Collector.SetBrand()
	f.Collector.SetMemory()
	f.Collector.SetMonitor()
	f.Collector.SetGraphicCard()

	return f.Collector.GetComputer() // возвращаем собранный компьютер
}

func main() {
	dns := GetCollector("DNS")
	honor := GetCollector("HONOR")

	factory := NewFactory(dns)
	dnsComputer := factory.CreateComputer()
	dnsComputer.Printer()

	factory.SetCollector(honor)
	honorComputer := factory.CreateComputer()
	honorComputer.Printer()
}

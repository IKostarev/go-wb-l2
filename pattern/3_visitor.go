/*
Посетитель (англ. visitor) — поведенческий шаблон проектирования,
описывающий операцию, которая выполняется над объектами других классов.
При изменении visitor нет необходимости изменять обслуживаемые классы.

плюс - новая функциональность в несколько классов добавляется сразу, не изменяя код этих классов.
минус - при изменении обсулживаемого класса нужно поменять код у шаблона.
      - Затруднение добавления новых классов поскольку нужно обновлять иерархию посетителя и его сыновый
*/

package main

import "fmt"

//----------INTERFACES----------
type ProductInfoRetriever interface {
	GetPrice() float32
	GetName() string
}

type Visitor interface {
	Visit(ProductInfoRetriever)
}

type Visitable interface {
	Accept(Visitor)
}

//----------PRODUCT----------
type Product struct {
	Price float32
	Name  string
}

func (p *Product) GetPrice() float32 {
	return p.Price
}

func (p *Product) Accept(v Visitor) {
	v.Visit(p)
}

func (p *Product) GetName() string {
	return p.Name
}

//----------PRODUCTS----------
type Rice struct {
	Product
}

type Pasta struct {
	Product
}

//----------VISITOR----------
type PriceVisitor struct {
	Sum float32
}

func (pv *PriceVisitor) Visit(p ProductInfoRetriever) {
	pv.Sum += p.GetPrice()
}

type NamePrinter struct {
	ProductList string
}

func (n *NamePrinter) Visit(p ProductInfoRetriever) {
	n.ProductList = fmt.Sprintf("%s\n%s", p.GetName(), n.ProductList)
}

func main() {
	products := make([]Visitable, 2)
	products[0] = &Rice{
		Product: Product{
			Price: 500,
			Name:  "Standart rice",
		},
	}
	products[1] = &Pasta{
		Product: Product{
			Price: 300,
			Name:  "France pastas",
		},
	}

	priceVisitor := &PriceVisitor{}

	for _, p := range products {
		p.Accept(priceVisitor)
	}

	fmt.Printf("Total: %f\n", priceVisitor.Sum)

	nameVisitor := &NamePrinter{}

	for _, p := range products {
		p.Accept(nameVisitor)
	}

	fmt.Printf("\nProduct list:\n----------\n%s", nameVisitor.ProductList)
}

/*
Шаблон фасад (англ. Facade) — структурный шаблон проектирования,
позволяющий скрыть сложность системы путём сведения всех возможных внешних вызовов к одному объекту,
делегирующему их соответствующим объектам системы

плюс - изоляция клиентов от поведения сложной системы
минус - сам интерфейс фасада может стать супер-классом (привязка к объекту), все последующие функции будут привязаны
*/
package main

import (
	"errors"
	"fmt"
	"time"
)

type Product struct { // продукты в магазине
	Title string
	Price float64
}

type Shop struct { // магазин
	Title    string
	Products []Product
}

// САМ ФАСАД РЕАЛИЗОВАН В МАГАЗИНЕ - СВЯЗКА ЮЗЕРА, БАНКА, КАРТЫ И ПРОДУКТА
func (s Shop) Sell(user User, product string) error { // покупка товара
	fmt.Printf("МАГАЗИН - Когда происходит покупка, магазин запрашивает наличие денег у пользователя\n")
	time.Sleep(2 * time.Second) // происходит запрос денег у пользователя

	err := user.Card.CheckBalance()

	if err != nil {
		return err
	}

	fmt.Printf("МАГАЗИН - Проверка может ли пользователь: %s купить товар\n", user.Name)
	time.Sleep(2 * time.Second)

	for _, prod := range s.Products {
		if prod.Title != product {
			continue // если нет товара, пропускаем его
		}

		if prod.Price > user.GetBalance() {
			return errors.New("МАГАЗИН - У пользователя недостаточно средств для покупки")
		}

		fmt.Printf("МАГАЗИН - Товар %s куплен пользователем - %s\n", s.Title, user.Name)
	}

	return nil
}

type Bank struct { // банк
	Title string
	Cards []Card
}

func (b Bank) CheckBalance(cardNum string) error {
	fmt.Printf("БАНК - Происходит проверка баланса на карте: %s\n", cardNum)
	time.Sleep(2 * time.Second) // происходит проверка баланса

	for _, card := range b.Cards {
		if card.Title != cardNum {
			continue // если не наша карта, пропускаем её
		}

		if card.Balance <= 0 {
			return errors.New("БАНК - Нет средств для совершения покупки")
		}
	}

	fmt.Printf("БАНК - Деньги есть!!!\n")

	return nil
}

type Card struct { // карта банка
	Title   string
	Bank    *Bank
	Balance float64
}

func (c Card) CheckBalance() error {
	fmt.Printf("КАРТА - Происходит запрос в банк для проверки средств\n")
	time.Sleep(2 * time.Second) // происходит запрос в банк
	return c.Bank.CheckBalance(c.Title)
}

type User struct { // пользователь
	Name string
	Card *Card
}

func (u User) GetBalance() float64 { // пользователь получает баланс
	return u.Card.Balance
}

var (
	bank = Bank{
		Title: "Tinkoff",
		Cards: []Card{},
	}

	testCard1 = Card{
		Title:   "Card_1",
		Bank:    &bank,
		Balance: 5000,
	}

	testCard2 = Card{
		Title:   "Card_2",
		Bank:    &bank,
		Balance: 1,
	}

	user1 = User{
		Name: "User_1",
		Card: &testCard1,
	}

	user2 = User{
		Name: "User_2",
		Card: &testCard2,
	}

	testProd = Product{
		Title: "Phone",
		Price: 400,
	}

	testShop = Shop{
		Title: "Shop_name",
		Products: []Product{
			testProd,
		},
	}
)

func main() {
	fmt.Printf("БАНК - Выпускает карты: %v и %v\n", testCard1, testCard2)
	bank.Cards = append(bank.Cards, testCard1, testCard2) // Добавляем наши тестовые карты

	fmt.Printf("Пользователь: %s\n", user1.Name)    // Первый пользователь
	err := testShop.Sell(user1, testProd.Title)   // Первый пользователь хочет совершить покупку тестового продукта

	if err != nil {
		fmt.Println(err.Error())   // проверка
		return
	}

	fmt.Printf("Пользователь: %s\n", user2.Name)    // Второй пользователь
	err = testShop.Sell(user2, testProd.Title)    // Второй пользователь хочет совершить покупку тестового продукта

	if err != nil {
		fmt.Println(err.Error())  // проверка
		return
	}
}

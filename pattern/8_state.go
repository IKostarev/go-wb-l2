/*
Состояние (англ. state) — поведенческий шаблон проектирования.
Используется в тех случаях, когда во время выполнения программы объект должен менять своё поведение в зависимости от своего состояния.

плюс - упрощает код
минус - может неоправданно усложнить код, если состояний мало и они редко меняются.
*/

package main

import (
	"fmt"
	"log"
)

type State interface {
	AddItem(int) error
	RequestItem() error
	InsertMoney(money int) error
	Dispenseitem() error
}

type VendMachine struct {
	hasItem       State
	itemRequested State
	hasMoney      State
	noItem        State
	currentState  State
	itemCount     int
	itemPrice     int
}

func NewVendingMachine(itemCount, itemPrice int) *VendMachine {
	v := &VendMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}

	hasItemState := &HasMoneyState{
		vendMachine: v,
	}

	itemRequestState := &ItemRequestState{
		vendMachine: v,
	}

	hasMoneyState := &HasMoneyState{
		vendMachine: v,
	}

	noItemState := &NoItemState{
		vendMachine: v,
	}

	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState

	return v
}

func (v *VendMachine) RequestItem() error {
	return v.currentState.RequestItem()
}

func (v *VendMachine) AddItem(count int) error {
	return v.currentState.AddItem(count)
}

func (v *VendMachine) InsertMoney(money int) error {
	return v.currentState.InsertMoney(money)
}

func (v *VendMachine) Dispenseitem() error {
	return v.currentState.Dispenseitem()
}

func (v *VendMachine) setState(s State) {
	v.currentState = s
}

func (v *VendMachine) incrementItemCount(count int) {
	v.itemCount = count
}

type NoItemState struct {
	vendMachine *VendMachine
}

func (no *NoItemState) RequestItem() error {
	return fmt.Errorf("Item out of stock")
}

func (no *NoItemState) AddItem(count int) error {
	no.vendMachine.incrementItemCount(count)
	no.vendMachine.setState(no.vendMachine.hasItem)
	return nil
}

func (no *NoItemState) InsertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

func (no *NoItemState) Dispenseitem() error {
	return fmt.Errorf("Item out of stock")
}

type ItemRequestState struct {
	vendMachine *VendMachine
}

func (i *ItemRequestState) RequestItem() error {
	return fmt.Errorf("Item already requested")
}

func (i *ItemRequestState) AddItem() error {
	return fmt.Errorf("Item Dispense in progress")
}

func (i *ItemRequestState) InsertMoney(money int) error {
	if money < i.vendMachine.itemPrice {
		fmt.Errorf("Insert money is less. Please insert %d", i.vendMachine.itemPrice)
	}

	fmt.Println("Money entered is ok")
	i.vendMachine.setState(i.vendMachine.hasMoney)
	return nil
}

func (i *ItemRequestState) Dispenseitem() error {
	return fmt.Errorf("Please insert money first")
}

type HasMoneyState struct {
	vendMachine *VendMachine
}

func (h *HasMoneyState) RequestItem() error {
	return fmt.Errorf("Item dispense in progress")
}

func (h *HasMoneyState) AddItem(count int) error {
	return fmt.Errorf("Item dispense in progress")
}

func (h *HasMoneyState) InsertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

func (h *HasMoneyState) Dispenseitem() error {
	fmt.Println("Dispensing Item")
	h.vendMachine.itemCount = h.vendMachine.itemCount - 1

	if h.vendMachine.itemCount == 0 {
		h.vendMachine.setState(h.vendMachine.noItem)
	} else {
		h.vendMachine.setState(h.vendMachine.hasItem)
	}
	return nil
}

type HasItemState struct {
	vendMachine *VendMachine
}

func (hi *HasItemState) RequestItem() error {
	if hi.vendMachine.itemCount == 1 {
		hi.vendMachine.setState(hi.vendMachine.noItem)
		return fmt.Errorf("No item present")
	}
	fmt.Printf("Item requestd\n")
	hi.vendMachine.setState(hi.vendMachine.itemRequested)
	return nil
}

func (hi *HasItemState) AddItem(count int) error {
	fmt.Printf("%d items added\n", count)
	hi.vendMachine.incrementItemCount(count)
	return nil
}

func (hi *HasItemState) InsertMoney(money int) error {
	return fmt.Errorf("Please select item first")
}

func (hi *HasItemState) Dispenseitem() error {
	return fmt.Errorf("Please select item first")
}

func main() {
	vedingMachine := NewVendingMachine(1, 10)

	err := vedingMachine.RequestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vedingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vedingMachine.Dispenseitem()
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println()

	err = vedingMachine.AddItem(2)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println()

	err = vedingMachine.RequestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vedingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vedingMachine.Dispenseitem()
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println()
}

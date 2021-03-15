package order

import "fmt"

type Order struct {
	Id     int
	Name   string
	Amount int
	Tax    int
}

func (v Order) CalculateTotal() {
	fmt.Println(v.Amount - v.Tax) // Todo: return calculated amount don't print it..
}

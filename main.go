package main

import (
	"fmt"

	"github.com/muhkj7/helloworld/app"
	"github.com/muhkj7/helloworld/domain"
	"github.com/muhkj7/helloworld/order"
)

func main() {
	u := domain.User{
		Firstname: "Muhammad",
		Lastname:  "Khan",
		Email:     "muhkj7@gmail.com",
	}
	fmt.Println(u.Email)

	u.SayHello()

	v := order.Order{
		Id:     12,
		Name:   "Khan",
		Amount: 205,
		Tax:    1500,
	}
	fmt.Println(v.Tax)

	v.CalculateTotal()

	fmt.Println(app.ServiceCalled())
}

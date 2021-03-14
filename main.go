package main

import (
	"fmt"

	"github.com/muhkj7/helloworld/app"
	"github.com/muhkj7/helloworld/domain"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println(domain.SayHello())
	fmt.Println(app.ServiceCalled())
}

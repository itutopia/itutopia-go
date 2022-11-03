package interfaces

import (
	"fmt"
)

type Phone interface {
	call()
}

type NokiaPhone struct {
}
type IPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func CallExample() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call() // I am Nokia, I can call you!
	phone = new(IPhone)
	phone.call() // I am iPhone, I can call you!
}

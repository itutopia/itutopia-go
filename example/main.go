package main

import (
	"app/example/go_example/constant"
	"testing"
)

func main() {
	//fmt.Println("hello go")

	//pointer.PointerGetValue()
	//array.ArrayExample()

	t := new(testing.T)
	constant.TestConstant(*&t)

}

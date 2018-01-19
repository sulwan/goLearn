package main

import (
	"fmt"

	"github.com/sulwan/goLearn/until"
)

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	until.Modify(array)
	fmt.Println("In main(), array values:", array)
}

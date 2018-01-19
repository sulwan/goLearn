package until

import (
	"fmt"
)

func Modify(array [5]int) {
	array[0] = 10
	fmt.Println("In modify(), array value:", array)
}

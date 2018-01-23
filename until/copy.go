package until

import (
	"fmt"
)

var (
	dataSlice    []int
	dataSliceTwo []int
)

func CopyData() {
	dataSlice := []int{1, 2, 3, 4, 5, 6}
	dataSliceTwo := []int{11, 22, 33}
	copy(dataSlice, dataSliceTwo)
	fmt.Println(dataSlice)

}

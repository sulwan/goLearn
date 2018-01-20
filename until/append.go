package until

import (
	"fmt"
)

func AppendDataSlice() {
	dataSlice := make([]int, 5, 10)
	fmt.Println(append(dataSlice, 6, 7, 8))
}

func AppendDataSliceArray() {
	dataSlice := make([]int, 5, 10)
	dataSli := []int{6, 7, 8}
	fmt.Println(append(dataSlice, dataSli...))
}

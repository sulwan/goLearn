package until

import (
	"fmt"
)

var dataSlice []int

func LenSlice() {
	dataSlice = make([]int, 5, 10)
	fmt.Println(len(dataSlice))
}

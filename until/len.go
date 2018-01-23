package until

import (
	"fmt"
)

func LenSlice() {
	dataSlice := make([]int, 5, 10)
	fmt.Println(len(dataSlice))
}

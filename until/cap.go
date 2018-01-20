package until

import (
	"fmt"
)

func CapExp() {
	var (
		sliceData [10]int
	)
	sliceData = [10]int{1, 2, 3, 4, 5, 6}
	fmt.Println(cap(sliceData))
}

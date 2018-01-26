package until

import (
	"fmt"
)

func GotoUntil() {
	i := 0
HERE:
	fmt.Println(i)
	i++
	if i < 10 {
		goto HERE
	}
}

package until

import (
	"fmt"
)

func IndefiniteUntil(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

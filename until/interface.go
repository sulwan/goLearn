package until

import (
	"fmt"
)

func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		default:
			fmt.Println(arg, "is an unknow type.")
		}
	}
}

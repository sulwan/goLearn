package until

import (
	"fmt"
)

func SliceArray() {
	var (
		dateArray [10]int
		dateSlice []int
	)
	dateArray = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	dateSlice = dateArray[:5]
	fmt.Println("Elements of dataArray:")
	for _, v := range dateArray {
		fmt.Print(v, " ")
	}

	fmt.Println("\n Elements of dateSlice: ")

	for _, v := range dateSlice {
		fmt.Print(v, " ")
	}

	fmt.Println()

}

func Vslicek() {
	var (
		sliceData []int
	)
	sliceData = []int{1, 2, 3, 4, 5, 6, 7, 8}
	for i := 0; i < len(sliceData); i++ {
		fmt.Println("sliceData[", i, "] =", sliceData[i])
	}
}

func Vslicekv() {
	var (
		sliceData []int
	)
	sliceData = []int{1, 2, 3, 4, 5, 6}

	for i, v := range sliceData {
		fmt.Println("sliceData[", i, "] =", v)
	}

}

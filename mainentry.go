package main

import (
	"fmt"

	"github.com/snehasis007/corelib2/base"
)

// func random() int {
// 	return rand.Intn(MAX-MIN) + MIN
// }

func main() {
	//var x base.AST = base.AST{}

	// for _, char := range []rune{0x56, 0x34} {
	// 	fmt.Println(string(char))
	// }
	//fmt.Printf("%d %#04x %U '%c'\n", 0x3A6, 934, '\u03A6', '\U000003A7')
	//arr := []int{1, 23, 13, 10}

	arr2 := []base.DataItemImpl{{Index: 0, Value: 9}, {Index: 1, Value: 7}, {Index: 2, Value: 5}, {Index: 3, Value: 1}}
	//base.QuickSort(arr2, 0, len(arr2)-1)
	base.Con()
	fmt.Println(arr2)
	//fmt.Printf("%d")

	//http.Start()
	//https.Handler()

}

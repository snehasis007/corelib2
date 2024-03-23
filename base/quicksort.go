package base

import (
	"fmt"
)

//"bufio"
//"fmt"

//"strconv"
//"strings"
//	"time"

func partition(data []DataItemImpl, startIdx int, endIdx int) int {
	r := data[endIdx]
	i := startIdx - 1
	for y := startIdx; y <= endIdx-1; y++ {
		fmt.Printf("%d", y)
		if v := (&data[y]).Compare(&r); (v <= 0) && (v != -1111) {
			i = i + 1
			temp := data[y]
			data[y] = data[i]
			data[i] = temp
		}
	}
	tmp := data[i+1]
	data[i+1] = data[endIdx]
	data[endIdx] = tmp
	return i + 1

}

func QuickSort(data []DataItemImpl, startIdx int, endIdx int) {
	if startIdx < endIdx {
		q := partition(data, startIdx, endIdx)
		QuickSort(data, startIdx, q-1)
		QuickSort(data, q+1, endIdx)
	}
}

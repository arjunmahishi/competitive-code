// Problem link: https://www.codewars.com/kata/51ba717bb08c1cd60f00002f/train/go

package kata

import (
	"strconv"
	"strings"
)

func Solution(list []int) string {
	cont := []int{}
	finalString := ""
	for i := 0; i < len(list); i++ {
		cont = append(cont, list[i])
		if i != len(list)-1 && list[i]+1 == list[i+1] {
			continue
		}
		finalString += formatRange(cont)
		cont = []int{}
	}
	return strings.TrimRight(finalString, ",")
}

func formatRange(arr []int) string {
	if len(arr) >= 3 {
		return strconv.Itoa(arr[0]) + "-" + strconv.Itoa(arr[len(arr)-1]) + ","
	}
	str := ""
	for _, ele := range arr {
		str += strconv.Itoa(ele) + ","
	}
	return str
}

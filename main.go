package main

import "fmt"

//"test/simple"

func qStort(arr []int) {
	if len(arr) < 2 {
		return
	}

	idx := partition(arr)
	qStort(arr[:idx])
	qStort(arr[idx+1:])
}

func partition(arr []int) int {
	p := arr[len(arr)-1]
	i := 0
	for j := 0; j < len(arr); j++ {
		if arr[j] < p {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[len(arr)-1] = arr[len(arr)-1], arr[i]
	return i
}

func main() {
	a := []int{1, 3, 9, 34, 6, 8, 5}
	qStort(a)
	fmt.Println(a)

}

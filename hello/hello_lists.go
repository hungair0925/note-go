package main

import "fmt"

func main() {
	var hoge [3]int
	hoge[0] = 100
	hoge[1] = 200
	//[100 200 0]
	fmt.Println(hoge)
	//this code is not working. => ppend(hoge, 300)

	var fuga []int = []int{100, 200}
	//[100 200]
	fmt.Println(fuga)
	fuga = append(fuga, 300)
	//[100 200 300]
	fmt.Println(fuga)

	var nestedList [3][3]int = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(nestedList)

	var nestedSlice = [][]int{
		[]int{0, 1},
		[]int{2, 3},
		[]int{4, 5, 6},
	}
	fmt.Println(nestedSlice)
}

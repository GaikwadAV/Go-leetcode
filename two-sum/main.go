package main

import "fmt"

func twoSum(nums []int, target int) []int {
	mymap := make(map[int]int)

	for i, v := range nums {
		x := target - v

		found, ok := mymap[x]
		if ok {
			return []int{found, i}
		}
		mymap[v] = i

	}
	return nil
}
func main() {
	nums := []int{1, 0, 2, 7}
	target := 9
	r := twoSum(nums, target)
	fmt.Println(r)
}

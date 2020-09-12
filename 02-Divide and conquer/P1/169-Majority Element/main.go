package main

func majorityElement(nums []int) int {
	//数组非空。且数组中总是有多数元素。
	maps := map[int]int{}
	for _, num := range nums {
		_, ok := maps[num]
		if ok {
			maps[num]++
		} else {
			maps[num] = 1
		}
	}
	var res = 0
	for k, v := range maps {
		if v > len(nums)/2 {
			res = k
		}
	}
	return res
}
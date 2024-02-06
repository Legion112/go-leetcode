package TwoSum

import "log"

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int, len(nums))
	for index, value := range nums {
		hashmap[value] = index
	}
	for firstIndex, firstValue := range nums {
		if secondIndex, exist := hashmap[target-firstValue]; exist && secondIndex != firstIndex {
			return []int{firstIndex, secondIndex}
		}
	}
	log.Println("No value to return")
	return nil
}

package main

func twoSum(numbers []int, target int) [2]int {
	var ix, iy int
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target && i != j {
				ix = i
				iy = j
				return [2]int{ix, iy}
			}
		}
	}
	return [2]int{ix, iy}
}

func main() {

}

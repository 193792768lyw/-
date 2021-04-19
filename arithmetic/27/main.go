package main

func main() {

}

func removeElement(nums []int, val int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			continue
		} else {
			nums[res] = nums[i]
			res++
		}
	}
	return res
}

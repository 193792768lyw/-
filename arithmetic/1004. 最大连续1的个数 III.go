package main

func main() {
	temp := map[string]interface{}{}
	delete(temp, "kkk")

	//fmt.Println(longestOnes([]int{1,1,1,0,0,0,1,1,1,1,0}, 2))
}
func longestOnes(A []int, K int) int {
	res := 0
	for i := 0; i < len(A); i++ {
		tempK := K
		temp := 0
		for j := i; j < len(A); j++ {
			if A[j] == 0 {
				if tempK == 0 {
					break
				}
				temp++
				tempK--
			} else {
				temp++
			}

		}
		if temp > res {
			res = temp
		}
	}

	return res
}

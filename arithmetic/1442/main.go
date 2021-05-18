package main

func main() {

}

//func countTriplets(arr []int) (ans int) {
//	n := len(arr)
//	s := make([]int, n+1)
//	for i, val := range arr {
//		s[i+1] = s[i] ^ val
//	}
//	for i := 0; i < n; i++ {
//		for j := i + 1; j < n; j++ {
//			for k := j; k < n; k++ {
//				if s[i] == s[k+1] {
//					ans++
//				}
//			}
//		}
//	}
//	return
//}

func countTriplets(arr []int) (ans int) {
	n := len(arr)
	s := make([]int, n+1)
	for i, val := range arr {
		s[i+1] = s[i] ^ val
	}
	for i := 0; i < n; i++ {
		for k := i + 1; k < n; k++ {
			if s[i] == s[k+1] {
				ans += k - i
			}
		}
	}
	return
}

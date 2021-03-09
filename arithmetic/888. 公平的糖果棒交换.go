package main

func main() {

}

func fairCandySwap(a, b []int) []int {
	sumA := 0
	// 将 AA 中的数字存入哈希表
	setA := map[int]struct{}{}
	for _, v := range a {
		sumA += v
		setA[v] = struct{}{}
	}

	sumB := 0
	for _, v := range b {
		sumB += v
	}

	delta := (sumA - sumB) / 2
	for i := 0; i < len(b); i++ {
		y := b[i]
		x := y + delta
		if _, has := setA[x]; has {
			return []int{x, y}
		}
	}
	return nil
}

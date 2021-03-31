package main

import "fmt"

// 全排列
func main() {
	res := make([][]int, 0)

	var fun func(chs []int, start int)
	fun = func(chs []int, start int) {
		if start == len(chs)-1 {
			res = append(res, append([]int(nil), chs...))
			//如果已经到了数组的最后一个元素，前面的元素已经排好，输出。
		}
		for i := start; i <= len(chs)-1; i++ {
			if i == start || chs[i] != chs[start] {
				//把第一个元素分别与后面的元素进行交换，递归的调用其子数组进行排序
				chs[i], chs[start] = chs[start], chs[i]
				fun(chs, start+1)
				chs[i], chs[start] = chs[start], chs[i]
				//子数组排序返回后要将第一个元素交换回来。
				//如果不交换回来会出错，比如说第一次1、2交换，第一个位置为2，子数组排序返回后如果不将1、2
				//交换回来第二次交换的时候就会将2、3交换，因此必须将1、2交换使1还是在第一个位置
			}

		}
	}
	fun([]int{1, 1}, 0)
	fmt.Println(res)
}

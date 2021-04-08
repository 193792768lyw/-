package main

import "time"

func main() {
	//var wg interface{}
	//wg = &sync.WaitGroup{}
	//
	//fmt.Println(reflect.TypeOf(wg))
	//done := make(chan bool)
	//go func() {
	//	select {
	//	case done <- true:
	//	default:
	//		fmt.Println("bdfvfhj")
	//		return
	//	}
	//
	//}()
	////for  {
	////	select {
	////	case  h := <- ch :
	////		fmt.Println(h)
	////	}
	////
	////}
	//time.Sleep(5*time.Second)
	//for v := range done{
	//	fmt.Println(v)
	//}
	//time.Sleep(5*time.Second)
	////fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 3))
	////fmt.Println(SafeClose(nil))
	ch := make(chan int)
	go func() {
		time.Sleep(1 * time.Second)
		select {
		case <-time.After(time.Second * 2):
			//app.G.Log.WithFields(c.GetLoggerFields("logger_fields")).WithField("data", v).Error("发送数据错误")
			return
		case ch <- 88:
		}
	}()
	close(ch)
	time.Sleep(5 * time.Second)

}

func SafeClose(ch chan int) (justClosed bool) {
	defer func() {
		if recover() != nil {
			// 一个函数的返回结果可以在defer调用中修改。
			justClosed = true
		}
	}()

	// 假设ch != nil。
	close(ch)   // 如果 ch 已关闭，将 panic
	return true // <=> justClosed = true; return
}

/*
输入：nums = [2,5,6,0,0,1,2], target = 3
输出：false
*/
func search(nums []int, target int) bool {
	len := len(nums)
	for _, num := range nums {
		if num == target {
			return true
		}
		if num > target {
			break
		}
	}

	for i := len - 1; i >= 0; i-- {
		if nums[i] == target {
			return true
		}
		if nums[i] < target {
			break
		}
	}
	return false
}

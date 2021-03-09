package main

import "fmt"

func main() {
	myQueue := Constructor()
	myQueue.Push(1)              // queue is: [1]
	myQueue.Push(2)              // queue is: [1, 2] (leftmost is front of the queue)
	fmt.Println(myQueue.Peek())  // return 1
	fmt.Println(myQueue.Pop())   // return 1, queue is [2]
	fmt.Println(myQueue.Empty()) // return false

}

/*
将一个栈当作输入栈，用于压入 \texttt{push}push 传入的数据；另一个栈当作输出栈，用于 \texttt{pop}pop 和 \texttt{peek}peek 操作。

每次 \texttt{pop}pop 或 \texttt{peek}peek 时，若输出栈为空则将输入栈的全部数据依次弹出并压入输出栈，这样输出栈从栈顶往栈底的顺序就是队列从队首往队尾的顺序。

*/

type MyQueue struct {
	inStack, outStack []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(x int) {
	q.inStack = append(q.inStack, x)
}

func (q *MyQueue) in2out() {
	for len(q.inStack) > 0 {
		q.outStack = append(q.outStack, q.inStack[len(q.inStack)-1])
		q.inStack = q.inStack[:len(q.inStack)-1]
	}
}

func (q *MyQueue) Pop() int {
	if len(q.outStack) == 0 {
		q.in2out()
	}
	x := q.outStack[len(q.outStack)-1]
	q.outStack = q.outStack[:len(q.outStack)-1]
	return x
}

func (q *MyQueue) Peek() int {
	if len(q.outStack) == 0 {
		q.in2out()
	}
	return q.outStack[len(q.outStack)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.inStack) == 0 && len(q.outStack) == 0
}

//type MyQueue struct {
//	arr []int
//}
//
///** Initialize your data structure here. */
//func Constructor() MyQueue {
//	return MyQueue{arr: []int{}}
//}
//
///** Push element x to the back of queue. */
//func (this *MyQueue) Push(x int) {
//	this.arr = append(this.arr, x)
//}
//
///** Removes the element from in front of queue and returns that element. */
//func (this *MyQueue) Pop() int {
//	res := this.arr[0]
//	this.arr = this.arr[1:]
//	return res
//}
//
///** Get the front element. */
//func (this *MyQueue) Peek() int {
//	return this.arr[0]
//}
//
///** Returns whether the queue is empty. */
//func (this *MyQueue) Empty() bool {
//	return len(this.arr) == 0
//}
//
///**
// * Your MyQueue object will be instantiated and called as such:
// * obj := Constructor();
// * obj.Push(x);
// * param_2 := obj.Pop();
// * param_3 := obj.Peek();
// * param_4 := obj.Empty();
// */

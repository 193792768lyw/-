package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRange(t *testing.T) {
	// arr
	words1 := [...]string{"Go", "语言", "高性能", "编程"}
	for i, s := range words1 {
		fmt.Println(i, s)
	}
	// slice
	/*
		变量 words 在循环开始前，仅会计算一次，如果在循环中修改切片的长度不会改变本次循环的次数。
		迭代过程中，每次迭代的下标和值被赋值给变量 i 和 s，第二个参数 s 是可选的。
		针对 nil 切片，迭代次数为 0。
	*/
	words := []string{"Go", "语言", "高性能", "编程"}
	for i, s := range words {
		words = append(words, "test")
		fmt.Println(i, s)
	}
	// range 还有另一种只遍历下标的写法，这种写法与 for 几乎没什么差异了。
	for i := range words {
		fmt.Println(i, words[i])
	}
	fmt.Println("--------------------")
	// map
	/*
		和切片不同的是，迭代过程中，删除还未迭代到的键值对，则该键值对不会被迭代。
		在迭代过程中，如果创建新的键值对，那么新增键值对，可能被迭代，也可能不会被迭代。
		针对 nil 字典，迭代次数为 0
		有可能打印如下
		two: 2
		three: 3
		one: 1
	*/
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	for k, v := range m {
		delete(m, "two")
		m["four"] = 4
		fmt.Printf("%v: %v\n", k, v)
	}
	fmt.Println("--------------------")
	// channel
	/*
		发送给信道(channel) 的值可以使用 for 循环迭代，直到信道被关闭。
		如果是 nil 信道，循环将永远阻塞。
	*/
	ch := make(chan string)
	go func() {
		ch <- "Go"
		ch <- "语言"
		ch <- "高性能"
		ch <- "编程"
		close(ch)
	}()
	for n := range ch {
		fmt.Println(n)
	}
}

/*
我们可以用一个非常简单的例子来证明 range 迭代时，返回的是拷贝。
persons 是一个长度为 3 的切片，每个元素是一个结构体。
使用 range 迭代时，试图将每个结构体的 no 字段增加 10，但修改无效，因为 range 返回的是拷贝。
使用 for 迭代时，将每个结构体的 no 字段增加 100，修改有效。
*/
func TestModifyForAndRange(t *testing.T) {
	persons := []struct{ no int }{{no: 1}, {no: 2}, {no: 3}}
	for _, s := range persons {
		s.no += 10
	}
	for i := 0; i < len(persons); i++ {
		persons[i].no += 100
	}
	fmt.Println(persons)
}

func TestRangeCopy(t *testing.T) {
	persons := []*struct{ no int }{{no: 1}, {no: 2}, {no: 3}}
	for _, s := range persons {
		s.no += 10
	}
	for i := 0; i < len(persons); i++ {
		persons[i].no += 100
	}
	for _, v := range persons {
		fmt.Println(v.no)
	}
}

func generateWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func BenchmarkForIntSlice(b *testing.B) {
	nums := generateWithCap(1024 * 1024)
	for i := 0; i < b.N; i++ {
		len := len(nums)
		var tmp int
		for k := 0; k < len; k++ {
			tmp = nums[k]
		}
		_ = tmp
	}
}

func BenchmarkRangeIntSlice(b *testing.B) {
	nums := generateWithCap(1024 * 1024)
	for i := 0; i < b.N; i++ {
		var tmp int
		for _, num := range nums {
			tmp = num
		}
		_ = tmp
	}
}

/*

goos: windows
goarch: amd64
pkg: Study/example/hpg-range
BenchmarkForIntSlice-4              3505            331621 ns/op
BenchmarkRangeIntSlice-4            3475            352602 ns/op

generateWithCap 用于生成长度为 n 元素类型为 int 的切片。
从最终的结果可以看到，遍历 []int 类型的切片，for 与 range 性能几乎没有区别。
*/

type Item struct {
	id  int
	val [4096]byte
}

func BenchmarkForStruct(b *testing.B) {
	var items [1024]Item
	for i := 0; i < b.N; i++ {
		length := len(items)
		var tmp int
		for k := 0; k < length; k++ {
			tmp = items[k].id
		}
		_ = tmp
	}
}

func BenchmarkRangeIndexStruct(b *testing.B) {
	var items [1024]Item
	for i := 0; i < b.N; i++ {
		var tmp int
		for k := range items {
			tmp = items[k].id
		}
		_ = tmp
	}
}

func BenchmarkRangeStruct(b *testing.B) {
	var items [1024]Item
	for i := 0; i < b.N; i++ {
		var tmp int
		for _, item := range items {
			tmp = item.id
		}
		_ = tmp
	}
}

/*
仅遍历下标的情况下，for 和 range 的性能几乎是一样的。
items 的每一个元素的类型是一个结构体类型 Item，Item 由两个字段构成，一个类型是 int，一个是类型是 [4096]byte，也就是说每个 Item 实例需要申请约 4KB 的内存。
在这个例子中，for 的性能大约是 range (同时遍历下标和值) 的 2000 倍。

*/

func generateItems(n int) []*Item {
	items := make([]*Item, 0, n)
	for i := 0; i < n; i++ {
		items = append(items, &Item{id: i})
	}
	return items
}

func BenchmarkForPointer(b *testing.B) {
	items := generateItems(1024)
	for i := 0; i < b.N; i++ {
		length := len(items)
		var tmp int
		for k := 0; k < length; k++ {
			tmp = items[k].id
		}
		_ = tmp
	}
}

func BenchmarkRangePointer(b *testing.B) {
	items := generateItems(1024)
	for i := 0; i < b.N; i++ {
		var tmp int
		for _, item := range items {
			tmp = item.id
		}
		_ = tmp
	}
}

/*
运行结果如下：
goos: darwin
goarch: amd64
pkg: example/hpg-range
BenchmarkForPointer-8             271279              4160 ns/op
BenchmarkRangePointer-8           264068              4194 ns/op
切片元素从结构体 Item 替换为指针 *Item 后，for 和 range 的性能几乎是一样的。
而且使用指针还有另一个好处，可以直接修改指针对应的结构体的值。
*/

/*
3 总结
range 在迭代过程中返回的是迭代值的拷贝，如果每次迭代的元素的内存占用很低，
那么 for 和 range 的性能几乎是一样，例如 []int。但是如果迭代的元素内存占用较高，
例如一个包含很多属性的 struct 结构体，那么 for 的性能将显著地高于 range，
有时候甚至会有上千倍的性能差异。
对于这种场景，建议使用 for，如果使用 range，建议只迭代下标，通过下标访问迭代值，
这种使用方式和 for 就没有区别了。如果想使用 range 同时迭代下标和值，
则需要将切片/数组的元素改为指针，才能不影响性能。
*/

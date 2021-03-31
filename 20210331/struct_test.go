package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
type Value struct {
    Name   string
    Gender string
}

func main() {
    v1 := Value{Name: "煎鱼", Gender: "男"}
    v2 := Value{Name: "煎鱼", Gender: "男"}
    if v1 == v2 {
        fmt.Println("脑子进煎鱼了")
        return
    }

    fmt.Println("脑子没进煎鱼")
}
*/

/*
type Value struct {
    Name   string
    Gender *string
}

func main() {
    v1 := Value{Name: "煎鱼", Gender: new(string)}
    v2 := Value{Name: "煎鱼", Gender: new(string)}
    if v1 == v2 {
        fmt.Println("脑子进煎鱼了")
        return
    }

    fmt.Println("脑子没进煎鱼")
}
*/

/*
type Value struct {
    Name   string
    GoodAt []string
}

func main() {
    v1 := Value{Name: "煎鱼", GoodAt: []string{"炸", "煎", "蒸"}}
    v2 := Value{Name: "煎鱼", GoodAt: []string{"炸", "煎", "蒸"}}
    if v1 == v2 {
        fmt.Println("脑子进煎鱼了")
        return
    }

    fmt.Println("脑子没进煎鱼")
}
*/

/*
type Value1 struct {
    Name string
}

type Value2 struct {
    Name string
}

func main() {
    v1 := Value1{Name: "煎鱼"}
    v2 := Value2{Name: "煎鱼"}
    if v1 == v2 {
        fmt.Println("脑子进煎鱼了")
        return
    }

    fmt.Println("脑子没进煎鱼")
}
显然，会直接报错：
# command-line-arguments
./main.go:18:8: invalid operation: v1 == v2 (mismatched types Value1 and Value2)


那是不是就完全没法比较了呢？并不，我们可以借助强制转换来实现：
 if v1 == Value1(v2) {
  fmt.Println("脑子进煎鱼了")
  return
 }
*/

/*
在 Go 语言中，Go 结构体有时候并不能直接比较，当其基本类型包含：slice、map、function 时，是不能比较的。若强行比较，就会导致出现例子中的直接报错的情况。
而指针引用，其虽然都是 new(string)，从表象来看是一个东西，但其具体返回的地址是不一样的。
因此若要比较，则需改为：

func main() {
    gender := new(string)
    v1 := Value{Name: "煎鱼", Gender: gender}
    v2 := Value{Name: "煎鱼", Gender: gender}
    ...
}

这样就可以保证两者的比较。如果我们被迫无奈，被要求一定要用结构体比较怎么办？
这时候可以使用反射方法 reflect.DeepEqual，如下：
*/

type Value struct {
	Name   string
	GoodAt []string
}

func TestSt(t *testing.T) {
	v1 := Value{Name: "煎鱼", GoodAt: []string{"炸", "煎", "蒸"}}
	v2 := Value{Name: "煎鱼", GoodAt: []string{"炸", "煎", "蒸"}}
	if reflect.DeepEqual(v1, v2) {
		fmt.Println("脑子进煎鱼了")
		return
	}

	fmt.Println("脑子没进煎鱼")
}

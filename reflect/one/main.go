package main

import (
	"fmt"
	"reflect"
)

func main() {
	n := 123
	p := &n
	vp := reflect.ValueOf(p)
	fmt.Println(vp.CanSet(), vp.CanAddr()) // false false
	vn := vp.Elem()                        // 取得vp的底层指针值引用的值的代表值
	fmt.Println(vn.CanSet(), vn.CanAddr()) // true true
	vn.Set(reflect.ValueOf(789))           // <=> vn.SetInt(789)
	fmt.Println(n)                         // 789
}

/*


type name1 struct {
	name string
}

func main() {
	ty := reflect.TypeOf(name1{})
	fmt.Println(ty.Kind())
	fmt.Println(ty.NumField())
	//ta := reflect.ArrayOf(5, reflect.TypeOf(123))
	//fmt.Println(ta.String()) // [5]int
	//tc := reflect.ChanOf(reflect.SendDir, ta)
	//fmt.Println(tc) // chan<- [5]int
	//tp := reflect.PtrTo(ta)
	//fmt.Println(tp) // *[5]int
	//
	//ts := reflect.SliceOf(tp)
	//fmt.Println(ts) // []*[5]int
	//
	//tm := reflect.MapOf(ta, tc)
	//fmt.Println(tm) // map[[5]int]chan<- [5]int
	//tf := reflect.FuncOf([]reflect.Type{ta},
	//	[]reflect.Type{tp, tc}, false)
	//fmt.Println(tf) // func([5]int) (*[5]int, chan<- [5]int)
	//tt := reflect.StructOf([]reflect.StructField{
	//	{Name: "Age", Type: reflect.TypeOf("abc")},
	//})
	//fmt.Println(tt)            // struct { Age string }
	//fmt.Println(tt.NumField()) // 1
}















type T struct {
	X    int  `max:"99" min:"0" default:"0"`
	Y, Z bool  `optional:"yes"`
}

func main1() {
	t := reflect.TypeOf(T{})
	x := t.Field(0).Tag
	y := t.Field(1).Tag
	z := t.Field(2).Tag
	fmt.Println(reflect.TypeOf(x)) // reflect.StructTag
	// v的类型为string
	v, present := x.Lookup("max")
	fmt.Println(len(v), present)      // 2 true
	fmt.Println(x.Get("max"))         // 99
	fmt.Println(x.Lookup("optional")) //  false
	fmt.Println(y.Lookup("optional")) // yes true
	fmt.Println(z.Lookup("optional")) // yes true
}



type F func(string, int) bool
func (f F) m(s string) bool {
	return f(s, 32)
}
func (f F) M() {}

type I interface{m(s string) bool; M()}

func main() {
	var x struct {
		F F
		i I
	}
	tx := reflect.TypeOf(x)
	fmt.Println(tx.Kind())        // struct
	fmt.Println(tx.NumField())    // 2
	fmt.Println(tx.Field(1).Name) // i
	// 包路径（PkgPath）是非导出字段（或者方法）的内在属性。
	fmt.Println(tx.Field(0).PkgPath) //
	fmt.Println(tx.Field(1).PkgPath) // main

	tf, ti := tx.Field(0).Type, tx.Field(1).Type
	fmt.Println(tf.Kind())               // func
	fmt.Println(tf.IsVariadic())         // false
	fmt.Println(tf.NumIn(), tf.NumOut()) // 2 1
	t0, t1, t2 := tf.In(0), tf.In(1), tf.Out(0)
	// 下一行打印出：string int bool
	fmt.Println(t0.Kind(), t1.Kind(), t2.Kind())

	fmt.Println(tf.NumMethod(), ti.NumMethod()) // 1 2
	fmt.Println(tf.Method(0).Name)              // M
	fmt.Println(ti.Method(1).Name)              // m
	_, ok1 := tf.MethodByName("m")
	_, ok2 := ti.MethodByName("m")
	fmt.Println(ok1, ok2) // false true
}





type T []interface{m()}
func (T) m() {}

func main2() {
	tp := reflect.TypeOf(new(interface{}))
	tt := reflect.TypeOf(T{})
	fmt.Println(tp.Kind(), tt.Kind()) // ptr slice

	// 使用间接的方法得到表示两个接口类型的reflect.Type值。
	ti, tim := tp.Elem(), tt.Elem()
	fmt.Println(ti.Kind(), tim.Kind()) // interface interface

	fmt.Println(tt.Implements(tim))  // true
	fmt.Println(tp.Implements(tim))  // false
	fmt.Println(tim.Implements(tim)) // true

	// 所有的类型都实现了任何空接口类型。
	fmt.Println(tp.Implements(ti))  // true
	fmt.Println(tt.Implements(ti))  // true
	fmt.Println(tim.Implements(ti)) // true
	fmt.Println(ti.Implements(ti))  // true
}

func main1() {
	type A = [16]int16
	var c <-chan map[A][]byte
	tc := reflect.TypeOf(c)

	fmt.Println(tc.Kind())    // chan
	fmt.Println(tc.ChanDir()) // <-chan
	tm := tc.Elem()
	ta, tb := tm.Key(), tm.Elem()
	fmt.Println(tm.Kind(), ta.Kind(), tb.Kind()) // map array slice
	tx, ty := ta.Elem(), tb.Elem()

	// byte是uint8类型的别名。
	fmt.Println(tx.Kind(), ty.Kind()) // int16 uint8
	fmt.Println(tx.Bits(), ty.Bits()) // 16 8
	fmt.Println(tx.ConvertibleTo(ty)) // true
	fmt.Println(tb.ConvertibleTo(ta)) // false
	//var s int16 = 99
	//var ddd uint8 = 9
	//s = uint8(ddd)
	//fmt.Println(s)


	// 切片类型和映射类型都是不可比较类型。
	fmt.Println(tb.Comparable()) // false
	fmt.Println(tm.Comparable()) // false
	fmt.Println(ta.Comparable()) // true
	fmt.Println(tc.Comparable()) // true
}

*/

package main // import "x.y/foo"

//
//type A struct {
//	n int
//}
//
//func (a A) m() {
//	fmt.Println("A", a.n)
//}
//
//type I interface {
//	m()
//}
//
//func Bar(i I) {
//	i.m()
//}

//type F func(int) bool
//func (f F) Validate(n int) bool {
//	return f(n)
//}
//func (f *F) Modify(f2 F) {
//	*f = f2
//}
//
//type B bool
//func (b B) IsTrue() bool {
//	return bool(b)
//}
//func (pb *B) Invert() {
//	*pb = !*pb
//}
//
//type I interface {
//	Load()
//	Save()
//}
//
//func PrintTypeMethods(t reflect.Type) {
//	fmt.Println(t, "has", t.NumMethod(), "methods:")
//	for i := 0; i < t.NumMethod(); i++ {
//		fmt.Print(" method#", i, ": ", t.Method(i).Name, "\n")
//	}
//}
//
//func main() {
//	var s struct {
//		F
//		*B
//		I
//	}
//
//
//	PrintTypeMethods(reflect.TypeOf(s))
//	fmt.Println()
//	PrintTypeMethods(reflect.TypeOf(&s))
//
//	var ff I
//	ff = s
//	ff = &s
//	fmt.Println(ff)
//}
//
//
//type MyInt int
//func (mi MyInt) IsOdd() bool {
//	return mi%2 == 1
//}
//
//type Age MyInt
//
//type X struct {
//	MyInt
//}
//func (x X) Double() MyInt {
//	return x.MyInt + x.MyInt
//}
//
//type Y struct {
//	Age
//}
//
//type Z X

//type I interface {
//	m()
//}
//
//type T struct {
//	I
//}
//
//func main() {
//	var t T
//	var i = &t
//	t.I = i
//	i.m() // 将调用t.m()，然后再次调用i.m()，......
//}

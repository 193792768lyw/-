package main

import (
	"fmt"
	"github.com/chai2010/errors"
	"io/ioutil"
	"testing"
)

func loadConfig() error {
	_, err := ioutil.ReadFile("/path/to/file")
	if err != nil {
		return errors.Wrap(err, "read failed")
	}
	return err
	// ...
}

func setup() error {
	err := loadConfig()
	if err != nil {
		return errors.Wrap(err, "invalid config")
	}

	return err
	// ...
}

func main() {



	//if err := setup(); err != nil {
	//
	//	for i, x := range err.(errors.Error).Caller() {
	//		fmt.Printf("caller:%d: %s\n", i, x.FuncName)
	//	}
	//	//log.Fatal(err)
	//}
	//
	//fmt.Println(errors.New("ooooooo"))

	// ...
}

//func main() {
//	fmt.Println(increaseA()) // 1
//	//fmt.Println(increaseB()) // 1
//}
//
//func increaseA() int {
//	var i int
//	defer func() {
//		//fmt.Println("9999")
//		i++
//		fmt.Println(i,"vvv000000000000000000")
//		//fmt.Printf("%p\n",&i)
//	}()
//	//fmt.Printf("%p\n",&i)
//	//fmt.Println("55555")
//	return i
//}

func increaseB() (r int) {
	defer func() {
		//fmt.Println("9999")
		r++
		//fmt.Printf("%p\n",&r)
	}()
	//fmt.Printf("%p\n",&r)
	//fmt.Println("4444")
	return r
}

type S struct {
}

func (d *S) m(x interface{}) {
}
func (d *S) m1(x interface{}) {
}

func g(x *interface{}) {
}

func Test33(t *testing.T) {

}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
	"sync"
	"testing"
	"unsafe"
)

// InterfaceStructure 定义了一个interface{}的内部结构
type InterfaceStructure struct {
	pt uintptr // 到值类型的指针
	pv uintptr // 到值内容的指针
}

// asInterfaceStructure 将一个interface{}转换为InterfaceStructure
func asInterfaceStructure(i interface{}) InterfaceStructure {
	return *(*InterfaceStructure)(unsafe.Pointer(&i))
}

func TestInterfaceStructure(t *testing.T) {
	var i1, i2 interface{}
	var v1 int = 0x0AAAAAAAAAAAAAAA
	var v2 int = 0x0BBBBBBBBBBBBBBB
	i1 = v1
	i2 = v2
	var a *int
	fmt.Printf("sizeof interface{} = %d\n", unsafe.Sizeof(a))
	fmt.Printf("sizeof interface{} = %d\n", unsafe.Sizeof(i1))
	fmt.Printf("i1 %x %+v\n", i1, asInterfaceStructure(i1))
	fmt.Printf("i2 %x %+v\n", i2, asInterfaceStructure(i2))
	var nilInterface interface{}
	// 所以对于一个interface{}类型的nil变量来说，它的两个指针都是0。
	fmt.Printf("nil interface = %+v\n", asInterfaceStructure(nilInterface))
}

func TestAssignInterfaceNil(t *testing.T) {
	var p *int = nil
	var i interface{} = p
	fmt.Printf("%v %+v is nil %v\n", i, asInterfaceStructure(i), i == nil)
}
func TestAssignInterfaceNil1(t *testing.T) {

	fmt.Println(reflect.TypeOf((*error)(nil)).Elem().Name())

	//var i interface{} = nil
	//fmt.Printf("%v %+v is nil %v\n", i, asInterfaceStructure(i), i == nil)
}

func TestStructureAllMethod(t *testing.T) {
	var wg sync.WaitGroup
	typ := reflect.TypeOf(&wg)
	fmt.Println(reflect.Indirect(reflect.ValueOf(&wg)).Type().Name())

	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)

		argv := make([]string, 0, method.Type.NumIn())
		returns := make([]string, 0, method.Type.NumOut())
		// j 从 1 开始，第 0 个入参是 wg 自己。
		for j := 1; j < method.Type.NumIn(); j++ {
			argv = append(argv, method.Type.In(j).Name())
		}
		for j := 0; j < method.Type.NumOut(); j++ {
			returns = append(returns, method.Type.Out(j).Name())
		}
		log.Printf("func (w *%s) %s(%s) %s",
			typ.Elem().Name(),
			method.Name,
			strings.Join(argv, ","),
			strings.Join(returns, ","))
	}
}
type RetainedRes struct {
	HeadPublic []*string `json:"head_public" form:"head_public"`
	HeadOuter  []*string `json:"head_outer" form:"head_outer"`
	HeadInner  []*string `json:"head_inner" form:"head_inner"`
	BodyInner  interface{} `json:"body_inner" form:"body_inner"`
	BodyOuter  interface{} `json:"body_outer" form:"body_outer"`
	TaskId     string      `json:"task_id" form:"task_id"`
}


func TestMaing(t *testing.T) {
	a := RetainedRes{
		HeadPublic: nil,

		TaskId:     "rrrrr",
	}
	dd , _ := json.Marshal(a)
	fmt.Println(string(dd))
	fmt.Println(len(a.HeadPublic))
}



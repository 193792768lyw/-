package main

import (
	"fmt"
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
	fmt.Printf("sizeof interface{} = %d\n", unsafe.Sizeof(i1))
	fmt.Printf("i1 %x %+v\n", i1, asInterfaceStructure(i1))
	fmt.Printf("i2 %x %+v\n", i2, asInterfaceStructure(i2))
	var nilInterface interface{}
	fmt.Printf("nil interface = %+v\n", asInterfaceStructure(nilInterface))
}


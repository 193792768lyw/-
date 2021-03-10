package main

import (
	"fmt"
	"reflect"
	"testing"
)

type Flags uint

const (
	FlagUp           Flags = 1 << iota // is up
	FlagBroadcast                      // supports broadcast access capability
	FlagLoopback                       // is a loopback interface
	FlagPointToPoint                   // belongs to a point-to-point link
	FlagMulticast                      // supports multicast access capability
)

func IsUp(v Flags) bool     { return v&FlagUp == FlagUp }
func TurnDown(v *Flags)     { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }

func TestOne(t *testing.T) {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10001 true"
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10000 false"
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))   // "10010 false"
	fmt.Printf("%b %t\n", v, IsCast(v)) // "10010 true"
}

const (
	Apple, Banana = iota + 1, iota + 2
	Cherimoya, Durian
	Elderberry, Fig
)

func TestTwo(t *testing.T) {
	fmt.Println(Apple, Banana, Cherimoya, Durian, Elderberry, Fig)
}

/*
reflect.ValueOf() 获取指针对应的反射值。
reflect.Indirect() 获取指针指向的对象的反射值。
(reflect.Type).Name() 返回类名(字符串)。
(reflect.Type).Field(i) 获取第 i 个成员变量。
*/
type Account struct {
	Username string
	Password string
}

func TestReflect(t *testing.T) {
	typ := reflect.Indirect(reflect.ValueOf(Account{})).Type()
	fmt.Println(typ.Name()) // Account

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fmt.Println(field.Name) // Username Password
	}
}

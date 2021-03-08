package go_函数

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"testing"
)

func TestFun(t *testing.T) {
	//不依赖具体的文件对象
	// func CloseFile(f *File) error
	//var CloseFile = (*File).Close
	//
	//// 不依赖具体的文件对象
	//// func ReadFile(f *File, offset int64, data []byte) int
	//var ReadFile = (*File).Read
	//
	//// 文件处理
	//f, _ := OpenFile("foo.dat")
	//ReadFile(f, 0, data)
	//CloseFile(f)
}

// 文件对象
type File struct {
	fd int
}

// 关闭文件
func (f *File) Close() error {
	// ...
	return nil
}

// 读文件数据
func (f *File) Read(offset int64, data []byte) int {
	// ...
	return 6
}

func TestCache(t *testing.T) {
	dd := Cache{
		m: make(map[string]string),
	}
	dd.m["ppp"] = "liuyaowu"
	fmt.Println(dd.Lookup("ppp"))
}

type Cache struct {
	m map[string]string
	sync.Mutex
}

func (p *Cache) Lookup(key string) string {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	return p.m[key]
}

func (p *Cache) Lock() {
	fmt.Println("pppppppppppppppp")
}
func (p *Cache) Unlock() {
	fmt.Println("pppppppppppppppp")

}

type UpperWriter struct {
	io.Writer
}

func (p *UpperWriter) Write(data []byte) (n int, err error) {
	return p.Writer.Write(bytes.ToUpper(data))
}
func TestInter(t *testing.T) {
	fmt.Fprintln(&UpperWriter{os.Stdout}, "hello, world")
}

type UpperString string

//func (s UpperString) String() string {
//	return strings.ToUpper(string(s))
//}

//type fmt.Stringer interface {
//	String() string
//}

func TestPrint(t *testing.T) {
	fmt.Fprintln(os.Stdout, UpperString("hello, world"))
	//var ss int = 99
	//var dd interface{} = ss
	//va,ok := dd.(float64)
	//fmt.Println(va,ok)
}

//var (
//	a io.ReadCloser = (*os.File)(f) // 隐式转换, *os.File 满足 io.ReadCloser 接口
//	b io.Reader     = a             // 隐式转换, io.ReadCloser 满足 io.Reader 接口
//	c io.Closer     = a             // 隐式转换, io.ReadCloser 满足 io.Closer 接口
//	d io.Reader    = c.(io.Reader) // 显式转换, io.Closer 不满足 io.Reader 接口
//)

type TB struct {
	testing.TB
}

func (p *TB) Fatal(args ...interface{}) {
	fmt.Println("TB.Fatal disabled!", args)
}

func TestMk(t *testing.T) {
	var tb testing.TB = new(TB)
	tb.Log("Hello, playground")
}

func TestError(t *testing.T) {

	//err := syscall.Chmod(":invalid path:", 0666)
	//fmt.Println(err.Error())
	//if err != nil {
	//	vv , ok := err.(syscall.Errno)
	//	log.Fatal(vv , ok)
	//}
	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
		}
	}()
	panic("00")
}
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

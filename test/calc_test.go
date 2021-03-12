package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

//
//func TestAdd(t *testing.T) {
//	if ans := Add(1, 2); ans != 3 {
//		t.Errorf("1 + 2 expected be 3, but %d got", ans)
//	}
//
//	if ans := Add(-10, -20); ans != -30 {
//		t.Errorf("-10 + -20 expected be -30, but %d got", ans)
//	}
//}
//
//func TestMul1(t *testing.T) {
//	t.Run("pos", func(t *testing.T) {
//		if Mul(2, 3) != 6 {
//			t.Fatal("fail")
//		}
//
//	})
//	t.Run("neg", func(t *testing.T) {
//		if Mul(2, -3) != -6 {
//			t.Fatal("fail")
//		}
//	})
//}
//
//func TestMul(t *testing.T) {
//	cases := []struct {
//		Name           string
//		A, B, Expected int
//	}{
//		{"pos", 2, 3, 6},
//		{"neg", 2, -3, -6},
//		{"zero", 2, 0, 0},
//	}
//
//	for _, c := range cases {
//		t.Run(c.Name, func(t *testing.T) {
//			fmt.Println(c)
//			if ans := Mul(c.A, c.B); ans != c.Expected {
//				t.Fatalf("%d * %d expected %d, but %d got",
//					c.A, c.B, c.Expected, ans)
//			}
//		})
//	}
//}

//type calcCase struct {
//	Name           string
//	A, B, Expected int
//}
//
//func createMulTestCase(t *testing.T, c *calcCase) {
//	//t.Helper()
//	t.Run(c.Name, func(t *testing.T) {
//		if ans := Mul(c.A, c.B); ans != c.Expected {
//			t.Fatalf("%d * %d expected %d, but %d got",
//				c.A, c.B, c.Expected, ans)
//		}
//	})
//}
//
//func TestMul(t *testing.T) {
//
//	createMulTestCase(t, &calcCase{"one", 2, 3, 6})
//	createMulTestCase(t, &calcCase{"two", 2, -3, 6})
//	createMulTestCase(t, &calcCase{"three", 2, 0, 0}) // wrong case
//}

/*
准备(setup)和回收(teardown)
*/

func setup() {
	fmt.Println("Before all tests")
}

func teardown() {
	fmt.Println("After all tests")
}

func Test1(t *testing.T) {
	fmt.Println("I'm test1")
}

func Test2(t *testing.T) {
	t.Errorf("pppp")
	fmt.Println("I'm test2")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, r.URL)
	w.Write([]byte("hello world"))
}

func handleError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("failed", err)
	}
}

func TestConn(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	handleError(t, err)
	defer ln.Close()

	http.HandleFunc("/hello", helloHandler)
	go http.Serve(ln, nil)

	resp, err := http.Get("http://" + ln.Addr().String() + "/hello")
	handleError(t, err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	handleError(t, err)

	if string(body) != "hello world" {
		t.Fatal("expected hello world, but got", string(body))
	}
}

func TestConn1(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	helloHandler(w, req)
	bytes, _ := ioutil.ReadAll(w.Result().Body)

	if string(bytes) != "hello world" {
		t.Fatal("expected hello world, but got", string(bytes))
	}
}

func BenchmarkHello(b *testing.B) {
	//for i := 0; i < b.N; i++ {
	//	fmt.Sprintf("hello")
	//}
	v := 0.099
	fmt.Println(v * 100)
	fmt.Println(fmt.Sprintf("%.2f", v*100))
}

func TestAddmain(t *testing.T) {
	http.HandleFunc("/", handler)
	//http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func TestAdd(t *testing.T) {
	dd := strings.Split("/a/b", "/")
	fmt.Println(dd)
}

func TestMainn(t *testing.T) {
	si := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("%v  len %d \n", si, len(si))

	test1(si)

	fmt.Printf("%v  len %d \n", si, len(si))
}

func test1(si []int) {
	si = append(si[:3], si[1:]...)
}

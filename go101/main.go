package main

import "fmt"

func main() {

	s := [...]struct{ age int }{
		{44}, {55}, {66},
	}

	for _, v := range s {
		fmt.Println(v)
		s[1].age = 8888
	}
	fmt.Println(s)

}

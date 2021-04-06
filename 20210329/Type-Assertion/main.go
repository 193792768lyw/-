package main

import (
	"encoding/csv"
	"os"
)

//
//func main() {
//	var i interface{} = 10
//	t1 := i.(int)
//	fmt.Println(t1)
//
//	fmt.Println("=====分隔线=====")
//
//	t2 := i.(string)
//	fmt.Println(t2)
//}

func main() {
	f, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF")

	w := csv.NewWriter(f)
	data := [][]string{
		{"1", "test1", "test1-3"},
		{"2", "test2", "test2-1"},
		{"3", "test3", "test3-1"},
		{"3", "test3", "刘耀武"},
	}

	w.WriteAll(data)
}

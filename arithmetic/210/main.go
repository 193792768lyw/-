package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	s := `  {
    "duplicate_83075415": [
      "merge_78dd692d"
    ],
    "duplicate_9c91e7cd": [
      "max_2bfe204f",
      "merge_2af6b446"
    ],
    "max_2bfe204f": [
      "merge_2af6b446"
    ],
    "merge_78dd692d": [
      "insert_d0cced32"
    ],
    "merge_2af6b446": [
      "merge_bd7dbae8"
    ],
    "merge_bd7dbae8": [
      "merge_78dd692d"
    ],
    "presto_58384822": [
      "duplicate_83075415"
    ],
    "presto_06ee0c23": [
      "duplicate_9c91e7cd"
    ],
    "presto_ea7ade44": [
      "merge_bd7dbae8"
    ],
    "insert_d0cced32": [
      "duplicate_8f486100"
    ],
    "duplicate_8f486100": [
      "selectmodel_4fa81cba"
    ],
    "selectmodel_4fa81cba": [
      "IMongo_53280647"
    ],
    "IMongo_53280647": []
  }`
	s1 := map[string]interface{}{}
	json.Unmarshal([]byte(s), &s1)

	nodeMap := []string{}
	nodeMapIndex := map[string]int{}
	count := 0
	for k, _ := range s1 {
		nodeMap = append(nodeMap, k)
		nodeMapIndex[k] = count
		count++
	}
	prerequisites := make([][]int, 0)
	for k, v := range s1 {
		n1 := nodeMapIndex[k]
		if in, ok := v.([]interface{}); ok {
			for _, node := range in {
				index := nodeMapIndex[node.(string)]
				prerequisites = append(prerequisites, []int{index, n1})
			}
		}
	}

	fmt.Println(findOrder(nodeMap, prerequisites))
}

type EdgeNode struct {
	adjvex int
	next   *EdgeNode
}

type VertexNode struct {
	in        int
	data      string
	firstEdge *EdgeNode
}

func findOrder(nodes []string, prerequisites [][]int) []string {
	var (
		edges  = make([]*VertexNode, len(nodes))
		result []string
	)

	for i := 0; i < len(nodes); i++ {
		edges[i] = &VertexNode{
			in:        0,
			data:      nodes[i],
			firstEdge: nil,
		}
	}

	for _, info := range prerequisites {
		edge := &EdgeNode{
			adjvex: info[0],
		}
		edge.next = edges[info[1]].firstEdge
		edges[info[1]].firstEdge = edge
		edges[info[0]].in++
	}

	q := []int{}
	for i := 0; i < len(nodes); i++ {
		if edges[i].in == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		result = append(result, edges[u].data)
		for e := edges[u].firstEdge; e != nil; e = e.next {
			edges[e.adjvex].in--
			if edges[e.adjvex].in == 0 {
				q = append(q, e.adjvex)
			}
		}
	}
	if len(result) != len(nodes) {
		return []string{}
	}
	return result
}

/*
func main() {
	fmt.Println(findOrder(2, [][]int{{1, 0}}))
}

type EdgeNode struct {
	adjvex int
	next   *EdgeNode
}

type VertexNode struct {
	in        int
	data      int
	firstEdge *EdgeNode
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		edges  = make([]VertexNode, numCourses)
		result []int
	)
	//for i := 0 ; i < numCourses ; i ++{
	//	edges[i] = &VertexNode{
	//		in:        0,
	//		data:      i,
	//		firstEdge: nil,
	//	}
	//}

	for _, info := range prerequisites {
		edge := &EdgeNode{
			adjvex: info[0],
		}
		edge.next = edges[info[1]].firstEdge
		edges[info[1]].firstEdge = edge
		edges[info[0]].in++
	}

	q := []int{}
	for i := 0; i < numCourses; i++ {
		if edges[i].in == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		result = append(result, u)
		for e := edges[u].firstEdge; e != nil; e = e.next {
			edges[e.adjvex].in--
			if edges[e.adjvex].in == 0 {
				q = append(q, e.adjvex)
			}
		}
	}
	if len(result) != numCourses {
		return []int{}
	}
	return result
}
*/
func findOrder1(numCourses int, prerequisites [][]int) []int {
	var (
		edges  = make([][]int, numCourses)
		indeg  = make([]int, numCourses)
		result []int
	)

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
		indeg[info[0]]++
	}

	q := []int{}
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		result = append(result, u)
		for _, v := range edges[u] {
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}
	if len(result) != numCourses {
		return []int{}
	}
	return result
}

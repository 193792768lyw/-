package main

import "fmt"

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

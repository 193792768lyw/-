package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Print(pRoot *TreeNode) [][]int {

	// write code here
	res := make([][]int, 0)
	if pRoot == nil {
		return res
	}
	quene := make([]*TreeNode, 0)
	quene = append(quene, pRoot)
	lev := 1
	for len(quene) > 0 {
		temp := make([]int, 0)
		quene1 := make([]*TreeNode, 0)
		for _, r := range quene {
			temp = append(temp, r.Val)
			if r.Left != nil {
				quene1 = append(quene1, r.Left)
			}
			if r.Right != nil {
				quene1 = append(quene1, r.Right)
			}
		}

		if lev%2 == 0 {
			for l, r := 0, len(temp)-1; l < r; {
				temp[l], temp[r] = temp[r], temp[l]
				l++
				r--
			}

		}

		res = append(res, temp)
		lev++
		quene = quene1
	}

	return res
}

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

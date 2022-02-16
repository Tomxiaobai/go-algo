package test

import (
	"fmt"
	leetcodemediumlesson "go-algo/leetcode-medium-lesson"
	"testing"
)

func TestSetZero(t *testing.T) {
	matrix := [][]int{{1, 2, 0}, {3, 4, 5}, {2, 5, 7}}
	fmt.Println("Before test:", matrix)
	leetcodemediumlesson.SetZeros(matrix)
	fmt.Println("After test:", matrix)
}
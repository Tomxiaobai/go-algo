package leetcodemediumlesson

import (
	"sort"
)

func SetZeros(matrix [][]int) {
	row, col := len(matrix), len(matrix[0])
	flagRow0, flagCol0 := false, false
	for _, v := range matrix[0] {
		if v == 0 {
			flagRow0 = true
		}
	}

	for _, v := range matrix {
		if v[0] == 0 {
			flagCol0 = true
		}
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if flagCol0 {
		for _, r := range matrix {
			r[0] = 0
		}
	}

	if flagRow0 {
		for i := 0; i < col; i++ {
			matrix[0][i] = 0
		}
	}
}

func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	for _, v := range strs {
		s := []byte(v)
		sort.Slice(s, func(i, j int) bool {return s[i] < s[j]})
		sortedStr := string(s)  // key
		mp[sortedStr] = append(mp[sortedStr], v)
	}
	ans := make([][]string, 0, len(mp))

	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}

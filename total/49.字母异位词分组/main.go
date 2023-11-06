package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	size := len(strs)
	m := make(map[string][]string, size)

	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sortedStr := string(s)
		m[sortedStr] = append(m[sortedStr], str)
	}

	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}

	return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := groupAnagrams(strs)
	fmt.Println(res)
}

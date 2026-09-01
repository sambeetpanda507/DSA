package main

import (
	"fmt"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	result := [][]string{}
	groups := make(map[string][]string)
	for _, str := range strs {
		hash := [26]int{}
		for _, s := range str {
			index := s - 'a'
			hash[index] += 1
		}

		var c strings.Builder
		for i, val := range hash {
			if val > 0 {
				fmt.Fprintf(&c, "%d%c", val, i+'a')
			}
		}

		key := c.String()
		if len(groups[key]) == 0 {
			groups[key] = []string{str}
		} else {
			groups[key] = append(groups[key], str)
		}
	}

	for _, val := range groups {
		result = append(result, val)
	}

	return result
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagrams(strs)
	fmt.Println(result)
}

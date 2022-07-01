package util

import (
	"strings"
)

func FilterPrefixList(arr *[]string, prefix string) []string {
	s1 := make([]string, 0)

	for _, v := range *arr {
		if strings.Index(v, prefix) == 0 && v != prefix {
			s1 = append(s1, v)
		}
	}

	return s1
}

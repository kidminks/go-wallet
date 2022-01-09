package utils

import "strings"

func SplitString(data string, splitter string) []string {
	return strings.Split(data, splitter)
}

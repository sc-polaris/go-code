package main

import (
	"fmt"
	"strings"
)

func splitWordsBySeparator(words []string, separator byte) []string {
	var res []string
	for _, w := range words {
		for _, s := range strings.Split(w, string(separator)) {
			fmt.Println(s)
			if s != "" {
				res = append(res, s)
			}
		}
	}
	return res
}

func main() {

}

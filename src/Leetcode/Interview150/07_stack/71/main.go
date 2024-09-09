package main

import (
	"strings"
)

func simplifyPath(path string) string {
	var st []string
	for _, v := range strings.Split(path, "/") {
		if v == ".." {
			if len(st) > 0 {
				st = st[:len(st)-1]
			}
		} else if v != "" && v != "." {
			st = append(st, v)
		}
	}
	return "/" + strings.Join(st, "/")
}

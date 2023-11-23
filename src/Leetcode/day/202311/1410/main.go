package main

import "strings"

func entityParser(text string) string {
	entityMap := map[string]string{
		"&quot;":  "\"",
		"&apos;":  "'",
		"&gt;":    ">",
		"&lt;":    "<",
		"&frasl;": "/",
		"&amp;":   "&",
	}

	n := len(text)
	var res strings.Builder
	for i := 0; i < n; {
		isEntity := false
		if text[i] == '&' {
			for k, v := range entityMap {
				if i+len(k) <= n && text[i:i+len(k)] == k {
					res.WriteString(v)
					isEntity = true
					i += len(k)
					break
				}
			}
		}
		if !isEntity {
			res.WriteByte(text[i])
			i++
		}
	}

	return res.String()
}

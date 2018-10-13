package core

import "fmt"

type line struct {
	text string
	no int
}

type token struct {
	text string
}

func tokenize(l *line) string {
	fmt.Println(l.text, l.no)
	return ""
}

package main

import (
	"fmt"
	"os"
	"bufio"
	"core"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fileName := os.Args[1]
	fmt.Println(fileName)

	file, err := os.Open(fileName)
	check(err)

	scanner := bufio.NewScanner(file)
	counter := 0
	for scanner.Scan() {
		counter++
		text := scanner.Text()
		core.tokenize({text: text, no: counter})
		fmt.Println(text)
	}
}
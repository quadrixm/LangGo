package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fileName := os.Args[1]
	fmt.Println(fileName)

	dat, err := ioutil.ReadFile(fileName)
	check(err)
	fmt.Print(string(dat))
}
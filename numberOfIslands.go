package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	files, err := ioutil.ReadDir("./matrices/")
	check(len(files), err)

	for _, f := range files {
		fmt.Println(f.Name())
	}
}

func check(lenght int, e error) {
	if e != nil {
		panic(e)
	} else if lenght == 0 {
		panic("oh no! folder is empty")
	}
}

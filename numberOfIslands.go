package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    files, err := ioutil.ReadDir("./matrices/")
	check(err)

	for _, f := range files {
		fmt.Println(f.Name())
	}
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
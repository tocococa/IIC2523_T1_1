package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	files, err := ioutil.ReadDir("./matrices/")
	check(err)
	checkDirIsEmpty(len(files))

	for _, f := range files {
		dat := openFile(f.Name())
		parseFile(dat) // retornar estructura
		// resolver problema
		f := createFile("numberOfIslands_<id>.txt")
		writeFileAndClose("2", f)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkDirIsEmpty(lenght int) {
	if lenght == 0 {
		panic("oh no! folder is empty")
	}
}

func openFile(fileName string) string {
	dat, err := os.ReadFile("./matrices/" + fileName)
	check(err)
	return string(dat)
}

func createFile(name string) *os.File {
	f, err := os.Create("./respuestas/" + name)
	check(err)
	//defer f.Close()
	return f
}

func writeFileAndClose(data string, f *os.File) {
	d := []byte(data)
	_, err := f.Write(d)
	check(err)
	defer f.Close()
}

func parseFile(dat string) {
	fmt.Print(dat + "\n")
}

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	files, err := ioutil.ReadDir("./matrices/")
	check(err)
	checkDirIsEmpty(len(files))

	for _, f := range files {
		dat := openFile(f.Name())
		// todo: sacar id y cols del nombre del archivo
		matrix := parseFile(dat, 5) // retornar matriz
		fmt.Println(matrix)

		// todo: resolver problema
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
	return f
}

func writeFileAndClose(data string, f *os.File) {
	d := []byte(data)
	_, err := f.Write(d)
	check(err)
	defer f.Close()
}

//func getIdAndCols(fileName string) []int {

//}

func parseFile(dat string, cols int) [][]int {
	list := strings.Split(dat, ",")
	rows := len(list) / cols
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			intVar, _ := strconv.Atoi(list[i*cols+j])
			matrix[i][j] = intVar
		}
	}
	return matrix
}

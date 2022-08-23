package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	answersDirectory := "./respuestas/"
	correctAnswers := map[string]string{
		"1": "0",
		"2": "6",
		"3": "1",
		"4": "19",
		"5": "1",
		"6": "641",
	}
	grade := 1

	f, err := os.Open(answersDirectory)
	if err != nil {
		fmt.Println(err)
		return
	}
	files, err := f.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range files {
		if !v.IsDir() {
			fileName := v.Name()
			fmt.Println("fileName: ", fileName)
			matrixInfo := strings.Split(fileName, "_")
			matrixIndex := strings.TrimSuffix(matrixInfo[1], ".txt")
			fmt.Println("matrixIndex: ", matrixIndex)

			numberOfIslands, err := os.ReadFile(answersDirectory + fileName)
			if err != nil {
				panic(err)
			}

			if string(numberOfIslands) == correctAnswers[matrixIndex] {
				//fmt.Println("Bien")
				grade++
			} else {
				//fmt.Println("Mal")
			}
		}
	}
	fmt.Println("----------------------------------")
	fmt.Println("Nota Código:", grade)
	theory := 7
	fmt.Println("Nota Explicaciones: ", theory)
	eval := 1
	fmt.Println("Ponderador Evaluación de Pares y Código de Ética: ", eval)
	finalGrade := (float64(grade)*0.7 + float64(theory)*0.3) * float64(eval)
	fmt.Println("Nota Final: ", finalGrade)
}

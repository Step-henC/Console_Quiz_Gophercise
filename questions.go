package main

import (
	"fmt"
	"os"
	"strings"
)

type qAnda struct {
	question string
	answer   string
}

func parseFile(filename string) []string {

	os.Open(filename)

	quizBytes, err := os.ReadFile(filename)

	if err != nil {

		fmt.Printf("ERROR: %v", err)

		os.Exit(1)

	}

	res := string(quizBytes)

	quizArr := strings.Split(res, ",")

	return quizArr

}

func createQuiz(csvLines []string) []qAnda {

	if len(csvLines) <= 0 {

		return []qAnda{}
	}

	var quiz []qAnda

	for i := 0; i < len(csvLines); i += 2 {

		var elem qAnda

		elem.question = csvLines[i]
		elem.answer = csvLines[i+1]

		quiz = append(quiz, elem)

	}

	return quiz
}

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

	return strings.Split(res, ",")

}

func createQuiz(csvLines []string) []qAnda {

	if len(csvLines) <= 0 {

		return []qAnda{}
	}

	var quiz []qAnda

	for i, line := range csvLines {

		var elem qAnda

		if i%2 == 0 {

			elem.question = line

		} else {
			elem.answer = line
		}

		quiz = append(quiz, elem)

	}

	return quiz
}

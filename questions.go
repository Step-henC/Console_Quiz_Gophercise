package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type qAnda struct {
	question string
	answer   string
}

type quiz struct {
	quesAndAnswerList []qAnda
	correctAnswers    int
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

func createQuiz(csvLines []string) quiz {

	if len(csvLines) <= 0 {

		return quiz{}
	}

	var res []qAnda

	var test quiz

	for i := 0; i < len(csvLines); i += 2 {

		var elem qAnda

		elem.question = csvLines[i]
		elem.answer = csvLines[i+1]

		res = append(res, elem)

	}

	test.quesAndAnswerList = res

	return test
}

func (q quiz) performConsoleQuiz() {

	correctCount := 0

	for _, quesResponse := range q.quesAndAnswerList {

		fmt.Println(quesResponse.question + "?")

		sc := bufio.NewScanner(os.Stdin)

		for sc.Scan() {

			if sc.Text() == quesResponse.answer {

				fmt.Println("Correct!")
				correctCount++

			} else {

				fmt.Println("Sorry, the correct answer is: " + quesResponse.answer)

			}

		}

	}

	q.correctAnswers = correctCount

}

func (q quiz) gradeQuiz() int {

	passing := 70

	grade := (q.correctAnswers / len(q.quesAndAnswerList) * 100)

	if grade > passing {

		fmt.Printf("You got %v out of %v, keep up the good work! \n", q.correctAnswers, len(q.quesAndAnswerList))

	} else {

		fmt.Printf("You got %v out of %v, better luck next time. \n", q.correctAnswers, len(q.quesAndAnswerList))

	}

	return grade
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Hi! Ready to get started on your Quiz? (y/n)")

	for scanner.Scan() {

		if scanner.Text() == "y" {

			res := parseFile("problems.csv")
			quiz := createQuiz(res)

			fmt.Println(quiz[0])
		}
	}

}

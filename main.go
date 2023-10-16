package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	correctCount := 0

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Hi! Ready to get started on your Quiz? (y/n)")

	for scanner.Scan() {

		if scanner.Text() == "y" {

			res := parseFile("problems.csv")
			quiz := createQuiz(res)

			for _, quesResponse := range quiz {

				fmt.Println(quesResponse.question + "?")

				sc := bufio.NewScanner(os.Stdin)

				for sc.Scan() {

					if sc.Text() == quesResponse.answer {

						fmt.Println("Correct!")
						correctCount++

					} else {

						fmt.Println("Sorry, the correct answer is: " + quesResponse.answer)
					}

					break

				}

			}

			passing := 70

			if (correctCount / len(quiz) * 100) > passing {

				fmt.Printf("You got %v out of %v, keep up the good work! \n", correctCount, len(quiz))
				break

			} else {

				fmt.Printf("You got %v out of %v, better luck next time. \n", correctCount, len(quiz))
				break

			}

		} else {

			fmt.Println("Sorry to see you go. Have a Nice day!")
			break

		}
	}

}

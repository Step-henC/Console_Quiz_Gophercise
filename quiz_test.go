package main

import (
	"testing"
)

func TestCreateEmptyQuiz(t *testing.T) {

	arrStr := parseFile("_testProbs.csv")

	if len(arrStr) != 0 {
		t.Errorf("Expected empty file but got file with content length: %v ", len(arrStr))
	}

}

func TestCreateNewQuiz(t *testing.T) {

	res := parseFile("problems.csv")

	question := res[0]
	answer := res[1]

	if len(res) != 26 {
		t.Errorf("Expected quiz bank with 13 questions and answers [26 total], but got: [ %v ] Q and A's", len(res))
	}

	if question != "5+5" && answer != "10" {
		t.Errorf("Expected 5+5=10 but got %v %v", question, answer)

	}

}

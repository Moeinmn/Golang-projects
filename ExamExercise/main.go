package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/fatih/color"
	"io"
	"log"
	"os"
	"strings"
)

type UserAnswer struct {
	question      string
	answer        string
	correctAnswer string
}

type ExamResult struct {
	results []UserAnswer
}

func (res ExamResult) printResults() {
	for index, val := range res.results {
		if val.answer == val.correctAnswer {
			color.Cyan("Question %v: %v \n Answer : %v \n CurrectAnswer : %v \n", index, val.question, val.answer, val.correctAnswer)
		} else {
			color.Red("Question %v: %v \n Answer : %v \n CurrectAnswer : %v \n", index, val.question, val.answer, val.correctAnswer)
		}
	}
}

// InputPrompt receives a string value using the label
func InputPrompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		n, _ := fmt.Fprint(os.Stderr, label+" ")

		// Printing the number of bytes written
		fmt.Println(n)

		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

func main() {
	// open file
	f, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("Error with cleaning file")
		}
	}(f)

	examResults := ExamResult{}

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		question := rec[0]
		answer := rec[1]

		userAnswer := InputPrompt(question)

		result := UserAnswer{
			question:      question,
			answer:        userAnswer,
			correctAnswer: answer,
		}

		examResults.results = append(examResults.results, result)

	}
	examResults.printResults()
}

package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	q string
	a string
}

func main() {

	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)

	if err != nil {
		exitProgram(fmt.Sprintf("Failed to open csv file: %s", *csvFilename))
	}

	defer file.Close()

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()
	if err != nil {
		exitProgram("Failed to parse csv file")
	}

	problems := parseLines(lines)
	correct := 0

	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s\n", i+1, problem.q)
		var answer string
		n, err := fmt.Scanf("%s\n", &answer)
		_ = n
		if err != nil {
			exitProgram("Failed to read answer")
		}

		if answer == problem.a {
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d questions.\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	s := make([]problem, len(lines))

	for i, line := range lines {
		s[i] = problem{
			line[0],
			strings.TrimSpace(line[1]),
		}
	}

	return s
}

func exitProgram(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

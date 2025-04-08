package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	results := make([]string, 0)
	questions := make([]string, 0)
	file, err := os.Open("problems.csv")
	check(err)
	defer file.Close()

	cReader := bufio.NewReader(os.Stdin)
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	check(err)

	for _, row := range records {
		for j, cell := range row {
			if j == 0 {
				questions = append(questions, cell)
			} else {
				// estos resultados los tengo que guardar
				results = append(results, cell)
			}
		}
	}

	total := len(results)
	correct := 0
	for i, v := range results {

		fmt.Printf("%s =\n", questions[i])
		text, err := cReader.ReadString('\n')
		check(err)
		text = strings.TrimSpace(text)
		if text == v {
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, total)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

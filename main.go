package main

import (
	"fmt"
	"strings"
)

type PhraseCase struct {
	Phrase  string
	Divisor int
}

func main() {
	count := 30
	cases := []PhraseCase{
		{Phrase: "fizz", Divisor: 3},
		{Phrase: "buzz", Divisor: 5},
	}

	for i := 1; i <= count; i++ {
		phrases := EvalPhraseCases(i, cases)
		fmt.Println(strings.Join(phrases, " "))
	}
}

func EvalPhraseCases(num int, cases []PhraseCase) []string {
	var phrases []string

	for _, c := range cases {
		if num%c.Divisor == 0 {
			phrases = append(phrases, c.Phrase)
		}
	}

	if len(phrases) == 0 {
		phrases = append(phrases, fmt.Sprintf("%d", num))
	}

	return phrases
}

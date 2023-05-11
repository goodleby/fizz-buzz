package main

import (
	"reflect"
	"testing"
)

func TestEvalCases(t *testing.T) {
	type args struct {
		num   int
		cases []PhraseCase
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Fizz",
			args: args{
				num: 3,
				cases: []PhraseCase{
					{Phrase: "fizz", Divisor: 3},
					{Phrase: "buzz", Divisor: 5},
				},
			},
			want: []string{"fizz"},
		},
		{
			name: "Buzz",
			args: args{
				num: 5,
				cases: []PhraseCase{
					{Phrase: "fizz", Divisor: 3},
					{Phrase: "buzz", Divisor: 5},
				},
			},
			want: []string{"buzz"},
		},
		{
			name: "FizzBuzz",
			args: args{
				num: 15,
				cases: []PhraseCase{
					{Phrase: "fizz", Divisor: 3},
					{Phrase: "buzz", Divisor: 5},
				},
			},
			want: []string{"fizz", "buzz"},
		},
		{
			name: "None",
			args: args{
				num: 7,
				cases: []PhraseCase{
					{Phrase: "fizz", Divisor: 3},
					{Phrase: "buzz", Divisor: 5},
				},
			},
			want: []string{"7"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EvalPhraseCases(tt.args.num, tt.args.cases); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EvalPhraseCases() = %v, want %v", got, tt.want)
			}
		})
	}
}

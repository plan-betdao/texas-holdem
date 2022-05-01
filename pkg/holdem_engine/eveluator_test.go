package holdem_engine

import (
	"testing"
)

func TestEvaluateCards(t *testing.T) {

	inputs := [10][7][2]byte{
		{{'s', 'a'}, {'s', 'k'}, {'s', 'q'}, {'s', 't'}, {'s', 'j'}, {'s', '9'}, {'s', '8'}},
		{{'s', 'k'}, {'s', '7'}, {'s', 'j'}, {'s', '9'}, {'s', 't'}, {'s', '2'}, {'s', 'q'}},
		{{'s', 't'}, {'d', '3'}, {'d', 't'}, {'h', 't'}, {'c', 't'}, {'s', '5'}, {'s', 'q'}},
		{{'d', '3'}, {'s', '3'}, {'h', '3'}, {'d', '2'}, {'s', '2'}, {'s', 'a'}, {'d', '7'}},
		{{'d', 'a'}, {'s', 'k'}, {'s', '7'}, {'s', '5'}, {'s', '4'}, {'s', '3'}, {'d', '5'}},
		{{'d', 'k'}, {'d', 'q'}, {'d', 't'}, {'s', 'j'}, {'h', '5'}, {'h', '2'}, {'d', '9'}},
		{{'c', '3'}, {'c', '2'}, {'d', 'k'}, {'s', 'k'}, {'c', 'k'}, {'s', 'a'}, {'s', '6'}},
		{{'c', '3'}, {'c', '2'}, {'d', 'k'}, {'s', 'k'}, {'c', 'a'}, {'s', 'a'}, {'s', '6'}},
		{{'c', '3'}, {'c', '2'}, {'d', 'q'}, {'s', 'k'}, {'c', 'a'}, {'s', 'a'}, {'s', '6'}},
		{{'c', '3'}, {'c', '2'}, {'d', 'q'}, {'s', 'k'}, {'c', 'a'}, {'s', 't'}, {'s', '6'}},
	}

	expect_results := [10]EvaluateResult{
		{
			category: "royal_flush",
			value:    [6]int{9, 14, 13, 12, 11, 10},
			picks:    [5][2]byte{{'s', 'a'}, {'s', 'k'}, {'s', 'q'}, {'s', 'j'}, {'s', 't'}}},
		{
			category: "straight_flush",
			value:    [6]int{8, 13, 12, 11, 10, 9},
			picks:    [5][2]byte{{'s', 'k'}, {'s', 'q'}, {'s', 'j'}, {'s', 't'}, {'s', '9'}}},
		{
			category: "four_of_a_kind",
			value:    [6]int{7, 10, 10, 10, 10, 5},
			picks:    [5][2]byte{{'s', 't'}, {'d', 't'}, {'h', 't'}, {'c', 't'}, {'s', '5'}}},
		{
			category: "full_house",
			value:    [6]int{6, 3, 3, 3, 2, 2},
			picks:    [5][2]byte{{'d', '3'}, {'s', '3'}, {'h', '3'}, {'d', '2'}, {'s', '2'}}},
		{
			category: "flush",
			value:    [6]int{5, 13, 7, 5, 4, 3},
			picks:    [5][2]byte{{'s', 'k'}, {'s', '7'}, {'s', '5'}, {'s', '4'}, {'s', '3'}}},
		{
			category: "straight",
			value:    [6]int{4, 13, 12, 11, 10, 9},
			picks:    [5][2]byte{{'d', 'k'}, {'d', 'q'}, {'s', 'j'}, {'d', 't'}, {'d', '9'}}},
		{
			category: "three_of_a_kind",
			value:    [6]int{3, 13, 13, 13, 14, 6},
			picks:    [5][2]byte{{'d', 'k'}, {'s', 'k'}, {'c', 'k'}, {'s', 'a'}, {'s', '6'}}},
		{
			category: "two_pairs",
			value:    [6]int{2, 14, 14, 13, 13, 6},
			picks:    [5][2]byte{{'c', 'a'}, {'s', 'a'}, {'d', 'k'}, {'s', 'k'}, {'s', '6'}}},
		{
			category: "pair",
			value:    [6]int{1, 14, 14, 13, 12, 6},
			picks:    [5][2]byte{{'c', 'a'}, {'s', 'a'}, {'s', 'k'}, {'d', 'q'}, {'s', '6'}}},
		{
			category: "highcard",
			value:    [6]int{0, 14, 13, 12, 10, 6},
			picks:    [5][2]byte{{'c', 'a'}, {'s', 'k'}, {'d', 'q'}, {'s', 't'}, {'s', '6'}}},
	}

	for index, _ := range inputs {
		t.Run(expect_results[index].category, func(t *testing.T) {
			if EvaluateCards(inputs[index]) != expect_results[index] {
				t.Fatal("fail")
			}
		})
	}

}

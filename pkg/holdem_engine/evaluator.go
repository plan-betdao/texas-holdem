package holdem_engine

import (
	"bytes"
	"texas-holdem/pkg/math"
)

var kind_order = map[byte]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	't': 10,
	'j': 11,
	'q': 12,
	'k': 13,
	'a': 14}

var category_order = map[string]int{
	"royal-flush":     9,
	"straight-flush":  8,
	"four-of-a-kind":  7,
	"full-house":      6,
	"flush":           5,
	"straight":        4,
	"three-of-a-kind": 3,
	"two-pairs":       2,
	"pair":            1,
	"highcard":        0}

type EvaluateResult struct {
	category string
	value    [6]int
	picks    [5][2]byte
}

func SortByKinds(cards [7][2]byte) {

}

func GetSortedCards(cards [7][2]byte) {

}

func GetFlushSuitCards(cards [7][2]byte) {

}

var royal_flush = [4][5][2]byte{
	{{'s', 'a'}, {'s', 'k'}, {'s', 'q'}, {'s', 'j'}, {'s', 't'}},
	{{'h', 'a'}, {'h', 'k'}, {'h', 'q'}, {'h', 'j'}, {'h', 't'}},
	{{'d', 'a'}, {'d', 'k'}, {'d', 'q'}, {'d', 'j'}, {'d', 't'}},
	{{'c', 'a'}, {'c', 'k'}, {'c', 'q'}, {'c', 'j'}, {'c', 't'}},
}

func GetValues(category string, picks [5][2]byte) (values [6]int) {
	values[0] = category_order[category]
	for index, v := range picks {
		values[index+1] = kind_order[v[1]]
	}
	return
}

func EvaluateCards(cards [7][2]byte) (result EvaluateResult) {
	for _, v := range royal_flush {
		set := math.Intersect(cards, v, func(i1 interface{}, i2 interface{}) bool {
			i1s2, ok := i1.([2]byte)
			if !ok {
				return false
			}

			i2s2, ok := i2.([2]byte)
			if !ok {
				return false
			}

			if bytes.Equal(i1s2[:], i2s2[:]) {
				return true
			}

			return false
		})
		if len(set) == 5 {
			result.category = "royal-flush"
			// a := (*[5]interface{})(set)
			// result.picks =
			result.value = GetValues(result.category, result.picks)
		}
	}
	return
}

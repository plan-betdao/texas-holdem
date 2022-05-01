package holdem_engine

import (
	"sort"
	"texas-holdem/pkg/math"
)

const (
	suit = 0
	kind = 1
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
	"royal_flush":     9,
	"straight_flush":  8,
	"four_of_a_kind":  7,
	"full_house":      6,
	"flush":           5,
	"straight":        4,
	"three_of_a_kind": 3,
	"two_pairs":       2,
	"pair":            1,
	"highcard":        0}

type EvaluateResult struct {
	category string
	value    [6]int
	picks    [5][2]byte
}

func SortByKinds(cards [7][2]byte) {

}

func GetSortedCards(cards [7][2]byte) [7][2]byte {
	groups := math.GroupBy(cards[:7], func(t1 [2]byte, t2 [2]byte) bool {
		return kind_order[t1[kind]] < kind_order[t2[kind]]
	}, func(t1 [2]byte) int {
		return kind_order[t1[kind]]
	})
	var keys []int
	for k := range groups {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	math.Sort(keys, func(t1 int, t2 int) bool {
		return len(groups[t1]) < len(groups[t2])
	})

	math.Reverse(keys)

	sorded_cards := [][2]byte{}
	for _, v := range keys {
		sorded_cards = append(sorded_cards, groups[v]...)
	}

	return *(*[7][2]byte)(sorded_cards)
}

func GetFlushSuitCards(cards [7][2]byte) [][2]byte {
	groups := math.GroupBy(cards[:7], func(t1 [2]byte, t2 [2]byte) bool {
		return t1[suit] < t2[suit]
	}, func(t1 [2]byte) byte {
		return t1[suit]
	})

	for _, g := range groups {
		if len(g) >= 5 {
			math.Sort(g, func(t1 [2]byte, t2 [2]byte) bool {
				// 这里用> 替换< 为了反转
				return kind_order[t1[kind]] > kind_order[t2[kind]]
			})
			return g
		}
	}
	return nil
}

var g_possiable_straights [10][5]int

func init() {
	for i := 14; i > 5; i-- {
		for j := i; j > i-5; j-- {
			g_possiable_straights[14-i][i-j] = j
		}
	}
	g_possiable_straights[9] = [5]int{5, 4, 3, 2, 14}
}

func ListStraights(cards [7][2]byte) [][2]byte {
	order_cards := math.GroupBy(cards[:7], func(t1 [2]byte, t2 [2]byte) bool {
		return kind_order[t1[kind]] < kind_order[t2[kind]]
	}, func(t1 [2]byte) int {
		return kind_order[t1[kind]]
	})

	for _, s := range g_possiable_straights {
		has := [][2]byte{}
		for _, v := range s {
			if _, ok := order_cards[v]; !ok {
				break
			}
			has = append(has, order_cards[v]...)
		}
		if len(has) == 5 {
			// 这里可能有问题
			// return *(*[5][2]byte)(has)
			return has
		}
	}
	return nil
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

func ArrayElementEqual[T comparable](array []T) bool {
	only_one := map[T]T{}
	for _, v := range array {
		only_one[v] = v
	}
	if len(only_one) == 1 {
		return true
	}
	return false
}

func EvaluateCards(cards [7][2]byte) (result EvaluateResult) {
	sorted_cards := GetSortedCards(cards)
	sorted_kinds := func(input [7][2]byte) (outs [7]byte) {
		for i, v := range input {
			outs[i] = v[kind]
		}
		return
	}(sorted_cards)
	flush_suit_cards := GetFlushSuitCards(cards)
	straights := ListStraights(cards)

	for _, v := range royal_flush {
		set := math.Intersect(v[:5], cards[:7])
		if len(set) == 5 {
			result.category = "royal_flush"
			result.picks = *(*[5][2]byte)(set)
			result.value = GetValues(result.category, result.picks)
			return
		}
	}

	if straights != nil {
		set := math.Intersect(straights[:5], flush_suit_cards)
		if len(set) == 5 {
			result.category = "straight_flush"
			result.picks = *(*[5][2]byte)(set)
			result.value = GetValues(result.category, result.picks)
			return
		}
	}

	four := sorted_kinds[:4]
	if ArrayElementEqual(four) {
		result.category = "four_of_a_kind"
		result.picks = *(*[5][2]byte)(sorted_cards[:5])
		result.value = GetValues(result.category, result.picks)
		return
	}

	three := sorted_kinds[:3]
	two := sorted_kinds[3:5]
	if ArrayElementEqual(three) && ArrayElementEqual(two) {
		result.category = "full_house"
		result.picks = *(*[5][2]byte)(sorted_cards[:5])
		result.value = GetValues(result.category, result.picks)
		return
	}

	if flush_suit_cards != nil {
		result.category = "flush"
		result.picks = *(*[5][2]byte)(flush_suit_cards[:5])
		result.value = GetValues(result.category, result.picks)
		return
	}

	if straights != nil {
		result.category = "straight"
		result.picks = *(*[5][2]byte)(straights[:5])
		result.value = GetValues(result.category, result.picks)
		return
	}

	if ArrayElementEqual(sorted_kinds[:3]) {
		result.category = "three_of_a_kind"
		result.picks = *(*[5][2]byte)(sorted_cards[:5])
		result.value = GetValues(result.category, result.picks)
		return
	}

	if ArrayElementEqual(sorted_kinds[:2]) && ArrayElementEqual(sorted_kinds[2:4]) {
		result.category = "two_pairs"
		result.picks = *(*[5][2]byte)(sorted_cards[:5])
		result.value = GetValues(result.category, result.picks)
		return
	}

	if ArrayElementEqual(sorted_kinds[:2]) {
		result.category = "pair"
		result.picks = *(*[5][2]byte)(sorted_cards[:5])
		result.value = GetValues(result.category, result.picks)
		return
	}

	result.category = "highcard"
	result.picks = *(*[5][2]byte)(sorted_cards[:5])
	result.value = GetValues(result.category, result.picks)

	return
}

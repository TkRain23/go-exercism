package scrabble

import "strings"

var lut = []int{1, 3, 3, 2, 1, 4, 2, 4, 1, 8, 5, 1, 3, 1, 1, 3, 10, 1, 1, 1, 1, 4, 4, 8, 4, 10}

func Score(word string) (ret int) {
	word = strings.ToLower(word)
	for _, ch := range word {
		ret += lut[ch-'a']
	}
	return
}

package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func Dupes(s string, in []string) bool {
	for _, v := range in {
		if v == s {
			return true
		}
	}
	return false
}

func Anagrams(in *[]string) *map[string][]string {
	anagramIdx := make(map[string][]int)

	for i, w := range *in {
		word := []rune(w)
		sort.Slice(word, func(i, j int) bool {
			return word[i] < word[j]
		})
		tmpW := string(word)
		anagramIdx[tmpW] = append(anagramIdx[tmpW], i)
	}

	res := make(map[string][]string)
	var tmp []string

	for _, v := range anagramIdx {
		if len(v) < 2 {
			continue
		}
		tmp = []string{}
		for _, j := range v {
			if Dupes((*in)[j], tmp) == false {
				tmp = append(tmp, (*in)[j])
			}
		}

		zero := tmp[0]
		sort.Strings(tmp)
		res[zero] = tmp
	}

	return &res
}

func randSlice() []string {
	var res []string

	for i := 0; i < 5; i++ {
		x := 5 + rand.Intn(1)
		var word []rune

		for j := 0; j < x; j++ {
			word = append(word, rune('а'+rand.Intn(2)))
		}

		res = append(res, string(word))
		word = []rune{}
	}
	fmt.Println("Start slice:", res)
	return res
}

func main() {
	test := randSlice()
	fmt.Println("Result: ", Anagrams(&test))
}

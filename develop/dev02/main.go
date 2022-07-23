package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func Unpack(str string) string { // перевожу все в руны на случай юникода
	var last rune
	var nums []rune
	out := ""

	runeStr := []rune(str)

	if unicode.IsNumber(runeStr[0]) { // проверяю ошибку первого символа
		return "некорректный ввод"
	}

	last = runeStr[0]

	escape := last == 92 // присваеваем значение бэкслеша

	for _, i := range runeStr[1:] {
		if escape {
			last = i
			escape = false
			continue
		}

		if i == 92 {
			escape = true
			out = out + Result(last, nums)
			continue
		}

		if unicode.IsNumber(i) {
			nums = append(nums, i)
			continue
		}

		out = out + Result(last, nums)
		nums = []rune{}
		last = i
	}

	out = out + Result(last, nums)

	return out
}

func Result(r rune, i []rune) string { // собираем символы
	res := strings.Repeat(string(r), countRunes(i))
	return string(res)
}

func countRunes(rep []rune) int { // считаем количество повторов
	if len(rep) == 0 { // если цифр не дано, значит одно повторение
		return 1
	}

	n, err := strconv.Atoi(string(rep))

	if err != nil {
		log.Fatalf("Have error is - %s", err)
	}

	return n
}

func main() {
	str := "aa6a"

	fmt.Println(Unpack(str))
}

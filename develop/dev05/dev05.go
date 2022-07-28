package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func fileString(file string) []string {
	var res []string
	f, err := os.Open(file)
	defer f.Close()

	if err != nil {
		log.Fatal("You have error is opening files - ", err)
	}

	fileScan := bufio.NewScanner(f)

	for fileScan.Scan() {
		res = append(res, fileScan.Text())
	}

	return res
}

type flags struct {
	A        int
	B        int
	C        int
	c        bool
	i        bool
	v        bool
	F        bool
	n        bool
	e        bool
	regular  string
	fileName string
}

var Flags flags
var Input []string

func extractFlags() {
	flag.IntVar(&Flags.A, "A", 0, "считать строк после совпадения")
	flag.IntVar(&Flags.B, "B", 0, "считать строк перед совпадением")
	flag.IntVar(&Flags.C, "C", 0, "считать строк до и столько же после совпадения")
	flag.BoolVar(&Flags.c, "c", false, "количество результатов")
	flag.BoolVar(&Flags.i, "i", false, "игнорировать регистр")
	flag.BoolVar(&Flags.v, "v", false, "исключать находки")
	flag.BoolVar(&Flags.F, "F", false, "точное совпадение со строкой, не паттерн")
	flag.BoolVar(&Flags.F, "e", false, "заэскейпить паттерн")
	flag.BoolVar(&Flags.n, "n", false, "печать номеров строки")
	flag.Parse()

	if flag.NArg() < 2 {
		log.Fatal("Отсутствует имя файла либо паттерн для поиска!")
	}

	Flags.regular = flag.Args()[0]
	Flags.fileName = flag.Args()[1]
	fmt.Println(Flags)
}

func linesContains(lines []int, target int) bool {
	for _, i := range lines {
		if i == target {
			return true
		}
	}

	return false
}

func invertLines(lines []int) []int {
	var res []int

	if len(lines) == 0 {
		for i := 0; i < len(Input); i++ {
			res = append(res, i)
		}

		return res
	}

	for i := 0; i < len(Input); i++ {
		if !linesContains(lines, i) {
			res = append(res, i)
		}
	}

	return res
}

func getLinesBAC(lines []int) []int {
	var res []int

	for _, i := range lines {
		before := i - Flags.B
		if Flags.C > 0 {
			before = i - Flags.C
		}

		after := i + Flags.A
		if after > len(Input)-1 {
			after = len(Input) - 1
		}

		for j := before; j < after; j++ {
			if !linesContains(res, j) {
				res = append(res, j)
			}
		}
	}

	return res
}

func printLines(lines []int) {
	for _, i := range lines {
		var line string

		if Flags.n {
			line = strconv.Itoa(i+1) + ":\t"
		}

		line = line + Input[i]
		fmt.Println(line)
	}
}

func doRegulat(pattern string) ([]string, []int) {
	var res []string
	var idx []int

	reg, err := regexp.Compile(pattern)

	if err != nil {
		log.Fatal("Неверное выражение. Попробуйте добавить: -е")
	}

	for i, j := range Input {
		if Flags.i {
			j = strings.ToLower(j)
		}

		if reg.MatchString(j) {
			res = append(res, Input[i])
			idx = append(idx, i)
		}
	}

	return res, idx
}

func doFixed(pattern string) ([]string, []int) {
	var res []string
	var idx []int

	for i, j := range Input {
		if Flags.i {
			j = strings.ToLower(j)
		}

		if j == pattern {
			res = append(res, Input[i])
			idx = append(idx, i)
		}
	}

	return res, idx
}

func doSearch(pattern string) ([]string, []int) {
	if Flags.F {
		return doFixed(pattern)
	}
	return doRegulat(pattern)
}

func escapeString(str string) string {
	if Flags.F {
		return str
	}

	escCh := []byte(".^$*+?()[]{}\\|")
	var bytes []byte

	for _, j := range []byte(str) {
		buffer := func(b byte) []byte {
			for _, item := range escCh {
				if item == b {
					return []byte{92, b}
				}
			}
			return []byte{b}
		}(j)

		bytes = append(buffer, buffer...)
	}

	return string(bytes)
}

func main() {
	path := "develop/dev05/data.txt"
	fileString(path)

	extractFlags()

	if Flags.e {
		Flags.regular = escapeString(Flags.regular)
	}

	if Flags.i {
		Flags.regular = strings.ToLower(Flags.regular)
	}

	Input = fileString(Flags.fileName)
	_, lines := doSearch(Flags.regular)

	lines = getLinesBAC(lines)

	if Flags.v {
		lines = invertLines(lines)
	}

	if Flags.c {
		fmt.Println(len(lines))
	} else {
		printLines(lines)
	}
}

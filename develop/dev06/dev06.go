package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var OSFlags osFlags

type osFlags struct {
	Fields        []int
	Delim         byte
	SeparatedOnly bool
}

func fieldsSetup(arg string) {
	if arg == "" {
		return
	}

	fields := strings.Split(arg, ",")
	for _, i := range fields {
		f, err := strconv.Atoi(i)

		if err != nil {
			log.Fatal("Error in fielsd line")
		}
		OSFlags.Fields = append(OSFlags.Fields, f)
	}
}

func delimSetup(arg string) {
	OSFlags.Delim = arg[0]
}

func readFlags() {
	var Fields string
	var Delim string

	flag.StringVar(&Fields, "f", "", "выбор поля для отображения. Пример: -f 4,6,9")
	flag.StringVar(&Delim, "d", "\t", "разделитель для полей. Пример -d \\s")
	flag.BoolVar(&OSFlags.SeparatedOnly, "s", false, "вывод только тех строк, что содержат разделитель")
	flag.Parse()

	if Fields == "" {
		log.Fatalf("Введите номера столбцов для отображения!\nПример: -f 4,6,9\n")
	}

	fieldsSetup(Fields)
	delimSetup(Delim)
}

func getSeparated(separated []string) string {
	resLine := ""

	for _, i := range OSFlags.Fields {
		if i-1 < len(separated) {
			resLine += separated[i-1] + string(OSFlags.Delim)
		}
	}

	resLine = resLine[:len(resLine)-1]
	return resLine
}

func endlessRead() []string {
	var lines []string

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()
		if line == "exit" {
			break
		}

		separated := strings.Split(line, string(OSFlags.Delim))
		if len(separated) > 1 {
			resLine := getSeparated(separated)
			lines = append(lines, resLine)
		}

		if (len(separated) == 1) && (OSFlags.SeparatedOnly == false) {
			lines = append(lines, line)
			continue
		}
	}

	return lines
}

func main() {
	readFlags()
	lines := endlessRead()

	for _, s := range lines {
		fmt.Println(s)
	}
}

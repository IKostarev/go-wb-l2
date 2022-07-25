package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadFile(file string) []string { // читаю файл
	var str []string

	f, err := os.Open(file)
	defer f.Close()

	if err != nil {
		log.Fatalf("Ошибка при чтении файла - %s", err)
	}

	scan := bufio.NewScanner(f)

	for scan.Scan() {
		str = append(str, scan.Text())
	}

	if err := scan.Err(); err != nil {
		log.Fatalf("Ошибка при сканировании текста - %s", err)
	}

	return str
}

func WriteFile(file string, str []string) { // пишу в файл
	f, err := os.OpenFile(file, os.O_TRUNC|os.O_RDWR, 0664)
	defer f.Close()

	if err != nil {
		log.Fatalf("Ошибка при записи - %s", err)
	}

	for _, j := range str {
		f.WriteString(j + "\n")
	}
}

func DeleteDuplicate(str []string) []string { // удаляю копии
	var res []string

	temp := make(map[string]struct{})

	for _, i := range str {
		if _, ok := temp[i]; ok == false {
			temp[i] = struct{}{}
			res = append(res, i)
		}
	}

	return res
}

func ReverseSlice(str []string) []string { // разворачиваю слайс
	if len(str) < 2 {
		return str
	}

	for i := 0; i < len(str)/2; i++ {
		str[i], str[len(str)-1-i] = str[len(str)-1-i], str[i]
	}

	return str
}

func Key(keys []string, key string) (int, bool) { // определяю введенные ключи
	for i, k := range keys {
		if k == key {
			return i, true
		}
	}

	return -1, false
}

func SortNumbers(str []string, left, right int, collection []string) []string { // сортирую по цифрам
	l := left
	r := right
	center, _ := strconv.Atoi(collection[(left+right)/2])

	for l <= r {
		collectionLeft, _ := strconv.Atoi(collection[l])
		for collectionLeft < center {
			l++
			collectionLeft, _ = strconv.Atoi(collection[l])
		}

		collectionRight, _ := strconv.Atoi(collection[r])
		for collectionRight > center {
			r--
			collectionRight, _ = strconv.Atoi(collection[r])
		}

		if l <= r {
			str[r], str[l] = str[l], str[r]
			collection[l], collection[r] = collection[l], collection[r]
			l++
			r--
		}
	}

	if r > left {
		SortNumbers(str, left, r, collection)
	}

	if l < right {
		SortNumbers(str, l, right, collection)
	}

	return str
}

func Compare(str, collection []string, l, r, idx int) { // совпадение символов
	if (len(collection[l]) < idx+1) || (len(collection[r]) < idx+1) {
		return
	}

	if collection[l][idx] == collection[r][idx] {
		Compare(str, collection, l, r, idx)
	}

	if collection[l][idx] > collection[r][idx] {
		collection[l], collection[r] = collection[r], collection[l]
		str[l], str[r] = str[r], str[l]
		return
	}
}

func SortQuick(str []string, left, right int, collection []string) []string { // быстрая сортировка
	l := left
	r := right
	center := collection[(left+right)/2][0]

	for l <= r {
		for collection[l][0] < center {
			l++
		}

		for collection[r][0] > center {
			r--
		}

		if l <= r {
			str[r], str[l] = str[l], str[r]
			collection[l], collection[r] = collection[r], collection[l]

			if collection[r][0] == collection[l][0] {
				Compare(str, collection, l, r, 0)
			}
			l++
			r--
		}
	}

	if r > left {
		SortQuick(str, left, r, collection)
	}

	if l < right {
		SortQuick(str, l, right, collection)
	}

	return str
}

func PrepareSorted(str []string, key string) []string { // подготовительная сортировка, получаю строку и ключи
	var (
		collection []string
		prepared   []string
		wrong      []string
	)

	keys := strings.Split(key, " ")

	if i, ok := Key(keys, "-k"); ok == false { // проверка ключа -k
		if (i + 1) > len(keys)-1 {
			return append([]string{"Ошибка ввода"}, str...)
		}

		var (
			tempIn         []string
			tempCollection []string
		)

		k, err := strconv.Atoi(keys[i+1])

		if err != nil {
			return append([]string{"Неправильный параметр для колонки"}, str...)
		}

		for _, j := range str {
			tempStr := strings.Split(j, " ")

			if len(tempStr) >= k {
				tempIn = append(tempIn, j)
				tempCollection = append(tempCollection, tempStr[k-1])
			} else {
				wrong = append(wrong, j)
			}
		}

		if len(tempIn) == 0 {
			return append([]string{"Нет столбца"}, str...)
		}

		prepared = tempIn
		collection = tempCollection
	} else {
		prepared = str
		collection = str
	}

	if _, ok := Key(keys, "-n"); ok == true { // ключ -n
		var (
			tempIn         []string
			tempCollection []string
		)

		reg := regexp.MustCompile("^[0-9]+")

		for i, s := range collection {
			numSub := reg.FindString(s)

			if numSub != "" {
				tempIn = append(tempIn, prepared[i])
				tempCollection = append(tempCollection, numSub)
			} else {
				wrong = append(wrong, prepared[i])
			}
		}

		if len(tempIn) == 0 {
			return append([]string{"Подходящих по параметрам строк не найдено"}, str...)
		}

		prepared = tempIn
		collection = tempCollection

		str = SortNumbers(prepared, 0, len(prepared)-1, collection)

		if _, ok := Key(keys, "c"); ok == true { // ключ с
			str = ReverseSlice(str)
		}

		return str
	}

	str = SortQuick(prepared, 0, len(prepared)-1, collection)
	if _, ok := Key(keys, "c"); ok == true {
		str = ReverseSlice(str)
	}

	str = append(str, wrong...)
	if _, ok := Key(keys, "u"); ok == true { // ключ u
		str = DeleteDuplicate(str)
	}

	return str
}

func SortStrings(str []string, key string) []string { // подготавливаю строку сортируя ее
	return PrepareSorted(str, key)
}

func Printer(str []string) { // печатаю
	fmt.Println("--------------------")

	for _, i := range str {
		fmt.Println(i)
	}

	fmt.Println("--------------------")
}

func main() {
	path := "develop/dev03/test_data.txt"
	data := ReadFile(path)

	result := SortStrings(data, "-k 3 -u")

	for _, i := range result {
		fmt.Println(i)
	}

	WriteFile(path, result)
}

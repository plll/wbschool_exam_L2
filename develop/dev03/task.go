package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

# Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

# Дополнительное

# Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.

Вводить флаги потом название файла
*/

func removeDuplicates[T comparable](slice []T) []T {
	seen := make(map[T]bool)
	result := make([]T, 0)

	for _, num := range slice {
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}

	return result
}

func reverseSlice[T comparable](slice []T) []T {
	for i := 0; i < len(slice)/2; i++ {
		slice[i], slice[len(slice)-i-1] = slice[len(slice)-i-1], slice[i]
	}
	return slice
}

func sortByColumn[T any](slice []T) []T {
	//TOD0
	return slice
}

func main() {
	var numericSort, deleteRepeats, reverseResult bool
	var columnSort int
	flag.BoolVar(&numericSort, "n", false, "Sort numerically")
	flag.BoolVar(&deleteRepeats, "u", false, "Delete repeats")
	flag.BoolVar(&reverseResult, "r", false, "Reverse result")
	flag.IntVar(&columnSort, "k", 0, "Sort by column")
	flag.Parse()

	fmt.Println(numericSort)
	args := flag.Args()
	if args[0] == "" {
		log.Fatal("input filename to sort")
	}
	f, err := os.Open(args[0])
	if err != nil {
		panic(err)
	}
	rows := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		rows = append(rows, line)
	}
	f.Close()

	if deleteRepeats {
		rows = removeDuplicates(rows)
	}

	if columnSort != 0 && !numericSort {

	} else if columnSort != 0 && numericSort {

	} else if numericSort {
		rowsInt := make([]int, 0, len(rows))
		for i := 0; i < len(rows); i++ {
			j, err := strconv.Atoi(rows[i])
			if err != nil {
				panic(err)
			}
			rowsInt = append(rowsInt, j)
		}
		slices.Sort(rowsInt)
		fmt.Println(rowsInt)
		for i := 0; i < len(rowsInt); i++ {
			rows[i] = strconv.Itoa(rowsInt[i])
		}
	} else {
		slices.Sort(rows)
	}
	if reverseResult {
		rows = reverseSlice(rows)
	}
	for _, x := range rows {
		fmt.Println(x)
	}
}

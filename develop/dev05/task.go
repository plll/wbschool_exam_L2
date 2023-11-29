package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
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

func main() {
	var after, before, context int
	var count, ignoreCase, invert, fixed, lineNumber bool
	flag.IntVar(&after, "A", 0, "печатать +N строк после совпадения")
	flag.IntVar(&before, "B", 0, "печатать +N строк до совпадения")
	flag.IntVar(&context, "C", 0, "(A+B) печатать ±N строк вокруг совпадения")
	flag.BoolVar(&count, "c", false, "Вывести количество строк")
	flag.BoolVar(&invert, "v", false, "Вместо совпадения, исключать")
	flag.BoolVar(&ignoreCase, "i", false, "Игнорировать регистр")
	flag.BoolVar(&fixed, "F", false, "Точное совпадение со строкой, не паттерн")
	flag.BoolVar(&lineNumber, "n", false, "Напечатать номер строки")

	flag.Parse()

	args := flag.Args()
	if args[0] == "" {
		log.Fatal("input filename to grep")
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
	pattern := args[1]
	if ignoreCase {
		pattern = strings.ToLower(pattern)
	}
	matchRows := make([]int, 0)
	var line string
	for i := 0; i < len(rows); i++ {
		if ignoreCase {
			line = strings.ToLower(rows[i])
		} else {
			line = rows[i]
		}
		if fixed && !invert {
			if line == pattern {
				matchRows = append(matchRows, i)
			}
		} else if strings.Contains(line, pattern) && !invert {
			matchRows = append(matchRows, i)
		} else if !strings.Contains(line, pattern) && invert {
			matchRows = append(matchRows, i)
		}
	}
	lenMatch := len(matchRows)
	for i := 0; i < lenMatch; i++ {
		if after != 0 {
			for j := matchRows[i]; j < matchRows[i]+after+1; j++ {
				if j < len(rows) {
					matchRows = append(matchRows, j)
				}
			}
		}
		if before != 0 {
			for j := matchRows[i]; j != matchRows[i]-before-1; j-- {
				if j >= 0 {
					matchRows = append(matchRows, j)
				}
			}
		}
		if context != 0 {
			for j := matchRows[i]; j != matchRows[i]-context-1; j-- {
				if j >= 0 {
					matchRows = append(matchRows, j)
				}
			}
			for j := matchRows[i]; j < matchRows[i]+context+1; j++ {
				if j < len(rows) {
					matchRows = append(matchRows, j)
				}
			}
		}
	}
	matchRows = removeDuplicates(matchRows)
	slices.Sort(matchRows)
	if count {
		fmt.Println(len(matchRows))
		return
	}
	for _, x := range matchRows {
		if lineNumber {
			fmt.Println(x, rows[x])
		} else {
			fmt.Println(rows[x])
		}
	}
}

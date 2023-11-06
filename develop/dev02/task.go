package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func unpackString(s string) string {
	runes := []rune(s)
	_, ok := strconv.Atoi(s)
	if ok == nil {
		log.Fatal(errors.New("invalid string"))
	}
	unpackedRunes := make([]rune, 0)
	for i := 0; i < len(runes); i++ {
		if !unicode.IsDigit(runes[i]) && runes[i] != '\u005C' {
			fmt.Println(runes[i], i, len(runes)-1)
			if i < len(runes)-1 {
				if unicode.IsDigit(runes[i+1]) {
					for j := 0; j < int(runes[i+1]-'0'); j++ {
						unpackedRunes = append(unpackedRunes, runes[i])
					}
				} else {
					unpackedRunes = append(unpackedRunes, runes[i])
				}
			} else {
				unpackedRunes = append(unpackedRunes, runes[i])
			}
		} else if runes[i] == '\u005C' && unicode.IsDigit(runes[i+1]) {
			if runes[i-1] == '\u005C' {
				for j := 0; j < int(runes[i+1]-'0'); j++ {
					unpackedRunes = append(unpackedRunes, runes[i])
				}
			} else {
				if i < len(runes)-2 && unicode.IsDigit(runes[i+2]) {
					for j := 0; j < int(runes[i+2]-'0'); j++ {
						unpackedRunes = append(unpackedRunes, runes[i+1])
					}
				} else {
					unpackedRunes = append(unpackedRunes, runes[i+1])
				}
			}
		}
	}
	return string(unpackedRunes)
}

func main() {
	str := "qwe\\4\\5"
	fmt.Println(unpackString(str))
}

package main

import (
	"fmt"
	"reflect"
	"slices"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func findAnagramms(words []string) map[string][]string {
	result := make(map[string][]string)
	setWords := make([]map[rune]int, 0)
	// Приводим к нижнему регистру + кастим в слайс рун слово и строим счетчик рун, чтобы в дальнейшем находить анограммы
	for i := 0; i < len(words); i++ {
		words[i] = strings.ToLower(words[i])
		wordRunes := []rune(words[i])
		runesCounter := make(map[rune]int)
		for _, x := range wordRunes {
			runesCounter[x] += 1
		}
		setWords = append(setWords, runesCounter)
	}
	// Заводим слайс мап счетчиков букв, чтобы не переиспользовать слова
	usedWords := make([]map[rune]int, 0)
	// Для каждой анограммы проходим по слайсу и находим все схожие мапы счетчиков букв, пропускаем те, мапы счетчиков которые есть в UsedWords и в конце сортируем множество анограмм и добавляем в результат ( маппу слайсов стрингов)
	for i := 0; i < len(setWords)-1; i++ {
		row := make([]string, 0)
		found := false
		for _, m := range usedWords {
			if reflect.DeepEqual(m, setWords[i]) {
				found = true
			}
		}
		if !found {
			for j := i + 1; j < len(setWords); j++ {
				if reflect.DeepEqual(setWords[i], setWords[j]) {
					if !slices.Contains(row, words[j]) {
						row = append(row, words[j])
					}
				}
			}
			usedWords = append(usedWords, setWords[i])
			slices.Sort(row)
			result[words[i]] = row
		}
	}
	return result
}

func main() {
	words := []string{"листок", "слиток", "столик", "пятак", "пятка", "тяпка", "тяпка", "цуц"}
	anagrams := findAnagramms(words)
	fmt.Println(anagrams)
}

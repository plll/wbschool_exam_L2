package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type fieldsArr []int

func (i *fieldsArr) String() string {
	return fmt.Sprintf("%v", *i)
}

func (i *fieldsArr) Set(value string) error {
	for _, dt := range strings.Split(value, ",") {
		tmp, err := strconv.Atoi(dt)
		if err != nil {
			return err
		}
		*i = append(*i, tmp)
	}
	return nil
}

func main() {
	var fields fieldsArr
	var separated bool
	var delimiter string
	flag.Var(&fields, "f", "Номера колонок")
	flag.StringVar(&delimiter, "d", ",", "Разделитель")
	flag.BoolVar(&separated, "s", false, "Только строки с разделителем")
	flag.Parse()

	rows := make([][]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		row := scanner.Text()
		if row == "" {
			break
		}
		rowFields := strings.Split(row, delimiter)
		rows = append(rows, rowFields)
	}

	fieldsAmount := len(rows[0])
	for _, row := range rows {
		var output string
		if separated && len(row) != fieldsAmount {
			continue
		}
		for _, x := range fields {
			if output == "" {
				output += row[x]
			} else {
				output = output + delimiter + row[x]
			}
		}
		fmt.Println(output)
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func mapkey(m map[string]int, value int) (key string) {
	for k, v := range m {
		if v == value {
			key = k
			return
		}
	}
	return
}

func main() {
	fmt.Println("Введите задачу")
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	x := strings.Split(text, " ")

	f, err := strconv.Atoi(x[0])
	s, er := strconv.Atoi(x[2])

	if len(x) != 3 {
		fmt.Println("Выдача паники, так как строка не является математической операцией.")
		panic(err)
	}

	if (f > 10) || (f < 1) || (s > 10) || (f < 1) {
		fmt.Println("Выдача паники, так как значение не подходит.")
		panic(err)
	}

	var numbers map[string]int = map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	if err != nil {
		el, status := numbers[x[0]]
		if status == false {
			fmt.Println("Выдача паники, некорректный ввод.")
			panic(err)
		} else {
			f = el
		}
	}

	if er != nil {
		el, status := numbers[x[2]]
		if status == false {
			fmt.Println("Выдача паники, некорректный ввод.")
			panic(er)
		} else {
			s = el
		}
	}

	var answer int
	var answerRyms string

	switch x[1] {
	case "+":
		answer = f + s
	case "-":
		answer = f - s
	case "*":
		answer = f * s
	case "/":
		answer = f / s

	}

	if (answer) > 0 && (answer < 11) {
		answerRyms = mapkey(numbers, answer)
	}

	if (answer) > 10 && (answer < 20) {
		answerRyms = "X" + mapkey(numbers, answer-10)
	}

	if (answer) >= 20 && (answer < 30) {
		answerRyms = "XX" + mapkey(numbers, answer-20)
	}

	if (answer) >= 30 && (answer < 40) {
		answerRyms = "XXX" + mapkey(numbers, answer-30)
	}

	if (answer) >= 40 && (answer < 50) {
		answerRyms = "XL" + mapkey(numbers, answer-40)
	}

	if (answer) >= 50 && (answer < 60) {
		answerRyms = "L" + mapkey(numbers, answer-50)
	}

	if (answer) >= 60 && (answer < 70) {
		answerRyms = "LX" + mapkey(numbers, answer-60)
	}

	if (answer) >= 70 && (answer < 80) {
		answerRyms = "LXX" + mapkey(numbers, answer-70)
	}

	if (answer) >= 80 && (answer < 90) {
		answerRyms = "LXXX" + mapkey(numbers, answer-80)
	}

	if (answer) >= 90 && (answer < 100) {
		answerRyms = "XC" + mapkey(numbers, answer-90)
	}

	if answer == 100 {
		answerRyms = "C"
	}

	if (er == nil) && (err == nil) {
		fmt.Println(answer)
	} else if (er != nil) && (err != nil) && (answer > 0) {
		fmt.Println(answerRyms)
	} else {
		panic(err)
	}

}

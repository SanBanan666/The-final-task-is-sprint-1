package calculator

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func Calc(expression string) (float64, error) {
	// Удаляем пробелы
	expression = strings.ReplaceAll(expression, " ", "")

	// Проверяем выражение на корректность
	if !isValidExpression(expression) {
		return 0, errors.New("некорректное выражение")
	}

	// Преобразуем инфиксное выражение в постфиксное
	postfix, err := infixToPostfix(expression)
	if err != nil {
		return 0, err
	}

	// Вычисляем постфиксное выражение
	result, err := evaluatePostfix(postfix)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func isValidExpression(expression string) bool {
	re := regexp.MustCompile(`^[0-9+\-*/().]+$`)
	return re.MatchString(expression)
}

func infixToPostfix(expression string) ([]string, error) {
	var output []string
	var stack []string

	precedence := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
	}

	for i := 0; i < len(expression); i++ {
		char := string(expression[i])

		if isDigit(char) {
			num := char
			for i+1 < len(expression) && isDigit(string(expression[i+1])) {
				i++
				num += string(expression[i])
			}
			output = append(output, num)
		} else if char == "(" {
			stack = append(stack, char)
		} else if char == ")" {
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				return nil, errors.New("непарные скобки")
			}
			stack = stack[:len(stack)-1] // удаляем '('
		} else if precedence[char] > 0 {
			for len(stack) > 0 && precedence[stack[len(stack)-1]] >= precedence[char] {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, char)
		}
	}

	for len(stack) > 0 {
		output = append(output, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return output, nil
}

func evaluatePostfix(postfix []string) (float64, error) {
	var stack []float64

	for _, token := range postfix {
		if isDigit(token) {
			num, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, err
			}
			stack = append(stack, num)
		} else {
			if len(stack) < 2 {
				return 0, errors.New("недостаточно значений в стеке")
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var result float64
			switch token {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				if b == 0 {
					return 0, errors.New("деление на ноль")
				}
				result = a / b
			}
			stack = append(stack, result)
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("ошибка вычисления")
	}
	return stack[0], nil
}

func isDigit(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

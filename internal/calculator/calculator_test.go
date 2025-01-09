package calculator_test

import (
	. "Sprint1/internal/calculator"
	"testing"
)

func TestCalculateInvalidExpression(t *testing.T) {
	invalidExpressions := []string{
		"2--2",    // Два подряд оператора
		"2+2*",    // Оператор без второго числа
		"2/(2-2)", // Деление на ноль
		"abc",     // Некорректные символы
		"(2+3",    // Непарные скобки
		"2++2",    // Два подряд оператора
		"",        // Пустая строка
		"10/0",    // Деление на ноль
	}

	for _, expr := range invalidExpressions {
		_, err := Calc(expr)
		if err == nil {
			t.Errorf("Ожидалась ошибка для выражения '%s', но получено nil", expr)
		}
	}
}

func TestCalculateValidExpression(t *testing.T) {
	tests := []struct {
		expression string
		expected   float64
	}{
		{"2+2", 4},
		{"10-3", 7},
		{"3*4", 12},
		{"8/2", 4},
		{"(2+3)*4", 20},
		{"10/(2+3)", 2},
	}

	for _, tt := range tests {
		result, err := Calc(tt.expression)
		if err != nil {
			t.Errorf("Ошибка при вычислении выражения '%s': %v", tt.expression, err)
			continue
		}
		if result != tt.expected {
			t.Errorf("Для выражения '%s' ожидалось %v, получено %v", tt.expression, tt.expected, result)
		}
	}
}

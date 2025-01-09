package main

import (
	"Sprint1/internal/calculator"
	"testing"
)

func TestMainCalc(t *testing.T) {
	result, err := calculator.Calc("3+3")
	if err != nil {
		t.Errorf("Ожидалось успешное выполнение, но получена ошибка: %v", err)
	}
	if result != 6 {
		t.Errorf("Ожидалось 6, но получено %v", result)
	}
}

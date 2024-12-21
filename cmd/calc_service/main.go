package main

import (
	"Sprint1/internal/calculator"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

func main() {
	http.HandleFunc("/api/v1/calculate", calculateHandler)

	port := ":8080"
	log.Printf("Server is running on http://localhost%s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Method: %s, Path: %s", r.Method, r.URL.Path)
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil || strings.TrimSpace(req.Expression) == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(Response{Error: "Expression is not valid"})
		return
	}

	result, err := calculator.Calc(req.Expression)
	if err != nil {
		handleCalculationError(w, err)
		return
	}

	resp := Response{Result: formatResult(result)}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func handleCalculationError(w http.ResponseWriter, err error) {
	if errors.Is(err, errors.New("некорректное выражение")) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(Response{Error: "Expression is not valid"})
	} else if strings.Contains(err.Error(), "деление на ноль") {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(Response{Error: "Division by zero is not allowed"})
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{Error: "Internal server error"})
	}
}

func formatResult(result float64) string {
	if result == float64(int(result)) {
		return strconv.Itoa(int(result))
	}
	return strconv.FormatFloat(result, 'f', -1, 64)
}

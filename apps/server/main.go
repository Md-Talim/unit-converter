package main

import (
	"log"
	"net/http"
	"strconv"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		next.ServeHTTP(w, r)
	})
}

func home(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	w.Write([]byte("Building unit coverter project from roadmap.sh"))
}

func convert(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	valueStr := r.URL.Query().Get("value")
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		http.Error(w, "Invalid value", http.StatusBadRequest)
		return
	}

	var result float64
	if from == "grams" && to == "kilograms" {
		result = value / 1000
	} else if from == "kilograms" && to == "grams" {
		result = value * 1000
	} else {
		http.Error(w, "Conversion not supported", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"result":` + strconv.FormatFloat(result, 'f', -1, 64) + `}`))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("POST /convert", convert)

	log.Print("Starting server on :4000")

	err := http.ListenAndServe(":4000", corsMiddleware(mux))
	log.Fatal(err)
}

package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/md-talim/unit-converter/server/pkg/converter"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		next.ServeHTTP(w, r)
	})
}

func convert(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	unitCategory := r.URL.Query().Get("category")
	valueStr := r.URL.Query().Get("value")
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		http.Error(w, "Invalid value", http.StatusBadRequest)
		return
	}

	unitConverter := converter.GetConverter(unitCategory)
	if unitConverter == nil {
		http.Error(w, "Category not supported", http.StatusBadRequest)
		return
	}

	result, err := unitConverter.Convert(value, from, to)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"result":` + strconv.FormatFloat(result, 'f', -1, 64) + `}`))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /convert", convert)

	log.Print("Starting server on :4000")

	err := http.ListenAndServe(":4000", corsMiddleware(mux))
	log.Fatal(err)
}

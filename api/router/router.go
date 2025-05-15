package router

import (
	"net/http"

	"github.com/Alastair7/gonverter/api/resource/file"
)

func SetupRouter() http.Handler {
	router := http.NewServeMux()

	checkhealthHandler := func(w http.ResponseWriter, r *http.Request) {}
	fileHandler := file.NewFileHandler()

	router.HandleFunc("/checkhealth", checkhealthHandler)
	router.Handle("/api/file", checkCORS(fileHandler))

	return router
}

func checkCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		if origin != "*" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}

		w.Header().Add("Vary", "Origin")

		next.ServeHTTP(w, r)
	})
}

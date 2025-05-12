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
	router.Handle("/api/file", fileHandler)

	return router
}

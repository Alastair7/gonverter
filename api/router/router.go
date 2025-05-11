package router

import "net/http"

func SetupRouter() http.Handler {
	router := http.NewServeMux()

	checkhealthHandler := func(w http.ResponseWriter, r *http.Request) {}
	router.HandleFunc("/checkhealth", checkhealthHandler)

	return router
}

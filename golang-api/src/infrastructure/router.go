package infrastructure

import (
	"log"
	"net/http"

	"golang-api/interfaces/controllers"
)

func HandleRouter() {
	http.HandleFunc("/", handleMethod(http.MethodGet, controllers.Index))

	server := &http.Server{
		Addr: ":8081",
	}

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

type handlerType func(http.ResponseWriter, *http.Request) ([]byte, int)

func handleMethod(methodType string, handler handlerType) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == methodType {
			body, httpStatus := handler(w, r)

			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.WriteHeader(httpStatus)
			w.Write(body)
		}

		w.WriteHeader(http.StatusNotFound)
	}
}

package main

import (
	"encoding/json"
	"net/http"

	"github.com/luoying/SimpleGoWepPro/context/middleware"
)

func main() {
	http.HandleFunc("/companies", func(w http.ResponseWriter, r *http.Request) {
		c := Company{
			ID:      123,
			Name:    "Google",
			Country: "USA",
		}

		enc := json.NewEncoder(w)
		enc.Encode(c)
	})

	http.ListenAndServe("localhost:8080", &middleware.TimeoutMiddleware{
		Next: &middleware.AuthMiddleware{},
	})
}

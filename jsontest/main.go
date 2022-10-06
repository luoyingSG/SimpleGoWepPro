package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/companies", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			dec := json.NewDecoder(r.Body)
			company := Company{}
			err := dec.Decode(&company)
			if err != nil {
				log.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			enc := json.NewEncoder(w)
			err = enc.Encode(company)
			if err != nil {
				log.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	})

	http.ListenAndServe("localhost:8080", nil)
}

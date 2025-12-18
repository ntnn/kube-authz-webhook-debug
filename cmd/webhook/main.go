package main

import (
	"encoding/json"
	"log"
	"net/http"

	pp "github.com/k0kubun/pp/v3"
	authorizationv1 "k8s.io/api/authorization/v1"
)

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)
		log.Printf("Headers: %#v", r.Header)

		var data map[string]any
		json.NewDecoder(r.Body).Decode(&data)
		pp.Println(data)

		sar := authorizationv1.SubjectAccessReview{
			Status: authorizationv1.SubjectAccessReviewStatus{
				Allowed: false,
			},
		}
		respBytes, err := json.Marshal(sar)
		if err != nil {
			http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(respBytes)
	})

	server := http.Server{}
	server.Addr = "localhost:9090"
	server.Handler = handler
	server.ListenAndServe()
}

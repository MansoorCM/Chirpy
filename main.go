package main

import (
	"log"
	"net/http"
)

func main() {
	const port = "8080"

	mux := http.NewServeMux()
	config := apiConfig{fileserverHits: 0}

	rootDir := http.Dir(".")
	mux.Handle("/app/*", config.middleWareMetricsInc(http.StripPrefix("/app", http.FileServer(rootDir))))

	logoPath := "/assets/logo.png"
	mux.HandleFunc("/app"+logoPath, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "."+logoPath)
	})

	mux.HandleFunc("GET /api/healthz", handlerReadiness)

	mux.HandleFunc("GET /admin/metrics", config.middleWareMetricsGet)

	mux.HandleFunc("GET /api/reset", config.middleWareMetricsReset)

	corsMux := middlewareCors(mux)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

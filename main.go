package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"himanshu.com/sample/services"
	"himanshu.com/sample/setup"
)

func commonMiddleware(next http.Handler) http.Handler {
	config, err := setup.GetConfig()
	if err != nil {
		panic(err)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("host", config.Host)
		w.Header().Add("dbname", config.DB.DBName)
		w.Header().Add("dbuser", config.DB.Username)
		w.Header().Add("dbpassword", config.DB.Password)
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()
	r.Use(commonMiddleware)
	r.HandleFunc("/setup", services.SetupHandler).Methods("GET")
	r.HandleFunc("/register", services.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", services.LoginHandler).Methods("POST")
	r.HandleFunc("/verifyToken", services.VerifyTokenHandler).Methods("POST")
	http.ListenAndServe(":8888", r)
}

package main

import (
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gorilla/mux"
)

var port string

func init() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Println("No port specified, using 8080 as default.")
		port = "8080" // default value if no PORT env variable is set
	}
}

func parseEnvURL(URLEnv string) *url.URL {
	envContent := os.Getenv(URLEnv)
	parsedURL, err := url.ParseRequestURI(envContent)
	if err != nil {
		log.Fatal("Not a valid URL for env variable ", URLEnv, ": ", envContent, "\n")
	}

	return parsedURL
}

func parseEnvVar(envVar string) string {
	envContent := os.Getenv(envVar)

	if len(envContent) == 0 {
		log.Fatal("Env variable ", envVar, " missing, exiting.")
	}

	return envContent
}

// HealthHandler responds to /healthz endpoint for application monitoring
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/healthz", HealthHandler)
	router.HandleFunc("/login/oidc", OIDCHandler)
	router.HandleFunc("/login", LoginHandler)
	// router.HandleFunc("/logout", LogoutHandler) // TODO
	router.HandleFunc("/", AuthReqHandler) // TODO convert to wildcard

	log.Fatal(http.ListenAndServe(":8080", router))
}
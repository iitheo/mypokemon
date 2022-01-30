package main

import (
	"fmt"
	httproutespokemon "github.com/iitheo/theopokemon/pkg/api/routes/v1/httproutes"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func init() {

	env := strings.TrimSpace(strings.ToLower(os.Getenv("SERVERPOKEMON")))

	switch env {
	case "prod":
		//todo
	default:
		os.Setenv("PORT", "8081")
	}

}

func main() {
	port := os.Getenv("PORT")

	log.Printf("server running at %s", port)
	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           httproutespokemon.Router(),
		TLSConfig:         nil,
		ReadTimeout:       45 * time.Second,
		ReadHeaderTimeout: 0,
		WriteTimeout:      45 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Printf("Error starting server - %s\n", err.Error())
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// ComputerList contains all Computers who we can use to work with
var ComputerList []Computer

func main() {

	// Start Processing Shell Arguments or use Default Values defined i const.go
	httpPort, computerFilePath := processShellArgs()

	// Process Environment Variables
	httpPort, computerFilePath = processEnvVars(httpPort, computerFilePath)

	// Loading Computer CSV File to Memory File in Memory
	var loadComputerCSVFileError error
	if ComputerList, loadComputerCSVFileError = loadComputerList(computerFilePath); loadComputerCSVFileError != nil {
		log.Fatalf("Failed to load \"%s\". Please make sure there's no syntax or permission error.", computerFilePath)
	}

	// Init HTTP Router - mux
	router := mux.NewRouter()

	// Define Home Route
	router.HandleFunc("/", renderHomePage).Methods("GET")

	// Define Wakeup Api functions with a Computer Name
	router.HandleFunc("/api/wakeup/computer/{computerName}", restWakeUpWithComputerName).Methods("GET")
	router.HandleFunc("/api/wakeup/computer/{computerName}/", restWakeUpWithComputerName).Methods("GET")

	// Setup Webserver
	httpListen := fmt.Sprint(":", httpPort)
	log.Printf("Serving on \"%s\".", httpListen)

	log.Fatal(http.ListenAndServe(httpListen, handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(router)))
}

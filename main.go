package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/krol3/container_go/pkg"
)

// go run main.go run <cmd> <args>
func main() {

	// handle route using handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to new server!")
	})
	http.HandleFunc("/run", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GET params were:", r.URL.Query())

		// if only one expected
		param1 := r.URL.Query().Get("cmd")
		if param1 != "" {
			pkg.ExecuteRun(param1)
		}
	})

	// listen to port
	http.ListenAndServe(":5050", nil)

	if len(os.Args) <= 1 {
		fmt.Println("Missing the param")
		log.Println("Default param 0: %s", os.Args[0])
		os.Exit(1)
	} else {

		switch os.Args[1] {
		case "run":
			println("run")
			pkg.ExecuteRun("ls")
		default:
			log.Fatal("Unknown command. Use run <command_name>, like `run /bin/bash` or `run echo hello`")
			panic("help")
		}
	}
}

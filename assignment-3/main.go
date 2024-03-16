package main

import (
	"assignment-3/controllers"
	"fmt"
	"net/http"
)

var PORT = ":8080"

func main() {
	go controllers.UpdateData()
	http.HandleFunc("/status", controllers.StatusMonitoring)

	fmt.Println("app listen on port", PORT)
	http.ListenAndServe(PORT, nil)
}

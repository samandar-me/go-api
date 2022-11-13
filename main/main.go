package main

import (
	router2 "awesomeProject/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server is running...")
	router := router2.Router()
	log.Fatal(http.ListenAndServe(":4000", router))
}

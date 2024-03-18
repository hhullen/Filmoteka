package main

import (
	"controllers"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	controllers.NewControllerREST("/api/v1", mux)
	fmt.Println(http.ListenAndServe(":8888", mux))
}

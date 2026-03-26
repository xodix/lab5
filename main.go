package main

import (
	"fmt"
	"net/http"
)

const (
	PORT = 8080
)

func main() {
	addr := fmt.Sprintf("localhost:%v", PORT)
	http.ListenAndServe()
}

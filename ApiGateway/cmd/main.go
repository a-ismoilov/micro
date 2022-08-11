package main

import (
	"github.com/akbarshoh/microOLX/api"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", api.Handle())
}

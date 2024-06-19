package main

import (
	"fmt"
	"net/http"
)

func public() http.Handler {
	fmt.Println("building static files for development")
	return http.StripPrefix("/public/", http.FileServer(http.Dir("public")))
}

package server

import (
	"fmt"
	"log"
	"net/http"
)

type RoutingInterface struct {
	Route string
	Callback func()
}

func Serve(port int, routes []RoutingInterface) {
	fmt.Printf("Serving at http://localhost:%d/\n", port)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func ServeImages(path string) {
	http.Handle("/", http.FileServer(http.Dir("./dist")))

	http.Handle("/view", http.FileServer(http.Dir("./static")))

	fmt.Printf("Starting server at port 8080\n")
	fmt.Printf("Serving at http://localhost:8080/%s\n", path)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
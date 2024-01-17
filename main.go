package main

import (
	"fmt"
	"net/http"

	"github.com/devhulk/gooser-web/handlers"
)

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/get-sites", handlers.GetSites)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	fmt.Println("Listening on port 3333")
	http.ListenAndServe(":3333", nil)

}

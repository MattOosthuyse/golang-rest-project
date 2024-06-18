package server

import (
	"fmt"
	"net/http"
)

func RestServer() {
	fmt.Println("In the rest server")

	http.HandleFunc("/", home)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed to listen and crashed")
	}
}

func home(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprint(w, "Hello from the home page")
	if err != nil {
		fmt.Println("Unable to handle home function")
		return
	}

}

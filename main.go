package main

import (
	"fmt"
	"net/http"
	"os"
)

func holaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "Hola")
	} else {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Local por defecto
	}

	http.HandleFunc("/hola", holaHandler)

	fmt.Printf("Servidor corriendo en http://localhost:%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

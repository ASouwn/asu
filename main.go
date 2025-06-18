package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ASouwn/asu/src-asu/router"
)

func main() {
	port := "3000"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	router.RouterInit("./src/app")
	fmt.Printf("Starting server on port http://localhost:%s ...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

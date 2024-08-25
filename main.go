package main

import (
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("routes/index.go")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	// Print the content (optional)
	fmt.Println("Content of example.go:")
	fmt.Println(string(content))

	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "go-run")
	if err != nil {
		fmt.Println("Error creating temp directory:", err)
		return
	}
	defer os.RemoveAll(tempDir)

	// r := chi.NewRouter()
	// r.Use(middleware.Logger)

	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("welcome"))
	// })
	// http.ListenAndServe(":3000", r)
}

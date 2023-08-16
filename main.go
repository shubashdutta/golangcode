package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/rs/cors"
    "github.com/shubash/saibaba/router"
)

func setupCors() *cors.Cors {
    return cors.New(cors.Options{
        AllowedOrigins:   []string{"https://saibabasevasadantrust.com"}, // Your frontend's HTTPS URL
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Authorization", "Content-Type"},
        AllowCredentials: true,
    })
}

func main() {
    r := router.Router()

    // Set up CORS middleware
    c := setupCors()
    handler := c.Handler(r)

    port := "8081" // Change the port to your desired value

    fmt.Printf("Server is ready and running on port %s...\n", port)
    log.Fatal(http.ListenAndServe(":"+port, handler))
}

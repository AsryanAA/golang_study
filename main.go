package main

import (
    "context"
    "fmt"
    "log"

    "google.golang.org/genai"
)

func main() {
    ctx := context.Background()
    client, err := genai.NewClient(ctx, &genai.ClientConfig{
        APIKey:  "AIzaSyBGf7Ga1eyy7t7AKiXJ0nOyPnFVzCcjtOU",
        Backend: genai.BackendGeminiAPI,
    })

    if err != nil {
        log.Fatal(err)
    }

    result, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.5-flash",
        genai.Text("Напиши красивое стихотворение про дочку"),
        nil,
    )
	
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(result.Text())
}
package main

import (
	"chatgpt-api-example/handler"
	"fmt"
	"log"
	"net/http"
)

const Port = 8080

func main() {
	grammarCheckerHandler := handler.NewGrammarCheckerHandler()

	http.HandleFunc("/grammar-check", grammarCheckerHandler.Check)
	log.Println("[HTTP] serving server on Port", Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", Port), nil); err != nil {
		log.Fatal("[HTTP] failed to serve server")
	}
}

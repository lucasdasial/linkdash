package main

import (
	"fmt"
	"log"
	"lucasdasial/rupi/internal/shortener"
	"os"
)

func main() {

	key := os.Getenv("SECRET_KEY")

	s, err := shortener.New(key)
	if err != nil {
		log.Fatalf("Shortener: %v", err)
	}

	shorted, err := s.Encode("https://example.com")
	if err != nil {
		log.Fatalf("Shortener: %v", err)
	}

	fmt.Printf("Valor encurtado: %v\n", shorted)

	deshorted, err := s.Decode(shorted)

	if err != nil {
		log.Fatalf("Shortner: %v", err)
	}

	fmt.Printf("Valor origianal: %v\n", deshorted)

}

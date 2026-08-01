package main

import (
	"fmt"
	"log"
	"lucasdasial/rupi/internal/shortener"
)

func main() {

	s, err := shortener.New("secret")
	if err != nil {
		log.Fatalf("shortener: %v", err)
	}

	short, err := s.Shorten("https://example.com")
	if err != nil {
		log.Fatalf("shortener: %v", err)
	}

	fmt.Println(short)

}

package main

import (
	"fmt"
	"lucasdasial/rupi/internal"
)

func main() {

	fmt.Println("RUPI:::Seja bem vindo! 🖖")
	fmt.Println("Qual url deseja encurtar?")

	askLink()

	for again() {
		askLink()
	}

	fmt.Println("Até a próxima! 🖐️")

}

func askLink() {
	var url string
	fmt.Scan(&url)
	fmt.Print("Link encurtado: ")
	fmt.Print("rupi.click/")
	fmt.Println(internal.GetShort(url))
}

func again() bool {
	fmt.Println("Desejar encurtar outro link?")

	var again bool
	fmt.Scan(&again)

	return again
}

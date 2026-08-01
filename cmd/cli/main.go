package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {

	fmt.Println(titleStyle.Render("RUPI · encurtador de links"))
	fmt.Println(hintStyle.Render("Seja bem vindo! 🖖"))
	fmt.Println()

	askLink()

	for again() {
		askLink()
	}

	fmt.Println(byeStyle.Render("Até a próxima! 🖐️"))

}

func askLink() {
	fmt.Println(promptStyle.Render("Qual link deseja encurtar?"))
	fmt.Print("> ")

	var link string
	fmt.Scan(&link)

	short := fmt.Sprintf("rupi.click/%s", "lucas")

	fmt.Println(resultLabelStyle.Render("Link encurtado:"))
	fmt.Println(resultBoxStyle.Render(short))
}

func again() bool {
	options := []string{"sim", "não", "s", "n"}
	optsMsg := fmt.Sprintf("Opções: %s", strings.Join(options, " | "))

	fmt.Println(promptStyle.Render("Deseja encurtar outro link?"))
	fmt.Println(hintStyle.Render(optsMsg))
	fmt.Print("> ")

	var awnser string
	fmt.Scan(&awnser)

	for !isCorrectAwnser(&awnser, options, optsMsg) {
		fmt.Print("> ")
		fmt.Scan(&awnser)
	}

	switch awnser {
	case "sim", "s":
		return true
	default:
		return false

	}

}

func isCorrectAwnser(awnser *string, options []string, message string) bool {

	if !slices.Contains(options, *awnser) {
		fmt.Println(errorStyle.Render(fmt.Sprintf("Resposta incompatível, use uma das seguintes opções: %s", message)))
		return false
	}

	return true
}

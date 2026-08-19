package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func askNumber(){
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println("==== É número primo ? ====")

	for {
		fmt.Print("Digite um número inteiro para saber se é primo: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Não foi possível fazer a leitura.\n")
			continue
		}

		number, err := strconv.ParseUint(strings.TrimSpace(input), 10, 64)
		if err != nil {
			fmt.Println("Informe um número válido (maior que zero).\n")
			continue
		}

		intNumber := int(number)

		isCousin := isACousinNumber(intNumber)

		if isCousin {
			fmt.Printf("Número %d é primo\n", number)
		} else {
			fmt.Printf("Número %d não é primo\n", number)
		}
	}
}

func isACousinNumber(number int) bool {
	if number == 1 {
		return false
	}

	if number == 2 {
		return true
	}

	if number % 2 == 0 {
		return false
	}

	//simplificar para: é divisivel por 3, 5 e 7 ?
	for i := 3; i < number; i++ {
		if number % i == 0 {
			return false
		}
	}

	return true
}

func main(){
	askNumber()
}
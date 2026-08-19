/*
Calculadora de Saúde
Você está construindo um app de saúde. Há um cadastro de usuários com nome, peso (kg) e altura (m). Ao medir novamente uma pessoa, os novos valores chegam separados e você precisa atualizar o registro existente — não criar um novo.

Implemente uma função que, dado um usuário e novos valores de peso e altura, atualize o struct original. Em seguida, calcule o IMC (peso/altura²) e classifique: abaixo do peso (< 18.5), normal (18.5–24.9), sobrepeso (25–29.9), obesidade (≥ 30).

Dica: O que garante que o struct original será modificado? Pense em como passar dados mutáveis para funções em Go.
*/
package main

import (
"bufio"
"fmt"
"os"
"strconv"
"strings"
)

type user struct {
	name string
	alt float32
	pes float32
	imc float32
}

func askUserData() (clearName string, fltAlt, fltPes float32){
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println("\n----- Calculadora de Saúde -----")

	for (clearName == ""){
		fmt.Print("Informe seu nome: ")
      	name, _ := reader.ReadString('\n') //aspas simples = caracter, dupla = string
      	clearName = strings.TrimSpace(name)
    }
      	
    for (fltAlt <= 0){
      	fmt.Print("Informe sua altura (m): ")
      	alt, _ := reader.ReadString('\n')
      	intAlt, err := strconv.ParseFloat(strings.TrimSpace(alt), 32)
      		
      	if err != nil {
      		fmt.Print("Informe um valor numérico válido\n")
      		continue
		}
			
		fltAlt = float32(intAlt)
    }
      	
    for (fltPes <= 0){
      	fmt.Print("Informe seu peso (kg): ")
      	pes, _ := reader.ReadString('\n')
      	intPes, err := strconv.ParseFloat(strings.TrimSpace(pes), 32)
      		
      	if err != nil {
      		fmt.Print("Informe um valor numérico válido\n")
      		continue
		}
			
		fltPes = float32(intPes)
    }
      	
    return
}

func healthCalculator(fltAlt, fltPes float32) (imc float32){
	imc = fltPes/(fltAlt * fltAlt)
    
    if (imc < 18.5) {
    	fmt.Println("IMC: abaixo do peso")
    } else if (imc < 24.9) {
    	fmt.Println("IMC: normal")  
    } else if (imc < 29.9) {
		fmt.Println("IMC: sobrepeso")
    } else {
    	fmt.Println("IMC: obesidade")
    }
    
    return
}

func main(){
	m := make(map[string]user)
	
	for true {
		clearName, fltAlt, fltPes := askUserData()
				
		fltImc := healthCalculator(fltAlt, fltPes)
		
		m[clearName] = user{clearName, fltAlt, fltPes, fltImc}
		
		fmt.Println("--- Valores salvos: ---")
		
		for key, val := range m {
			fmt.Printf("pacientes antigos: %s, imc: %.2f\n", key, val.imc)
		}
	}
}

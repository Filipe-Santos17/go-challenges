/*
1. As Temperaturas da Semana
Uma estação meteorológica envia, todo dia à meia-noite, sete temperaturas (uma para cada dia da semana anterior) em graus Celsius. Você recebe esses valores e precisa:

Calcular a média semanal, a menor e a maior temperatura.
Identificar se houve algum dia com temperatura abaixo de zero (geada).
As temperaturas devem ser armazenadas em uma estrutura fixa de sete posições.
Dica: Considere o tipo numérico mais adequado e uma estrutura de dados com tamanho fixo. Use uma constante para o número de dias.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const WeekDays = 7

var weekDays = [...]string{
	"Segunda",
	"Terça",
	"Quarta",
	"Quinta",
	"Sexta",
	"Sábado",
	"Domingo",
}

type WeekStats struct {
	Average   float64
	Min       float64
	Max       float64
	FrostDays int
}

func askUserDt() [WeekDays]float64 {
	reader := bufio.NewReader(os.Stdin)
	var temperatures [WeekDays]float64

	fmt.Println("===== Temperaturas da Semana =====")

	for i, day := range weekDays {
		for {
			fmt.Printf("Digite a temperatura da %s: ", day)

			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Não foi possível fazer a leitura.")
				continue
			}

			temp, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
			if err != nil {
				fmt.Println("Informe um número válido.")
				continue
			}

			temperatures[i] = temp
			break
		}
	}

	return temperatures
}

func calcWeeksTemp(temperatures [WeekDays]float64) WeekStats {
	stats := WeekStats{
		Min: temperatures[0],
		Max: temperatures[0],
	}

	for _, temp := range temperatures {
		if temp < stats.Min {
			stats.Min = temp
		}

		if temp > stats.Max {
			stats.Max = temp
		}

		if temp < 0 {
			stats.FrostDays++
		}

		stats.Average += temp
	}

	stats.Average /= WeekDays

	return stats
}

func main() {
	temperatures := askUserDt()
	stats := calcWeeksTemp(temperatures)

	fmt.Printf(`
===== Resumo da Semana =====
Média: %.1f°C
Menor temperatura: %.1f°C
Maior temperatura: %.1f°C
Dias com geada: %d
`, stats.Average, stats.Min, stats.Max, stats.FrostDays)
}
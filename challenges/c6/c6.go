package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Task struct {
	text   string
	status bool
}

var listTasks []Task
var reader = bufio.NewReader(os.Stdin)

func showOptions() {
	fmt.Print(`
===== Opções =====
(1) Adicionar uma tarefa (texto livre)
(2) Listar todas as tarefas com índice
(3) Marcar uma tarefa como concluída (pelo índice)
(4) Remover uma tarefa (pelo índice)
(5) Sair
Escolha sua opção: `)
}

func askUserAction() {
	for {
		showOptions()

		opt, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Não foi possivel ler a resposta!")
			continue
		}

		optNum, err := strconv.ParseUint(strings.TrimSpace(opt), 10, 64)

		if err != nil {
			fmt.Println("Opção invalida!")
			continue
		}

		choseAction(optNum)
	}
}

func addTask() {
	for {
		fmt.Print("Escreva a tarefa que deseja adicionar (enter para adicionar): ")

		task, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Não foi possivel adicionar sua task!")
		}

		clearTaskText := strings.TrimSpace(task)

		if clearTaskText == "" {
			fmt.Println("Preencha um texto valido")
			continue
		}

		listTasks = append(listTasks, Task{text: clearTaskText, status: false})
		fmt.Println("Task adicionada com sucesso!")

		break
	}
}

func showTasks() {
	if len(listTasks) == 0 {
		fmt.Println("\nsem tasks no momento")
		return
	}

	fmt.Println("\n===== Lista de tasks =====")

	for i, task := range listTasks {
		sts := "incompleta"

		if task.status {
			sts = "completa"
		}

		fmt.Printf("%d. %s - %s\n", i, task.text, sts)
	}
}

func completeTask() {
	fmt.Print("Informe o indice da task (enter para adicionar): ")

	ix, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Erro na leitura - Não foi possivel completar a task!")
		return
	}

	iNum, err := strconv.ParseUint(strings.TrimSpace(ix), 10, 64)

	if err != nil {
		fmt.Println("Indice invalido - Não foi possivel completar a task!")
		return
	}

	intINum := int(iNum)

	if intINum >= len(listTasks) {
		fmt.Println("Indice invalido - Inexistente!")
		return
	}

	listTasks[intINum].status = true

	fmt.Printf("Task concluida: %s\n", listTasks[intINum].text)
}

func removeTask() {
	fmt.Print("Informe o indice da task (enter para adicionar): ")

	ix, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Erro na leitura - Não foi possivel remover a task!")
		return
	}

	iNum, err := strconv.ParseUint(strings.TrimSpace(ix), 10, 64)

	if err != nil {
		fmt.Println("Indice invalido - Não foi possivel remover a task!")
		return
	}

	intINum := int(iNum)

	if intINum >= len(listTasks) {
		fmt.Println("Indice invalido - Inexistente!")
		return
	}

	text := listTasks[intINum].text

	listTasks = slices.Delete(listTasks, intINum, intINum+1)
	//listTasks = append(listTasks[:intINum], listTasks[intINum+1:]...)

	fmt.Printf("Task removida: %s\n", text)
}

func stopProgram() {
	fmt.Print("Programa Interrompido\n")
	os.Exit(0)
}

func choseAction(option uint64) {
	switch option {
	case 1:
		addTask()
	case 2:
		showTasks()
	case 3:
		completeTask()
	case 4:
		removeTask()
	case 5:
		stopProgram()
	default:
		fmt.Print("Opção invalida!\n")
	}
}

func main() {
	askUserAction()
}

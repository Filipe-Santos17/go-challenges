package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type dtChallenge struct {
	Port            int      `json:"port"`
	Debug           bool     `json:"debug"`
	Allowed_hosts   []string `json:"allowed_hosts"`
	Timeout_seconds int      `json:"timeout_seconds"`
}

func readFile() []byte {
	content, err := os.ReadFile("./challenges/c5/file.json")

	if err != nil {
		panic("Erro ao abrir o arquivo para leitura")
	}

	return content
}

func convertDtFileToJson() dtChallenge {
	contentFile := readFile()

	var dtCh dtChallenge

	err := json.Unmarshal(contentFile, &dtCh)

	if err != nil {
		panic("Erro ao converter o contéudo do arquivo para json")
	}

	return dtCh
}

func main() {
	data := convertDtFileToJson()

	fmt.Println(data)
}

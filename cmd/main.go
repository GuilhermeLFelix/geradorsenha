package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)

	fmt.Println("Informe um tamanho para sua senha: (Valor númerico e maior ou igual a 8)")

	var texto, _ = reader.ReadString('\n')

	tamanho, err := strconv.Atoi(strings.TrimSpace(texto))

	if err != nil {
		fmt.Print("Erro ao ler o valor informado: ", err)
		return
	}

	fmt.Println("Deseja caracteres especiais: (Sim/Nao)")

	texto, _ = reader.ReadString('\n')

	sChar := false

	if strings.TrimSpace(texto) == "Sim" || strings.ToLower(strings.TrimSpace(texto)) == "s" {
		sChar = true
	}

	fmt.Println("Deseja números: (Sim/Nao)")

	texto, _ = reader.ReadString('\n')

	sNum := false

	if strings.TrimSpace(texto) == "Sim" || strings.ToLower(strings.TrimSpace(texto)) == "s" {
		sNum = true
	}

	fmt.Println(GeraSenha(tamanho, sChar, sNum))
}

func GeraSenha(tamanho int, caracterEspecial bool, numeros bool) string {
	if tamanho < 8 {
		tamanho = 8
	}

	senha := ""

	letras := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
		"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T",
		"U", "V", "W", "X", "Y", "Z"}

	for range tamanho {
		senha += letras[rand.Intn(26)]
	}

	if numeros {
		var tamanho1 int

		if tamanho%2 == 0 {
			tamanho1 = tamanho / 2
		} else {
			tamanho1 = (tamanho - 1) / 2
		}

		for range tamanho1 {
			indiceSenha := rand.Intn(tamanho)

			senha = string(senha[:indiceSenha] + strconv.Itoa((rand.Intn(10))) + senha[indiceSenha+1:])
		}
	}

	if caracterEspecial {
		charsEspeciais := []string{"#", "%", "$", "*", "!"}

		indiceSenha := rand.Intn(tamanho)

		senha = string(senha[:indiceSenha] + charsEspeciais[rand.Intn(len(charsEspeciais))] + senha[indiceSenha+1:])
	}

	return senha
}

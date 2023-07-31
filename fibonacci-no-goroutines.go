package main

import "fmt"

/* Atenção para calcular o numero desejado em fibonacci altere a posição as seguintes campos:

- posicao := uint(45)

aperte para iniciar: ctrl+alt+n (no vscode)
aperte para parar: ctrl+alt+m (no vscode)

enjoy it :)

*/

func fibonnaci(posicao uint) uint {

	if posicao <= 1 {
		return posicao
	}

	return fibonnaci(posicao-2) + fibonnaci(posicao-1)
}

func main() {

	posicao := uint(45)
	fmt.Println(fibonnaci(posicao))

}

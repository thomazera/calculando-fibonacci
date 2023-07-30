package main

import "fmt"



/* Atenção para calcular o numero desejado em fibonacci altere a posição as seguintes campos: 

- calculos := make(chan int, "numero desejado")
- resultados := make(chan int, "numero desejado")
- for calculo:= 0; calculo <= 50; calculo++
- for i:= 0; i <= "numero desejado"; i++

aperte para iniciar: ctrl+alt+n (no vscode)
aperte para parar: ctrl+alt+m (no vscode)

enjoy it :)

*/

func main() {
	calculos := make(chan int, 50)
	resultados := make(chan int,50)

	go worker(calculos, resultados)

	for calculo:= 0; calculo <= 50; calculo++ {
		calculos <- calculo
	}
	close(calculos)

	for i:= 0; i <= 50; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}



}


// Utilizando a função worker para calcular a fibo com melhor desempenho
func worker(calculos <-chan int, resultados chan<- int ) {
	for numero := range calculos {
		resultados <- fibo(numero)
	}
}


// Criando a função que chama o calculo da fibonacci
func fibo(posicao int) int{
	if posicao <=  1 {
		return posicao
	}

	return fibo(posicao-2) + fibo(posicao-1)

}




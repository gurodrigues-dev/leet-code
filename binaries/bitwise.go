package main

import "fmt"

func main() {
	/*


		Operador	Nome	Descrição	Exemplo
		&	AND bit a bit	Compara cada bit dos dois valores. Retorna 1 somente se ambos forem 1.	a & b
		`	`	OR bit a bit	Compara cada bit. Retorna 1 se qualquer um dos bits for 1.
		^	XOR bit a bit	Retorna 1 se os bits forem diferentes.	a ^ b
		&^	AND NOT (Clear bits)	Limpa os bits definidos no segundo operando.	a &^ b
		<<	Deslocamento à esquerda	Move os bits para a esquerda, multiplicando por 2 a cada passo.	a << 1
		>>	Deslocamento à direita	Move os bits para a direita, dividindo por 2 a cada passo.	a >> 1

	*/

	a := 5
	b := 1

	fmt.Println("a & b =", a&b)
	fmt.Println("a | b =", a|b)
	fmt.Println("a ^ b =", a^b)
	fmt.Println("a &^ b =", a&^b)

	fmt.Println("a << 1 =", a<<1)
	fmt.Println("a >> 1 =", a>>1)

}

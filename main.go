package main

import (
	"fmt"
	"math"
)

func calcularQuantidadeMoedas(moedas []int, valor int) {

	contarMoedas := make(map[int]int)

	i := len(moedas) - 1

	for valor > 0 {

		valorMoeda := moedas[i]

		if valor >= valorMoeda {
			valor -= valorMoeda
			contarMoedas[valorMoeda]++
		}

		if valor < valorMoeda {
			i--
		}

	}

	fmt.Println("Quantidade de moedas utilizadas:")
	for _, moeda := range moedas {
		if contarMoedas[moeda] > 0 {
			fmt.Printf("%d moeda(s) de %d centavos\n", contarMoedas[moeda], moeda)
		}
	}
}

func main() {

	moedasDisponiveis := map[int][]int{
		1: {1, 2, 5, 10, 25, 50, 100},
		2: {1, 5, 10, 20, 50, 100},
		3: {1, 2, 5, 10, 20, 50, 100},
		4: {1, 5, 12, 24, 50, 100},
	}

	var opcao int
	var valor float64

	fmt.Println("Escolha um conjunto de moedas:")
	fmt.Println("1: {1, 2, 5, 10, 25, 50, 100}")
	fmt.Println("2: {1, 5, 10, 20, 50, 100}")
	fmt.Println("3: {1, 2, 5, 10, 20, 50, 100}")
	fmt.Println("4: {1, 5, 12, 24, 50, 100}")
	fmt.Print("Digite o número do conjunto (1 a 4): ")

	_, err := fmt.Scanf("%d", &opcao)
	if err != nil {
		fmt.Println("Erro ao ler o conjunto:", err)
		return
	}

	moedas, existe := moedasDisponiveis[opcao]
	if !existe {
		fmt.Println("Opção inválida.")
		return
	}

	fmt.Print("Digite o valor em reais (ex: 0.95): ")
	_, err = fmt.Scanf("%f", &valor)
	if err != nil {
		fmt.Println("Erro ao ler o valor:", err)
		return
	}

	valorCentavos := int(math.Round(valor * 100))

	calcularQuantidadeMoedas(moedas, valorCentavos)

}

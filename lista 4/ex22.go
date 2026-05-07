package main

import "fmt"

func main() {
	var codigos [10]int
	var saldos [10]float64

	// Cadastro inicial das contas
	for i := 0; i < 10; i++ {
		fmt.Printf("Digite o código da conta %d: ", i+1)
		fmt.Scan(&codigos[i])
		fmt.Printf("Digite o saldo inicial da conta %d: ", i+1)
		fmt.Scan(&saldos[i])
	}

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1 - Efetuar depósito")
		fmt.Println("2 - Efetuar saque")
		fmt.Println("3 - Consultar ativo bancário")
		fmt.Println("4 - Finalizar programa")

		var opcao int
		fmt.Print("Escolha uma opção: ")
		fmt.Scan(&opcao)

		if opcao == 4 {
			fmt.Println("Programa finalizado.")
			break
		}

		switch opcao {
		case 1:
			var codigo int
			var valor float64
			fmt.Print("Digite o código da conta: ")
			fmt.Scan(&codigo)
			fmt.Print("Digite o valor do depósito: ")
			fmt.Scan(&valor)

			encontrado := false
			for i := 0; i < 10; i++ {
				if codigos[i] == codigo {
					saldos[i] += valor
					fmt.Println("Depósito realizado. Novo saldo:", saldos[i])
					encontrado = true
					break
				}
			}
			if !encontrado {
				fmt.Println("Conta não encontrada.")
			}

		case 2:
			var codigo int
			var valor float64
			fmt.Print("Digite o código da conta: ")
			fmt.Scan(&codigo)
			fmt.Print("Digite o valor do saque: ")
			fmt.Scan(&valor)

			encontrado := false
			for i := 0; i < 10; i++ {
				if codigos[i] == codigo {
					if saldos[i] >= valor {
						saldos[i] -= valor
						fmt.Println("Saque realizado. Novo saldo:", saldos[i])
					} else {
						fmt.Println("Saldo insuficiente.")
					}
					encontrado = true
					break
				}
			}
			if !encontrado {
				fmt.Println("Conta não encontrada.")
			}

		case 3:
			soma := 0.0
			for _, saldo := range saldos {
				soma += saldo
			}
			fmt.Println("Ativo bancário total:", soma)

		default:
			fmt.Println("Opção inválida.")
		}
	}
}

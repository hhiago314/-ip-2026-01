package main

import "fmt"

func main() {
	// 0 = livre, 1 = ocupado
	janela := make([]int, 24)
	corredor := make([]int, 24)

	for {
		var escolha string
		fmt.Print("Deseja poltrona na (j)anela ou (c)orredor? (sair para encerrar): ")
		fmt.Scan(&escolha)

		if escolha == "sair" {
			break
		}

		var disponiveis []int
		if escolha == "j" {
			for i, val := range janela {
				if val == 0 {
					disponiveis = append(disponiveis, i)
				}
			}
			if len(disponiveis) == 0 {
				fmt.Println("Não há poltronas disponíveis na janela.")
				continue
			}
			fmt.Println("Poltronas livres na janela:", disponiveis)
			var pos int
			fmt.Print("Escolha a posição: ")
			fmt.Scan(&pos)
			if pos >= 0 && pos < 24 && janela[pos] == 0 {
				janela[pos] = 1
				fmt.Println("Reserva feita na janela posição", pos)
			} else {
				fmt.Println("Poltrona inválida ou já ocupada.")
			}
		} else if escolha == "c" {
			for i, val := range corredor {
				if val == 0 {
					disponiveis = append(disponiveis, i)
				}
			}
			if len(disponiveis) == 0 {
				fmt.Println("Não há poltronas disponíveis no corredor.")
				continue
			}
			fmt.Println("Poltronas livres no corredor:", disponiveis)
			var pos int
			fmt.Print("Escolha a posição: ")
			fmt.Scan(&pos)
			if pos >= 0 && pos < 24 && corredor[pos] == 0 {
				corredor[pos] = 1
				fmt.Println("Reserva feita no corredor posição", pos)
			} else {
				fmt.Println("Poltrona inválida ou já ocupada.")
			}
		} else {
			fmt.Println("Opção inválida.")
		}

		// Verifica se o ônibus está cheio
		cheio := true
		for _, v := range janela {
			if v == 0 {
				cheio = false
				break
			}
		}
		for _, v := range corredor {
			if v == 0 {
				cheio = false
				break
			}
		}
		if cheio {
			fmt.Println("Ônibus completamente cheio!")
			break
		}
	}
}

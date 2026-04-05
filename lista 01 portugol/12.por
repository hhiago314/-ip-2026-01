programa {
  inclua biblioteca Matematica --> mat

  funcao inicio() {
    inteiro horas_totais, grupos_de_3, horas_extras
    real valor_pagar

    leia(horas_totais)

    grupos_de_3 = horas_totais / 3
    
    horas_extras = horas_totais % 3

    valor_pagar = (grupos_de_3 * 10.00) + (horas_extras * 5.00)

    escreva("O VALOR A PAGAR E = ", mat.arredondar(valor_pagar, 2), "\n")
  }
}

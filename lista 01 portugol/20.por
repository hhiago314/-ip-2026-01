programa {
  
  funcao inicio() {
    inteiro horas, minutos, segundos, total_segundos

    leia(horas)
    leia(minutos)
    leia(segundos)

    total_segundos = (horas * 3600) + (minutos * 60) + segundos

    escreva("O TEMPO EM SEGUNDOS E = ", total_segundos, "\n")
  }
}

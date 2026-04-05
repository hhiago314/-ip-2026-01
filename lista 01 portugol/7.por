programa {
  inclua biblioteca Matematica --> mat

  funcao inicio() {
    real fahrenheit, polegadas
    real celsius, milimetros

    leia(fahrenheit)
    leia(polegadas)

    celsius = (5.0 * fahrenheit - 160.0) / 9.0

    milimetros = polegadas * 25.4

    escreva("O VALOR EM CELSIUS = ", mat.arredondar(celsius, 2), "\n")
    escreva("A QUANTIDADE DE CHUVA E = ", mat.arredondar(milimetros, 2), "\n")
  }
}
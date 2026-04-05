programa {
  inclua biblioteca Matematica --> mat

  funcao inicio() {
    inteiro n, i
    real fahrenheit, celsius

    leia(n)

    para (i = 1; i <= n; i++) {
      leia(fahrenheit)

      celsius = (5.0 * (fahrenheit - 32.0)) / 9.0

      real f_formatado = mat.arredondar(fahrenheit, 2)
      real c_formatado = mat.arredondar(celsius, 2)

      escreva(f_formatado, " FAHRENHEIT EQUIVALE A ", c_formatado, " CELSIUS\n")
    }
  }
}
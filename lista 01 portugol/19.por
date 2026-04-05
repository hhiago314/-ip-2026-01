programa {
  inclua biblioteca Matematica --> mat

  funcao inicio() {
    inteiro n, k
    real soma = 0.0

    leia(n)

    se (n <= 1) {
      escreva("Numero invalido!\n")
    }
    senao {
      para (k = 1; k <= n; k++) {
        soma = soma + (1.0 / k)
      }

      escreva(mat.arredondar(soma, 6), "\n")
    }
  }
}
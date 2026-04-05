programa {
  inclua biblioteca Matematica --> mat

  funcao inicio() {
    real nota
    caracter conceito

    leia(nota)

    se (nota >= 9.0 e nota <= 10.0) {
      conceito = 'A'
    }
    senao se (nota >= 7.5 e nota < 9.0) {
      conceito = 'B'
    }
    senao se (nota >= 6.0 e nota < 7.5) {
      conceito = 'C'
    }
    senao {
      conceito = 'D'
    }

    escreva("NOTA = ", mat.arredondar(nota, 1), " CONCEITO = ", conceito, "\n")
  }
}
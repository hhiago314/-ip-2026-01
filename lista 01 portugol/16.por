programa {
  inclua biblioteca Matematica --> mat

  funcao inicio() {
    real salario, salario_reajustado

    leia(salario)

    se (salario <= 300.00) {
      salario_reajustado = salario * 1.50
    }
    senao {
      salario_reajustado = salario * 1.30
    }

    escreva("SALARIO COM REAJUSTE = ", mat.arredondar(salario_reajustado, 2), "\n")
  }
}

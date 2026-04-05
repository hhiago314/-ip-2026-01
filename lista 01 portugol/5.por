programa {
  funcao inicio() {

    inteiro conta
    real consumo
    cadeia tipo
    real valor

    leia(conta)
    leia(consumo)
    leia(tipo)

    // cálculo
    se (tipo == "R") {
      valor = 5 + consumo * 0.05
    } senao {
      se (tipo == "C") {
        valor = 500 + consumo * 0.25
      } senao {
        valor = 800 + consumo * 0.04
      }
    }

    escreva("CONTA = ", conta, "\n")
    escreva("VALOR DA CONTA = ", valor, "\n")

  }
}
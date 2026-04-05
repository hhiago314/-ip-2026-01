programa {
  funcao inicio() {

    inteiro n1, n2, n3
    inteiro numero, quadrado

    leia(n1)
    leia(n2)
    leia(n3)

    
    se (n1 < 0 ou n1 > 9 ou n2 < 0 ou n2 > 9 ou n3 < 0 ou n3 > 9) {
      escreva("DIGITO INVALIDO\n")
    } senao {

      numero = n1 * 100 + n2 * 10 + n3
      quadrado = numero * numero

      escreva(numero, ", ", quadrado, "\n")
    }

  }
}
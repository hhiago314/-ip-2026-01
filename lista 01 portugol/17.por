programa {
  
  funcao inicio() {
    inteiro x, y, i, numero_atual

    leia(x)
    leia(y)

    se (x % 2 == 0) {
      
      numero_atual = x
      
      para (i = 1; i <= y; i++) {
        escreva(numero_atual, " ")
        
        numero_atual = numero_atual + 2
      }
      
      escreva("\n")
      
    }
    senao {
      escreva("O PRIMEIRO NUMERO NAO E PAR\n")
    }
  }
}

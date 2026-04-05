programa {
  funcao inicio() {
    real media, n1, n2, n3
    escreva("Digite a primeira nota: ")
    leia(n1)
     escreva("Digite a segunda nota: ")
     leia(n2)
      escreva("Digite a terceira nota: ")
      leia(n3)
     media = n1 + n2 + n3 /3
     escreva("\n MEDIA = ", media)
       se (media >= 6) {
      escreva("\nAPROVADO\n")
    } senao {
      escreva("\nREPROVADO\n")
    }
    
  }
}

programa {
  inclua biblioteca Matematica --> mat

  funcao inicio() {
    inteiro n, i, quadrado

    leia(n)

    para (i = 1; i <= n; i++) {
      
      se (i % 2 == 0) {
        quadrado = i * i
        
        escreva(i, "^2 = ", quadrado, "\n")
      }
    }
  }
}

programa {
  
  funcao inicio() {
    inteiro a1, razao, n
    inteiro soma = 0
    inteiro termo_atual
    inteiro i

    leia(a1)
    leia(razao)
    leia(n)

    termo_atual = a1

    para (i = 1; i <= n; i++) {
      soma = soma + termo_atual
      
      termo_atual = termo_atual + razao
    }

    escreva(soma, "\n")
  }
}
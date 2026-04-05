programa {
  inclua biblioteca Matematica --> mat

  funcao inicio() {
    real n1, n2, n3, media

    leia(n1)
    leia(n2)
    leia(n3)

    media = (n1 + n2 + n3) / 3.0

    escreva("MEDIA = ", mat.arredondar(media, 2), "\n")

    se (media >= 6.0) {
      escreva("APROVADO\n")
    }
    senao {
      escreva("REPROVADO\n")
    }
  }
}
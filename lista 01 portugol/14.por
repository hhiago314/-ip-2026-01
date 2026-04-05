programa {
  inclua biblioteca Matematica --> mat

  funcao inicio() {
    real h, a, area_base, volume
    
    leia(h)
    leia(a)

    area_base = (3 * mat.potencia(a, 2.0) * mat.raiz(3.0, 2.0)) / 2.0

    volume = (1.0 / 3.0) * area_base * h

    escreva("O VOLUME DA PIRAMIDE E = ", mat.arredondar(volume, 2), " METROS CUBICOS\n")
  }
}
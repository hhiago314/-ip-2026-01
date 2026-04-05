programa {
  inclua biblioteca Matematica --> mat

  funcao inicio() {
    real raio, altura, custo
    real area_circulo, area_lateral, area_total
    
    const real PI = 3.14159

    leia(raio)
    leia(altura)

    area_circulo = PI * (raio * raio)
    
    area_lateral = 2 * PI * raio * altura
    
    area_total = (2 * area_circulo) + area_lateral

    custo = area_total * 100.00

    escreva("O VALOR DO CUSTO E = ", mat.arredondar(custo, 2), "\n")
  }
}

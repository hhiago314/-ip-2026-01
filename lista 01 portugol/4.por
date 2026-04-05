programa {
  funcao inicio() {

    real salarioMinimo
    real consumo
    real custoPorKw
    real total
    real totalDesconto

    leia(salarioMinimo)
    leia(consumo)

    custoPorKw = (salarioMinimo * 0.7) / 100

    total = custoPorKw * consumo

    totalDesconto = total * 0.9

    escreva("Custo por kW: R$ ", custoPorKw, "\n")
    escreva("Custo do consumo: R$ ", total, "\n")
    escreva("Custo com desconto: R$ ", totalDesconto, "\n")

  }
}
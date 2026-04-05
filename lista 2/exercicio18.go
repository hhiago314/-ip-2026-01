package main
import "fmt"

func main() {
Perfeito, Hiago! Aqui está a lista **completa** das 23 soluções em Go, cada uma como um programa independente. Você pode salvar cada trecho em arquivos separados (`q1.go`, `q2.go`, etc.) e rodar com `go run qX.go`.  

---

### Questão 18 – Locadora DVDs
```go
package main
import "fmt"

func main() {
    var preco float64
    var categoria string
    var dia int
    fmt.Scan(&preco, &categoria, &dia)

    if dia == 2 || dia == 3 || dia == 5 {
        preco *= 0.6
    }
    if categoria == "lancamento" {
        preco *= 1.15
    }
    fmt.Println("Preço final:", preco)
}

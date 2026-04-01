package main
import f "fmt"
func main() {
	a, b, c := 0, 0, 0
	f.Print("digite o primeiro numero: ")
	f.Scan(&a)
	f.Print("digite o segundo numero: ")
	f.Scan(&b)
	f.Print("digite o terceiro numero: ")
	f.Scan(&c)
	maior:=dmaior(a, b, c)
	f.Printf("%d é o maior numero", maior)
}
func dmaior(a, b, c int) int{
	maior:=a
	if b > maior{
		maior=b
	}
	if c > maior{
		maior=c
	}
	return maior
}
package main
import f "fmt"
func main(){
	n:= 0
	f.Print("digite um numero: ")
	f.Scan(&n)
	fatorial:=ffatorial(n)
	f.Printf("o fatorial de %d é %d",n, fatorial)
}
func ffatorial(n int) int{
	resultado:=1
	for i := 1; i <= n; i++ {
		resultado*=i
	}
	return resultado
}
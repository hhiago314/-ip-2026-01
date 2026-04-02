package main
import f "fmt"
func main(){
	x, y, z := 0, 0, 0
	f.Print("digite o primeiro numero: ")
	f.Scan(&x)
	f.Print("digite o segundo numero: ")
	f.Scan(&y)
	f.Print("digite o terceiro numero: ")
	f.Scan(&z)
	media:=fmedia(x, y, z)
	f.Printf("A media dos numeros é %d", media)

}
func fmedia(x, y, z int) int{
	return (x+y+z)/3
}
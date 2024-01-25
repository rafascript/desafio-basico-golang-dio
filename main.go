package main

import "fmt"

const ebulicaoK float64 = 373.0

//função principal
func main() {

	var tempK float64 = ebulicaoK     //variável do valor da temperatura em graus fahrenheit
	var tempC float64 = (tempK - 273) //F to C
	//exibir resultado em tela
	fmt.Println("A temperatura de ebulição da água em K° : ", tempK)
	fmt.Println("A temperatura de ebulição da água em C° : ", tempC)

}

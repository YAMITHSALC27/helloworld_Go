package main

import "fmt"

func imprimir() {
	fmt.Println("texto desde la funcion imprimir")
}

func hola_string(s string) string {
	return "Hola " + s
}

func sumar(a int, b int) int {
	return a + b
}

func comparar_bool() {
	var a bool
	b := true
	a = false
	if a == b {
		fmt.Println("Los valores son iguales")
	} else {
		fmt.Println("Los valores son diferentes")
	}
}
func main() {
	//fmt.Println("texto desde la funcion main")
	//imprimir()
	//fmt.Println(hola_string("yamith"))
	//fmt.Println(sumar(3, 5))
	comparar_bool()
}

package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func imprimir() {
	fmt.Println("texto desde la funcion imprimir")
}

func hola_string(s string) string {
	return "Hola " + s
}

func sumar(a int, b int) int {
	return a + b
}

// func comparar_bool() {
// 	var a bool
// 	b := true
// 	a = false
// 	if a == b {
// 		fmt.Println("Los valores son iguales")
// 	} else {
// 		fmt.Println("Los valores son diferentes")
// 	}
// }

func arreglo() {
	var aprendices [5]string

	aprendices[0] = "Juan"
	aprendices[1] = "Maria"
	aprendices[2] = "Pedro"
	aprendices[3] = "Ana"
	aprendices[4] = "frank"
	fmt.Println(aprendices)
}

func tipos_de_datos() {
	var texto string = "fabian"
	var numero int = 1000
	var decimal float64 = 0.0001

	fmt.Println(reflect.TypeOf(texto))
	fmt.Println(reflect.TypeOf(numero))
	fmt.Println(reflect.TypeOf(decimal))
}

func converstringboolean() {
	var palabra string = "true"

	boolean, err := strconv.ParseBool(palabra)
	fmt.Println(boolean, reflect.TypeOf(boolean))
	fmt.Println("este es el error: ", err)
}

func convertbooltostring() {
	var palabra_bool bool = true

	strbool := strconv.FormatBool(palabra_bool)
	fmt.Println(strbool, reflect.TypeOf(strbool))
}

func tipos_variables() {
	var (
		nombre, apellido string = "yamith", "salcedo"
		edad             int    = 18
		pensionado       bool   = false
	)
	fmt.Printf("Nombre: %s\nApellido: %s\nEdad: %d\nPensionado: %t\n", nombre, apellido, edad, pensionado)

}
const pi= 3.1416
func area (radio float64) float64 {
	return pi * radio * radio
}

func valor_por_defecto(){
	var(
		nombre, apellido string
		edad int 
		peso float64
		estudiante bool
	)
	fmt.Printf("Nombre: %s\nApellido: %s\nEdad: %d\nPeso: %.2f\nEstudiante: %t\n", nombre, apellido, edad, peso, estudiante)
}



func main() {
	//fmt.Println("texto desde la funcion main")
	//imprimir()
	//fmt.Println(hola_string("yamith"))
	//fmt.Println(sumar(3, 5))
	//comparar_bool()
	// arreglo()
	// tipos_de_datos()
	//converstringboolean()
	//convertbooltostring()
	//tipos_variables()
	//valor_por_defecto()
	fmt.Printf("El area del circulo con radio 5 es: %.2f\n", area(5))
}

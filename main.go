package main

import (
	"Calculator/CalculatorApp"
	"fmt"
)

func main() {
	//Mensaje de pantalla para probar
	fmt.Println("Hola Mundo")
	fmt.Println("Estoy practicando una calculadora basica")
	opc := 0
	var num1, num2 int
	for {
		//Solicitamos el tipo de operación del usuario
		fmt.Println("Seleccionar opción")
		fmt.Println("1- Suma \n2- Resta \n3- Multiplicación \n4- Division \n5- Cerrar")
		//captura valores digitados en pantalla
		fmt.Scanf("%d\n", &opc)

		if opc == 1 {
			fmt.Println("Ingrese el primer número: ")
			fmt.Scanf("%d\n", &num1)
			fmt.Println("Ingrese el segundo número: ")
			fmt.Scanf("%d\n", &num2)
			CalculatorApp.Sum(num1, num2)
		} else if opc == 2 {
			fmt.Println("Ingrese el primer número: ")
			fmt.Scanf("%d\n", &num1)
			fmt.Println("Ingrese el segundo número: ")
			fmt.Scanf("%d\n", &num2)
			CalculatorApp.Subtract(num1, num2)
		} else if opc == 3 {
			fmt.Println("Ingrese el primer número: ")
			fmt.Scanf("%d\n", &num1)
			fmt.Println("Ingrese el segundo número: ")
			fmt.Scanf("%d\n", &num2)
			CalculatorApp.Multiplication(num1, num2)
		} else if opc == 4 {
			fmt.Println("Ingrese el primer número: ")
			fmt.Scanf("%d\n", &num1)
			fmt.Println("Ingrese el segundo número: ")
			fmt.Scanf("%d\n", &num2)
			CalculatorApp.Division(num1, num2)
		} else if opc == 5 {
			fmt.Println("Cerrando Calculadora")
			break
		} else {
			fmt.Println("Por favor seleccione una opción correcta")
		}
	}

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	// TODO: Instalar e importar la libreria Strcase
)

var options string = `
Seleccione una opción:

1. Convertir a capital case (ex. Capital)
2. Convertir a minúscula
3. Convertir a mayúscula
4. Convertir a Camel case (ex. camelCase)
5. Convertir a Snake case (ex. snake_case)
6. Convertir a Pascal case (ex. PascalCase)
7. Convertir a Kebab case (ex. kebab-case)
8. Convertir tu edad en años a días, semanas y meses
9. Convertir tu peso de kilogramos a libras
10. Convertir tu estatura de metros a pulgadas
11. Calcular tu índice de masa corporal imc
12. Mostrar mi resumen
`

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	fmt.Println(options)
	// Declaramos una variable input recibiendo lo que el usuario ingrese hasta que le de enter.
	input, _ := reader.ReadString('\n')
	input = input[:len(input)-1]

	// Declaramos una variable option y una variable err,
	// y con la función Atoi recibimos el string input convirtiendo el texto en un entero
	option, err := strconv.Atoi(input)
	if err != nil {
		panic("Debes escribir un número entre el 1 y el 12")
	}
	// if option < 1 && option > 12
	if option >= 1 && option <= 12 {
		if option == 1 {
			fmt.Println("Introduzca un texto")
			in, _ := reader.ReadString('\n')
			fmt.Println(strings.Title(strings.ToLower(in)))
		} else if option == 2 {
			fmt.Println("Introduzca un texto")
			in, _ := reader.ReadString('\n')
			fmt.Println(strings.ToLower(in))
		} else if option == 3 {
			fmt.Println("Introduzca un texto")
			in, _ := reader.ReadString('\n')
			fmt.Println(strings.ToUpper(in))
		} else if option == 4 {
			fmt.Println("Convertir a Camel case (ex. camelCase)")
			// TODO: Indicarle al usuario a traves de un Print que debe ingresar texto
			// TODO: Uso la función que me permite tomar texto de la terminal y lo guardo en una variable
			// TODO: Usar la libreria Strcase para convertir el texto, dentro de un Print
		} else if option == 5 {
			fmt.Println("Convertir a Snake case (ex. snake_case)")
		} else if option == 6 {
			fmt.Println("Convertir a Pascal case (ex. PascalCase)")
		} else if option == 7 {
			fmt.Println("Convertir a Kebab case (ex. kebab-case)")
		} else if option == 8 {
			fmt.Println("Convertir tu edad en años a días, semanas y meses")
		} else if option == 9 {
			fmt.Println("Convertir tu peso de kilogramos a libras")
		} else if option == 10 {
			fmt.Println("Convertir tu estatura de metros a pulgadas")
		} else if option == 11 {
			fmt.Println("Calcular tu índice de masa corporal imc")
		} else if option == 12 {
			fmt.Println("Mostrar mi resumen")
		}
	}

	// ¡Hola! <name> <lastname>. Tu edad en días es <age (days)>, en semanas es de <age (weeks)> y en meses es de <age (months)>
	/* fmt.Printf("¡Hola!, %s %s. Tu edad en días es %.2f, en semanas es de %.2f y en meses es de %.2f \n",
	name, lastName, age * 365, age * 52, age * 12)*/
}

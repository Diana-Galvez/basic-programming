package main

import (
	"bufio"
	"fmt"
	"os"
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
	var input string = ""
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	fmt.Println(options)
	input, _ = reader.ReadString('\n')
	input = input[:len(input)-1]
  fmt.Println(input)
	/*var age float64 = 0
	var name string = ""
	var lastName string = ""

  // indicamos al usuario lo que debe escribir y al final dar a enter, para que
	// quede guardado en la variable name
	fmt.Print("escribe tu nombre: ")
	// Una vez finalice el mensaje darle a enter
	name, _ = reader.ReadString('\n')
	name = name[:len(name)-1]

	// Te pide escribir tu apellido
	fmt.Print("escribe tu apellido: ")
	// Una vez hecho darle enter para que quede guardo en la variable lastName
  lastName, _ = reader.ReadString('\n')
	lastName = lastName[:len(lastName)-1]



	// Declaramos una variable age y una variable err,
	// y con ParseFloat recibimos el string "texto" y también recibe el grado de precisión
	// convirtiéndolo en decimal.
	age, err := strconv.ParseFloat(input, 2)
	if err != nil {
		panic("Debes escribir un número decimal (ex. 32.6)")
	}

// ¡Hola! <name> <lastname>. Tu edad en días es <age (days)>, en semanas es de <age (weeks)> y en meses es de <age (months)>
  fmt.Printf("¡Hola!, %s %s. Tu edad en días es %.2f, en semanas es de %.2f y en meses es de %.2f \n",
		 name, lastName, age * 365, age * 52, age * 12)*/
}

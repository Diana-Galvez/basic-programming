package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// var age int = 34
	// var name string = "diana patricia"
	// var lastName string = "galvez gaviria"
	// var numberPets int = 1
	// var height float32 = 1.68
	// var isWoman bool = true

	/*fmt.Println(
			"hola mi nombre es", name, lastName,
			", tengo", age, "años, tengo", numberPets,
			"mascota mido", height,
			"metros y soy mujer:", isWoman,
	 ) */

  /*fmt.Printf(
		"hola mi nombre es %s %s, tengo %3d años, tengo %d mascota, mido %.2f metros y soy mujer:, %t\n",
		 name, lastName, age, numberPets, height, isWoman,
	 )*/

	// fmt.Printf("Yo mido %.3f pulgadas ", height/0.0256)

	var age int = 0
	var name string = ""
	var lastName string = ""
	// var numberPets int = 0
	// var height float32 = 0
	// var isWoman bool = false

  // indicamos al usuario lo que debe escribir y al final dar a enter, para que
	// quede guardado en la variable name
	fmt.Print("escribe tu nombre: ")
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	// Una vez finalice el mensaje darle a enter
	name, _ = reader.ReadString('\n')
	name = name[:len(name)-1]

	// Te pide escribir tu apellido
	fmt.Print("escribe tu apellido: ")
	// Una vez hecho darle enter para que quede guardo en la variable lastName
  lastName, _ = reader.ReadString('\n')
	lastName = lastName[:len(lastName)-1]

  var input string = ""
	fmt.Print("escribe tu edad en años")
	input, _ = reader.ReadString('\n')
	input = input[:len(input)-1]
	age, err = strconv.Atoi(input)
	if err != nil {
		panic("Debes escribir un número entero")
	}

  fmt.Printf("¡Hola!, %s %s. Tu edad en días es\n", lastName, name)
}

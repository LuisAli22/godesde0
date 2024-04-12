package main

import (
	"fmt"
	"runtime"

	"github.com/LuisAli22/godesde0/ejercicios"
)

func main() {
	//variables.MuestroEnteros()
	/*estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)*/
	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Esto no es windows, es ", os)
	} else {
		fmt.Println("Esto es windows")

	}
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("%s \n", os)

	}
	number, message := ejercicios.GetNumberInfo("50")
	fmt.Printf("Number: %d \n", number)
	fmt.Printf("Message: %s \n", message)
}

package main

import (
	//	"fmt"

	"fmt"

	"github.com/LuisAli22/godesde0/interfaces/goroutines"
)

func main() {
	//variables.MuestroEnteros()
	/*estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)*/
	/*if os := runtime.GOOS; os == "Linux." || os == "OS X." {
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
	fmt.Printf("Message: %s \n", message)*/
	//teclado.IngresoNumeros()
	//iteraciones.Iterar()
	//fmt.Println(ejercicios.ObtenerNumeroYMostrarTabla())
	//files.SumaTabla()
	//files.LeoArchivo()
	//funciones.Calculos()
	//funciones.LlamarClosure()
	//funciones.Exponencial(2)
	//arreglos_slices.Capacidad()
	//mapas.MostrarMapas()
	//users.AltaUsuario()
	/*Pedro := new(modelos.Hombre)
	e.HumanosRespirando(Pedro)
	Maria := new(modelos.Mujer)
	e.HumanosRespirando(Maria)*/
	//d.VemosDefer()
	//d.EjemploPanic()
	go goroutines.MiNombreLentooo("Luis Ali")
	fmt.Println("Estoy aqui")
	var x string
	fmt.Scanln(&x)
}

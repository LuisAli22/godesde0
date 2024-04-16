package defer_panic

import (
	"fmt"
	"log"
)

func VemosDefer() {
	fmt.Println("Ese es un primer mensaje")
	defer fmt.Println("Ese es el mensaje final")
	fmt.Println("Ese es un segundo mensaje")
}
func EjemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrio un error que generò el panic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("Se encontro el valor 1")
	}
}

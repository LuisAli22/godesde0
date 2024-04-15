package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func mostrarTabla(numeroDeTabla int) {
	for i := 1; i <= 10; i++ {
		fmt.Println(numeroDeTabla, " X ", i, "= ", numeroDeTabla*i)
	}
}
func ObtenerNumeroYMostrarTabla() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Ingrese nùmero de tabla: ")
		if scanner.Scan() {
			numeroDeTabla, err := strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			}
			mostrarTabla(numeroDeTabla)
			break
		}

	}
}

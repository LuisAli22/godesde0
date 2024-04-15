package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func mostrarTabla(numeroDeTabla int) string {
	var texto string
	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintln(numeroDeTabla, " X ", i, "= ", numeroDeTabla*i)
	}
	return texto
}
func ObtenerNumeroYMostrarTabla() string {
	scanner := bufio.NewScanner(os.Stdin)
	var tabla string
	for {
		fmt.Println("Ingrese nùmero de tabla: ")
		if scanner.Scan() {
			numeroDeTabla, err := strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			}
			tabla = mostrarTabla(numeroDeTabla)
			break
		}

	}
	return tabla
}

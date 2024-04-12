package ejercicios

import "strconv"

func GetNumberInfo(strnum string) (int, string) {
	number, err := strconv.Atoi(strnum)
	if err != nil {
		return 0, "Hubo un error: " + err.Error()
	}
	if number > 100 {
		return number, "Es mayor a 100"
	} else {
		return number, "Es menor a 100"
	}
}

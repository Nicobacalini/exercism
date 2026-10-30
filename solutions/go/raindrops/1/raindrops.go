package raindrops

import "strconv"

func Convert(number int) string {
	resultado := ""

	if number%3 == 0 {
		resultado += "Pling"
	}
	if number%5 == 0 {
		resultado += "Plang"
	}
	if number%7 == 0 {
		resultado += "Plong"
	}
	if resultado == "" {
		return strconv.Itoa(number)
	}

	return resultado
}

package isbnverifier

import (
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {
	textoLimpio := strings.ReplaceAll(isbn, "-", "")

	if len(textoLimpio) != 10 {
		return false
	}

	resultado := 0
	multiplicador := 10

	for i := 0; i < 10; i++ {
		caracter := textoLimpio[i]
		var valor int

		if caracter == 'X' || caracter == 'x' {
			if i != 9 {
				return false
			}
			valor = 10
		} else {
			var err error
			valor, err = strconv.Atoi(string(caracter))
			if err != nil {
				return false
			}
		}

		resultado += valor * multiplicador
		multiplicador--
	}

	return resultado%11 == 0
}

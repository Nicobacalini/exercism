package phonenumber

import (
	"errors"
	"fmt"
	"strings"
)

func Number(phoneNumber string) (string, error) {
	var digitos strings.Builder

	for _, r := range phoneNumber {
		if r >= '0' && r <= '9' {
			digitos.WriteRune(r)
		} else if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			return "", errors.New("los numeros de telefono no pueden contener letras")
		} else if strings.ContainsRune("@:!", r) { 
			return "", errors.New("caracteres especiales no permitidos")
		}
	}
	numClean := digitos.String()
	largo := len(numClean)

	if largo < 10 {
		return "", errors.New("numero demasiado corto")
	}
	if largo > 11 {
		return "", errors.New("numero demasiado largo")
	}
	if largo == 11 {
		if numClean[0] != '1' {
			return "", errors.New("el codigo de pais de 11 digitos debe comenzar con 1")
		}
		numClean = numClean[1:]
	}

	if numClean[0] == '0' {
		return "", errors.New("el codigo de area no puede comenzar con cero")
	}
	if numClean[0] == '1' {
		return "", errors.New("el codigo de area no puede comenzar con uno")
	}

	if numClean[3] == '0' {
		return "", errors.New("el codigo de la central no puede comenzar con cero")
	}
	if numClean[3] == '1' {
		return "", errors.New("el codigo de la central no puede comenzar con uno")
	}

	return numClean, nil
}

func AreaCode(phoneNumber string) (string, error) {
	num, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return num[0:3], nil
}

func Format(phoneNumber string) (string, error) {
	num, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", num[0:3], num[3:6], num[6:10]), nil
}

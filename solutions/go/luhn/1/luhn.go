package luhn

import "strings"

func tieneNoNumeros(s string) bool {
    for _, c := range s {
        if c < '0' || c > '9' {
            return true
        }
    }
    return false
}

func Valid(id string) bool {

    id = strings.ReplaceAll(id, " ", "")
	if len(id) < 2 ||  tieneNoNumeros(id) {
        return false
    }
    
	sumaTotal := 0
	deboDuplicar := false
    for i:=len(id) -1 ; i >= 0; i-- {
        digito := int(id[i] - '0')
		if deboDuplicar {
            digito *= 2
            if digito > 9{
                digito -= 9
            }
        }
        sumaTotal += digito
        deboDuplicar = !deboDuplicar
    }
    return sumaTotal%10 == 0
}

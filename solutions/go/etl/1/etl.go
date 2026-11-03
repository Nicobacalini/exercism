package etl
import "strings"
func Transform(in map[int][]string) map[string]int {
	resultado := map[string]int{}

    for puntaje , letras := range in {
        for _, letra := range letras{
            resultado[strings.ToLower(letra)] = puntaje
        }
    }
    return resultado
}

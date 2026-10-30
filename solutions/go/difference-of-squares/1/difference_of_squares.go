package differenceofsquares

func SquareOfSum(n int) int {
	contador := 0
    for n > 0{
        contador = contador + n
        n--
    }
    return contador * contador
}

func SumOfSquares(n int) int {
	contador := 0
    for n > 0 {
        contador = contador + (n *n)
        n--
    }
    return contador
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

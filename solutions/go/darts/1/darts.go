package darts
import "math"

func Score(x, y float64) int {
	distancia := math.Sqrt((x*x)+(y*y))

    if distancia <= 1{
        return 10
    } else if distancia <= 5{
        return 5
    } else if distancia <=10{
        return 1
    }else{
        return 0
    }
}

// Package weather proporciona herramientas para el monitoreo meteorológico y
// la generación de previsiones climáticas dentro del territorio de Goblinocus.
// Permite registrar las condiciones actuales y predecir el clima futuro.
package weather

var (
	// CurrentCondition almacena una descripción textual del estado del clima actual (ej. "Soleado", "Lluvia de fuego").
	// Su uso está destinado únicamente a la lectura y actualización de reportes en tiempo real.
	CurrentCondition string

	// CurrentLocation almacena el nombre del área geográfica o región que se está evaluando actualmente.
	// Modificar esta variable permite cambiar el enfoque del reporte meteorológico a una nueva ubicación.
	CurrentLocation string
)

// Forecast toma una ubicación y una condición climática actual para generar
// un resumen detallado de la previsión del tiempo que se espera en esa zona.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

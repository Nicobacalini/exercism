package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    if len(birdsPerDay) == 0 {
        return 0
    }
	resultado := 0
    for i:= 0; i< len(birdsPerDay); i++ {
        resultado = resultado + birdsPerDay[i]
    }
    return resultado
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	const daysInWeek = 7
	startIndex := (week - 1) * daysInWeek
	endIndex := startIndex + daysInWeek
	totalBirds := 0
    
	for _, count := range birdsPerDay[startIndex:endIndex] {
		totalBirds += count
	}
	return totalBirds
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	if len(birdsPerDay) < 0 {
        return birdsPerDay
    }

    resultadonuevo:= birdsPerDay
    for i := 0; i < len(birdsPerDay); i = i +2 {
        resultadonuevo[i] = resultadonuevo[i]+ 1
    }

    return resultadonuevo
}

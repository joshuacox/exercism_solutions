package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	//return int((float64(productionRate) / 60) * (successRate / 100) + 0.5)
	return int((float64(productionRate) / 60) * (successRate / 100))
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var totalCost uint = 0
	for carsCount > 9 {
		carsCount = carsCount - 10
		totalCost += 95000
	}
	//totalCost = 95000 * uint(carsCount / 10)
	//carsCount = carsCount - (carsCount / 10) * 10
	totalCost = totalCost + uint(carsCount * 10000)
	return totalCost
}

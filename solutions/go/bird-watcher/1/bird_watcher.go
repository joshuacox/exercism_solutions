package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	count := 0
	for i := 0; i < len(birdsPerDay); i++ {
		count += birdsPerDay[i]
	}
	return count
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	bird_count := 0
	start_count := (week - 1) * 7
	end_count := start_count + 7
	for i := start_count; i < end_count; i++ {
		if i < len(birdsPerDay) {
		  bird_count += birdsPerDay[i]
		}
	}
	return bird_count

}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i+=2 {
		birdsPerDay[i]++
	}
	return birdsPerDay
}

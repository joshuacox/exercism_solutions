package speed

type Car struct {
	speed int
	battery int
	batteryDrain int
	distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed: speed,
		batteryDrain: batteryDrain,
		battery: 100,
	}
}

type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	//track_distance := 500
	//drain_delta := car.batteryDrain * track_distance
	//drain_delta := car.batteryDrain * car.speed
	//
	drain_delta := car.batteryDrain
	if drain_delta <= car.battery {
		car.battery = car.battery - drain_delta
		//car.distance = car.distance +  track_distance
		car.distance = car.distance +  car.speed
	}
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
		return (car.battery/car.batteryDrain)*car.speed >= track.distance
}

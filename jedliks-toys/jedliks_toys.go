package jedlik

import "fmt"

// Drive drives car one time. IF there is not enough battery to drive,
// the car will not move.
func (car *Car) Drive() {
	if car.battery-car.batteryDrain >= 0 {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
}

// DisplayDistance returns distance driven.
func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

// DisplayBattery returns battery percentage.
func (car Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

// Check if car could finish a race.
func (car Car) CanFinish(trackDistance int) bool {
	for car.battery > 0 {
		car.Drive()
	}

	return car.distance >= trackDistance
}

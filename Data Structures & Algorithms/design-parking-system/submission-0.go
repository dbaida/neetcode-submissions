const (
	carSizeBig int = iota + 1
	carSizeMedium
	carSizeSmall
)

type ParkingSystem struct {
	space map[int]int
}
func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		space: map[int]int{
			carSizeBig: big,
			carSizeMedium: medium,
			carSizeSmall: small,
		},
	}
}
func (this *ParkingSystem) AddCar(carType int) bool {
	spaceLeft := this.space[carType]
	if spaceLeft == 0 {
		return false
	}
	this.space[carType]--

	return true
}

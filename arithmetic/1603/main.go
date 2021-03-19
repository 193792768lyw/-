package main

//1603. 设计停车系统
func main() {

}

type ParkingSystem struct {
	count [3]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{count: [3]int{big, medium, small}}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.count[carType-1] > 0 {
		this.count[carType-1] -= 1
		return true
	}
	return false
}

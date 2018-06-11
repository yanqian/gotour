package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934


type car struct{
	gas_pedal uint16
	break_pedal uint16
	sterring_wheel int16
	top_speed_kmh float64
}

func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax)
}

func (c car) mph() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax/kmh_multiple)
}

func main() {
	a_car := car{gas_pedal: 65000,
				break_pedal: 0,
				sterring_wheel: 7855,
				top_speed_kmh: 255.56}
	b_car := car{65000, 0, 39, 500}
	fmt.Println(a_car.top_speed_kmh)
	fmt.Println(b_car.kmh())
	fmt.Println(b_car.mph())

}
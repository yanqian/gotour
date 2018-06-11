package main

import "fmt"
type car struct{
	gas_pedal uint16
	break_pedal uint16
	sterring_wheel int16
	top_speed_kmh float64
}

func main() {
	a_car := car{gas_pedal: 233,
				break_pedal: 0,
				sterring_wheel: 7855,
				top_speed_kmh: 255.56}
	b_car := car{100, 0, 39, 500}
	fmt.Println(a_car.top_speed_kmh)
	fmt.Println(b_car.gas_pedal)

}
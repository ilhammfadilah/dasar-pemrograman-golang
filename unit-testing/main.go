package main

import "math"

type Kubus struct {
	sisi float64
}

func (k Kubus) Volume() float64 {
	return math.Pow(k.sisi, 3)
}

func (k Kubus) Luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}

func (k Kubus) Keliling() float64 {
	return k.sisi * 12
}

func main() {}

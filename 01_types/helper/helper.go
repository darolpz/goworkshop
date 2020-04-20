package helper

import "math"

func sinHelper(number float64) float64 {
	return math.Sin(number)
}

func cosHelper(number float64) float64 {
	return math.Cos(number)
}

func SinCosHelper(number float64) ( sin,  cos float64){
		sin = sinHelper(number)
		cos = cosHelper(number)
		return
}
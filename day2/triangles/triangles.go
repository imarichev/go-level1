package triangles

import (
	"errors"
	"math"
)

func IsDotsCorrect(a float64, b float64, c float64) (error){
	if (a + b <= c || a + c <= b || b + c <= a) {
		return errors.New("incorrect dots")
	}
	return nil
}

func Square(a float64, b float64, c float64) (float64) {
	p :=(a + b + c) / 2.0
	return math.Sqrt(p * (p - a) * (p - b) * (p - c))
}

func GetSides(x1 int, y1 int, x2 int, y2 int, x3 int, y3 int)(float64, float64, float64){
	a := math.Sqrt(math.Pow(float64(x2 - x1), 2) + math.Pow(float64(y2 - y1), 2))
	b := math.Sqrt(math.Pow(float64(x3 - x2), 2) + math.Pow(float64(y3 - y2), 2))
	c := math.Sqrt(math.Pow(float64(x1 - x3), 2) + math.Pow(float64(y1 - y3), 2))
	return a, b, c
}

func IsRight(a float64, b float64, c float64)(bool){
	if a * a + b * b - c * c < 0.01|| a * a + c * c - b * b < 0.01 || b * b + c * c - a * a < 0.01{
		return true
	}else{
		return false
	}
}
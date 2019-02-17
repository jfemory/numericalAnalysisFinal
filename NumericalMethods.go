package main

import (
	"fmt"
)

//test
//Initial conditions
var acc = 0.01
var a = 0.0
var b = 180.0
var N = 720
var alpha = 100.0
var h = (b - a) / float64(N)

func main() {
	t := a
	w := alpha
	for i := 1; i < N; i++ {
		k1 := f(t, w)
		k2 := f(t+(h/2), w+(k1/2))
		k3 := f(t+(h/2), w+(k2/2))
		k4 := f(t+h, w+k3)
		w = w + (k1+(2*k2)+(2*k3)+k4)/6
		t = a + float64(i)*h
		fmt.Println(t, w)
	}
}

//func f is the function being evaluated by the 4RK method.
func f(t float64, w float64) float64 {
	return h * (-1 - 0.02*w + (acc * t))
}

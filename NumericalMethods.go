package main

import (
	"fmt"
)

type Entry struct {
	t float64
	w float64
}

//Initial conditions
var acc = 0.01
var a = 0.0
var b = 90.0
var N = 180
var alpha = 100.0
var h = (b - a) / float64(N)

func main() {
	outputRK := rungeKutta(a, alpha)
	fmt.Println(outputRK)
	outputAB := adamsBashforth(outputRK)
	fmt.Println("Runge-Kutta")
	printIt(outputRK)
	fmt.Println("Adams-Bashforth")
	printIt(outputAB)
}

//printIt takes an array of entries, and prints them to stdo
//in a format that can easily be copied out
func printIt(in []Entry) {
	for i := 0; i < len(in); i++ {
		fmt.Println(in[i].t, in[i].w)
	}
}

//adamsBashforth takes an array of t and w values, takes the
//first four values, and returns an array of AB values
func adamsBashforth(in []Entry) []Entry {
	output := make([]Entry, 4)
	for i := 0; i < 4; i++ {
		output[i] = in[i]
	}
	for j := 4; j < N; j++ {
		var z Entry
		w0, w1, w2, w3 := in[j-4], in[j-3], in[j-2], in[j-1]
		//Update t
		z.t = a + float64(j)*h
		//Predict new w value
		z.w = w3.w + ((h * ((55 * f(w3.t, w3.w)) -
			(59 * f(w2.t, w2.w)) + (37 * f(w1.t, w1.w)) -
			(9 * f(w0.t, w0.w)))) / 24)
		//Correct new w value
		z.w = w3.w + ((h * ((9 * f(z.t, z.w)) +
			(19 * f(w3.t, w3.w)) - (5 * f(w2.t, w2.w)) +
			f(w1.t, w1.w))) / 24)
		output = append(output, z)
	}
	return output
}

//rungeKutta takes an initial t and w and returns a slice based on
//the global step size of t and w
func rungeKutta(t0 float64, w0 float64) []Entry {
	output := make([]Entry, 0)
	t := t0
	w := w0
	output = append(output, Entry{t, w})
	for i := 1; i < N; i++ {
		k1 := h * f(t, w)
		k2 := h * f(t+(h/2), w+(k1/2))
		k3 := h * f(t+(h/2), w+(k2/2))
		k4 := h * f(t+h, w+k3)
		w = w + (k1+(2*k2)+(2*k3)+k4)/6
		t = a + float64(i)*h
		output = append(output, Entry{t, w})
	}
	return output
}

//func f is the function being evaluated by the numerical
//methods considered.
func f(t float64, w float64) float64 {
	return (-1 - 0.02*w + (acc * t))
}

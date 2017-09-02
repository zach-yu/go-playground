package main

import (
"fmt"
"math"
)

type Func func(float64) float64

//https://en.wikipedia.org/wiki/Closure_(computer_programming)
// function as an arg & return value
// f & dx outlived the scope of derivative()
func derivative(f Func, dx float64)  Func{
	return func (x float64) float64{
		return (f(x+dx) - f(x))/dx
	}
}


func main() {
	deriv1 := derivative(func(x float64) float64{return x*x*x}, 0.00001) //closure
	fmt.Printf("%f\n", deriv1(1))
	fmt.Printf("%f\n", deriv1(2))
	deriv2 := derivative(func(x float64) float64{return math.Log(x)}, 0.00001) //closure
	fmt.Printf("%f\n", deriv2(1))
	fmt.Printf("%f\n", deriv2(2))

}

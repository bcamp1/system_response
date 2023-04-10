package rk

import (
	"bcamp/rk4/vec"
)

func Step(f func(float64, vec.Vec) vec.Vec, y vec.Vec, t, h float64) vec.Vec {
	k1 := f(t, y)
	k2 := f(t+(0.5*h), vec.Add(y, vec.MulScalar(k1, 0.5*h)))
	k3 := f(t+(0.5*h), vec.Add(y, vec.MulScalar(k2, 0.5*h)))
	k4 := f(t+h, vec.Add(y, vec.MulScalar(k3, h)))

	y_next := make(vec.Vec, len(y))

	for i := range y_next {
		y_next[i] = y[i] + (h/6)*(k1[i]+2*k2[i]+2*k3[i]+k4[i])
	}

	return y_next
}

func ODE(f func(float64, vec.Vec) vec.Vec, y0 vec.Vec, t vec.Vec) vec.Matrix {
	y := make(vec.Matrix, len(t))
	for i := range y {
		y[i] = make(vec.Vec, len(y0))
	}
	h := t[1] - t[0]

	copy(y[0], y0)
	for i := 1; i < len(y); i++ {
		y_next := Step(f, y[i-1], t[i-1], h)
		copy(y[i], y_next)
	}

	return y
}

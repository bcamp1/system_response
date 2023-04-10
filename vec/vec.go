package vec

import (
	"image/color"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type Vec []float64

type Matrix []Vec

func (m Matrix) GetColumn(c int) Vec {
	vec := make(Vec, len(m))
	for i := range m {
		vec[i] = m[i][c]
	}

	return vec
}

func mul(a, b float64) float64 {
	return a * b
}

func div(a, b float64) float64 {
	return a / b
}

// func pow(a, b float64) float64 {
// 	return math.Pow(a, b)
// }

func add(a, b float64) float64 {
	return a + b
}

func sub(a, b float64) float64 {
	return a - b
}

func Range(start, end float64, n int) Vec {
	if n < 2 {
		panic("Must have 2 or more values (n)")
	}

	v := make(Vec, n)
	h := (end - start) / float64(n-1)

	for i := range v {
		v[i] = start + (h * float64(i))
	}

	return v
}

func Const(val float64, n int) Vec {
	if n < 1 {
		panic("Must have 1 or more values (n)")
	}

	v := make(Vec, n)

	for i := range v {
		v[i] = val
	}

	return v
}

func Exp(v Vec) Vec {
	return Function(v, math.Exp)
}

func Sin(v Vec) Vec {
	return Function(v, math.Sin)
}

func Cos(v Vec) Vec {
	return Function(v, math.Cos)
}

func Tan(v Vec) Vec {
	return Function(v, math.Tan)
}

func MulScalar(v Vec, s float64) Vec {
	return Function(v, func(x float64) float64 { return x * s })
}

func DivScalar(v Vec, s float64) Vec {
	return Function(v, func(x float64) float64 { return x / s })
}

func AddScalar(v Vec, s float64) Vec {
	return Function(v, func(x float64) float64 { return x + s })
}

func SubScalar(v Vec, s float64) Vec {
	return Function(v, func(x float64) float64 { return x - s })
}

func Pow(v Vec, s float64) Vec {
	return Function(v, func(x float64) float64 { return math.Pow(x, s) })
}

func vecOp(v1, v2 Vec, op func(float64, float64) float64) Vec {
	if len(v1) != len(v2) {
		panic("Lens not equal")
	}

	v3 := make(Vec, len(v1))
	for i := range v1 {
		v3[i] = op(v1[i], v2[i])
	}
	return v3
}

func Add(v1, v2 Vec) Vec {
	return vecOp(v1, v2, add)
}

func Sub(v1, v2 Vec) Vec {
	return vecOp(v1, v2, sub)
}

func Mul(v1, v2 Vec) Vec {
	return vecOp(v1, v2, mul)
}

func Div(v1, v2 Vec) Vec {
	return vecOp(v1, v2, div)
}

func Function(v Vec, f func(float64) float64) Vec {
	v_new := make(Vec, len(v))
	for i := range v_new {
		v_new[i] = f(v[i])
	}
	return v_new
}

func MakePoints(x Vec, y Vec) plotter.XYs {
	if len(x) != len(y) {
		panic("Length x != length y")
	}

	pts := make(plotter.XYs, len(x))
	for i := range pts {
		pts[i].X = x[i]
		pts[i].Y = y[i]
	}
	return pts
}

type PlotData struct {
	X, Y Vec
	C    color.Color
}

func Plot(plotData PlotData, fname string) {
	lineData := MakePoints(plotData.X, plotData.Y)

	p := plot.New()

	p.Title.Text = "Vec Plot"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	p.Add(plotter.NewGrid())

	l, err := plotter.NewLine(lineData)
	if err != nil {
		panic(err)
	}
	l.LineStyle.Width = vg.Points(0.75)
	//l.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
	l.LineStyle.Color = plotData.C

	p.Add(l)

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, fname); err != nil {
		panic(err)
	}
}

func PlotMulti(curves []PlotData, fname string) {

	p := plot.New()

	p.Title.Text = "Vec Plot"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	p.Add(plotter.NewGrid())

	for i := range curves {
		plotData := curves[i]
		lineData := MakePoints(plotData.X, plotData.Y)
		l, err := plotter.NewLine(lineData)
		if err != nil {
			panic(err)
		}
		l.LineStyle.Width = vg.Points(0.75)
		//l.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
		l.LineStyle.Color = plotData.C

		p.Add(l)

	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, fname); err != nil {
		panic(err)
	}
}

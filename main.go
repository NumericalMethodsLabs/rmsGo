package main

import (
	"fmt"
	"github.com/wcharczuk/go-chart"
	"math"
	"net/http"
)

func fx(x float64) float64 {
	return math.Cos(x)
}

func approximate(x float64, A []float64) float64 {
	Px := 0.0
	for i := 0; i < len(A); i++ {
		Px += A[i] * math.Pow(x,float64(i))
	}
	return Px
}

var xAxis []float64
var fxAxis []float64
var px []float64

func main() {
	//var n = 3 // степень многочлена

	//матрица Грамма
	//g[i][j] = Интеграл от 0 до Пи x^(i+j)
	g := [][]float64{
		{3.14159265, 4.9348022, 10.33542556, 24.35227276},
		{4.9348022, 10.33542556, 24.35227276, 61.20393696},
		{10.33542556, 24.35227276, 61.20393696, 160.23153226},
		{24.35227276, 61.20393696, 160.23153226, 431.47046111},
	}
	//d[i] = Интеграл от 0 до Пи x^(i)
	d := []float64{3.67759339e-17, -2.00000000, -6.28318531, -17.6088132}

	a := Execute(g, d)


	for i := 0.0; i < math.Pi; i += 0.01 {
		xAxis = append(xAxis, i)
		fxAxis = append(fxAxis, fx(i))
	}

	for _, val := range xAxis {
		px = append(px, approximate(val, a))
	}

	fmt.Println("a: ", a)
	fmt.Println("f(x): ", fxAxis[:5])
	fmt.Println("p(x): ", px[:5])

	http.HandleFunc("/", drawChart)
	http.ListenAndServe(":8000", nil)
}

func drawChart(res http.ResponseWriter, req *http.Request) {

	/*
	   The below will draw the same chart as the `basic` example, except with both the x and y axes turned on.
	   In this case, both the x and y axis ticks are generated automatically, the x and y ranges are established automatically, the canvas "box" is adjusted to fit the space the axes occupy so as not to clip.
	*/

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Style: chart.Style{
				//Show: true, //enables / displays the x-axis
			},
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				//Show: true, //enables / displays the y-axis
			},
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					//Show:        true,
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					StrokeWidth: 5.0,
					//FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
				},
				XValues: xAxis,
				YValues: fxAxis,
			},
		},
	}
	fmt.Println(px)
	fmt.Println(fxAxis)
	res.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, res)
}
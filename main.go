package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type Pontos struct {
	X float64
	Y float64
}

func histPlot(values plotter.Values, title string) {
	p  := plot.New()
	
	p.Title.Text = title

	hist, err := plotter.NewHist(values, 20)
	if err != nil {
			panic(err)
	}
	p.Add(hist)

	if err := p.Save(5*vg.Inch, 5*vg.Inch, fmt.Sprintf("%s.png", title)); err != nil {
			panic(err)
	}
}

func gerarPontosAleatorios () Pontos {
	rand.Seed(time.Now().UnixNano())

	pontos := Pontos{}

	pontos.X = rand.Float64()
	pontos.Y = rand.Float64()

	return pontos
}

func obterPis (pontosTeste int) plotter.Values {
	var totaisPi plotter.Values

	for i := 0; i < 1000; i ++ {
		total := 0

		for j := 0; j < pontosTeste; j ++ {
			pontos := gerarPontosAleatorios()

			inside := math.Pow(pontos.X, 2) + math.Pow(pontos.Y, 2)
			
			if inside <= 1 {
				total += 1
			}
		}

		pi := 4 * (float64(total) / float64(pontosTeste))

		totaisPi = append(totaisPi, pi)
	}

	return totaisPi
}

func main () {
	histPlot(obterPis(50), "50 pontos por teste")
	histPlot(obterPis(100), "100 pontos por teste")
	histPlot(obterPis(500), "500 pontos por teste")

}
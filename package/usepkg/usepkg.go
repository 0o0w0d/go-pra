package main

import (
	"fmt"
	"go-pra/package/usepkg/custompkg"

	"github.com/guptarohit/asciigraph"
	"github.com/tuckersGo/musthaveGo2/ch14/expkg"
)

func main()  {
	custompkg.PrintCustom()
	fmt.Println()

	expkg.PrintSample()
	data := []float64{ 3, 4, 5, 6, 9, 7, 5, 8, 5, 10, 2, 7, 2, 5, 6 }
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
}
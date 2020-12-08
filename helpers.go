package main

import (
	"fmt"
	"os"

	"github.com/wcharczuk/go-chart"
)

func getFullPath(fileName string) string {
	return fmt.Sprintf("%s%s", DataDir, fileName)
}

func plotDistribution(errorCodeDistribution map[string]int) {
	values := make([]chart.Value, 0, len(errorCodeDistribution))

	for k, v := range errorCodeDistribution {
		values = append(values, chart.Value{Value: float64(v), Label: k})
	}

	pie := chart.PieChart{
		Width:  1960,
		Height: 1080,
		Values: values,
	}

	f, err := os.Create("output/pie_chart.png")

	if err != nil {
		panic(err)
	}
	defer f.Close()
	pie.Render(chart.PNG, f)
}

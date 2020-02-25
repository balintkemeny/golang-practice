package main

import (
	"fmt"
	"math"
)

func nbMonths(startPriceOld, startPriceNew, savingperMonth int, percentLossByMonth float64) [2]int {
	i := 0
	savings := 0
	priceOld := float64(startPriceOld)
	priceNew := float64(startPriceNew)
	for {
		if float64(savings)+priceOld >= priceNew {
			break
		}
		if i%2 == 1 {
			percentLossByMonth += 0.5
		}
		//startPriceOld = int(float64(startPriceOld) * (1.0 - (percentLossByMonth / 100.0)))
		//startPriceNew = int(float64(startPriceNew) * (1.0 - (percentLossByMonth / 100.0)))
		priceOld = priceOld * (1.0 - (percentLossByMonth / 100.0))
		priceNew = priceNew * (1.0 - (percentLossByMonth / 100.0))
		savings += savingperMonth
		i++
	}
	return [2]int{i, savings + int(math.Round(priceOld-priceNew))}
}

func main() {
	fmt.Println(nbMonths(2000, 8000, 1000, 1.5))
	fmt.Println(nbMonths(12000, 8000, 1000, 1.5))
	fmt.Println(nbMonths(8000, 12000, 500, 1.0))
	fmt.Println(nbMonths(18000, 32000, 1500, 1.25))
	fmt.Println(nbMonths(7500, 32000, 300, 1.55))
}

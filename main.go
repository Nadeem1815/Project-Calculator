package main

import (
	"fmt"

	"github.com/Nadeem1815/project-calculator/filemanager"
	"github.com/Nadeem1815/project-calculator/prices"
)

func main() {

	taxRates := []float64{0, 0.07, 0.1, 0.15}

	donchans := make([]chan bool, len(taxRates))
	errChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		donchans[index] = make(chan bool)
		errChans[index] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(donchans[index], errChans[index])

		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// }

	}

	for index := range taxRates {
		select {
		case err := <-errChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-donchans[index]:
			fmt.Println("done!")
		}
	}

	// for _, errdonchan := range errChans {
	// 	<-errdonchan
	// }

	// for _, donedonchan := range donchans {
	// 	<-donedonchan
	// }

}

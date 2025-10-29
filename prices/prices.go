package prices

import (
	"fmt"

	"github.com/Nadeem1815/project-calculator/conversion"
	"github.com/Nadeem1815/project-calculator/filemanager"
)

type TaxIncludeJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludeJob) LoadData() {

	lines, err := filemanager.ReadLines("prices.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringToFloats(lines)

	if err != nil {
		fmt.Println(err)

		return
	}

	job.InputPrices = prices

}

func (job *TaxIncludeJob) Process() {
	job.LoadData()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludePrice := price * (1 + job.TaxRate)

		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludePrice)

	}
	fmt.Println(result)
}

// constructor
func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludeJob {

	return &TaxIncludeJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}

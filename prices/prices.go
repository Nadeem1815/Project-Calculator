package prices

import (
	"fmt"

	"github.com/Nadeem1815/project-calculator/conversion"
	"github.com/Nadeem1815/project-calculator/iomanger"
)

type TaxIncludeJob struct {
	IOManager         iomanger.IOManager `json:"-"`
	TaxRate           float64            `json:"tax_rate"`
	InputPrices       []float64          `json:"input_prices"`
	TaxIncludedPrices map[string]string  `json:"tax_include_prices"`
}

func (job *TaxIncludeJob) LoadData() error {

	lines, err := job.IOManager.ReadLines()

	if err != nil {

		return err
	}

	prices, err := conversion.StringToFloats(lines)

	if err != nil {

		return err
	}

	job.InputPrices = prices

	return job.IOManager.WriteResult(job)

}

func (job *TaxIncludeJob) Process(donchans chan bool, errorChan chan error) {
	err := job.LoadData()

	// errorChan <- errors.New("an error")
	if err != nil {
		// return err
		errorChan <- err
		return
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludePrice := price * (1 + job.TaxRate)

		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludePrice)

	}
	// fmt.Println(result)
	job.TaxIncludedPrices = result
	job.IOManager.WriteResult(job)
	donchans <- true

}

// constructor
func NewTaxIncludedPriceJob(io iomanger.IOManager, taxRate float64) *TaxIncludeJob {

	return &TaxIncludeJob{
		IOManager:   io,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}

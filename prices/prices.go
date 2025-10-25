package prices

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Nadeem1815/project-calculator/conversion"
)

type TaxIncludeJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludeJob) LoadData() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("could not open file")
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)

	var lines []string

	// scan return bool
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println("Reading the file content failed.")
		fmt.Println(err)
		file.Close()
		return
	}

	prices, err := conversion.StringToFloats(lines)

	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}

	job.InputPrices = prices
	file.Close()
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

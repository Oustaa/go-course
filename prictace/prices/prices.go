package prices

import (
	"bufio"
	"fmt"
	"os"

	"example.com/prictace/conversion"
)

type TaxIncludedPricejob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

func (job *TaxIncludedPricejob) LoadData() {
	file, err := os.Open("prices.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println("An error ocured whilke reading the file!")
		fmt.Println(err)
		file.Close()
		return

	}
	linesFloats, conversionErr := conversion.StringsToFloats(lines)

	if conversionErr != nil {
		fmt.Println(conversionErr)
	}

	job.InputPrices = linesFloats
	file.Close()
}

func (job *TaxIncludedPricejob) Process() {
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", price*(1+job.TaxRate))
	}

	job.TaxIncludedPrices = result

	fmt.Println(job)
}

func NewTaxIncludedPricejob(TaxRate float64) *TaxIncludedPricejob {
	return &TaxIncludedPricejob{
		TaxRate:     TaxRate,
		InputPrices: []float64{},
	}
}

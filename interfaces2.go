package main

import "fmt"

// Runtime Polymorphism

type taxSystem interface {
	calculateTax() int
}
type indianTax struct {
	taxPercentage int
	income        int
}

func (i *indianTax) calculateTax() int {
	tax := i.income * i.taxPercentage / 100
	return tax
}

type singaporeTax struct {
	taxPercentage int
	income        int
}

func (i *singaporeTax) calculateTax() int {
	tax := i.income * i.taxPercentage / 100
	return tax
}

type usaTax struct {
	taxPercentage int
	income        int
}

func (i *usaTax) calculateTax() int {
	tax := i.income * i.taxPercentage / 100
	return tax
}
func main() {
	indianTax := &indianTax{
		taxPercentage: 30,
		income:        1000,
	}
	singaporeTax := &singaporeTax{
		taxPercentage: 10,
		income:        2000,
	}

	useTax := &usaTax{
		taxPercentage: 20,
		income:        3000,
	}

	taxSystems := []taxSystem{indianTax, singaporeTax, useTax}
	totalTax := calculateTotalTax(taxSystems)

	fmt.Printf("Total Tax is %d\n", totalTax)
}
func calculateTotalTax(taxSystems []taxSystem) int {
	totalTax := 0
	for _, t := range taxSystems {
		totalTax += t.calculateTax() //This is where runtime polymorphism happens
	}
	return totalTax
}

package tax

type Repository interface {
	SaveTax(amount float64) error
}

func CalculateTaxAndSave(amount float64, repository Repository) error {
	tax := CalculateTax(amount)
	return repository.SaveTax(tax)
}

func CalculateTax(amount float64) float64 {
	if amount > 1000 {
		return 10.0
	}

	return 5.0
}

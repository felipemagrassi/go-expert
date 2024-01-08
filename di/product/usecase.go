package product

type ProductUseCase struct {
	repository ProductRepositoryInterface
}

func NewProductUseCase(repository ProductRepositoryInterface) *ProductUseCase {
	return &ProductUseCase{repository}
}

func (uc *ProductUseCase) GetProduct(id string) (Product, error) {
	return uc.repository.GetProduct(id)
}

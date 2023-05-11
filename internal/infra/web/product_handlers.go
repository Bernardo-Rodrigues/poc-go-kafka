package web

import (
	"encoding/json"
	"net/http"

	"test/internal/usecase"
)

type ProductHandlers struct {
	CreateProductUseCase *usecase.CreateProductUseCase
	ListProductsUseCase  *usecase.ListProductsUseCase
}

func NewProductHandlers(createProductUseCase *usecase.CreateProductUseCase, listProductsUseCase *usecase.ListProductsUseCase) *ProductHandlers {
	return &ProductHandlers{
		CreateProductUseCase: createProductUseCase,
		ListProductsUseCase:  listProductsUseCase,
	}
}

func (p *ProductHandlers) CreateProductHandler(response http.ResponseWriter, request *http.Request) {
	var input usecase.CreateProductInputDto
	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := p.CreateProductUseCase.Execute(input)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(output)
}

func (p *ProductHandlers) ListProductsHandler(response http.ResponseWriter, request *http.Request) {
	output, err := p.ListProductsUseCase.Execute()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(output)
}

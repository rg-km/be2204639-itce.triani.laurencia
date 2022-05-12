package api

import (
	"encoding/json"
	"net/http"
)

type ProductListErrorResponse struct {
	Error string `json:"error"`
}

type Product struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Category string `json:"category"`
}

type ProductListSuccessResponse struct {
	Products []Product `json:"products"`
}

func (api *API) productList(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	encoder := json.NewEncoder(w)

	response := ProductListSuccessResponse{}
	response.Products = make([]Product, 0)

	var arrProducts []Product
	products, err := api.productsRepo.SelectAll()
	for i := 0; i < len(products); i++ {
		newDataProduct := Product{
			Name:     products[i].ProductName,
			Price:    products[i].Price,
			Category: products[i].Category,
		}
		arrProducts = append(arrProducts, newDataProduct)
	}

	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(DashboardErrorResponse{Error: err.Error()})
			return
		}
	}()
	if err != nil {
		return
	}

	for _, product := range products {
		response.Products = append(response.Products, Product{
			Name:     product.ProductName,
			Price:    product.Price,
			Category: product.Category,
		})
	}
	encoder.Encode(response)
	// encoder.Encode(ProductListSuccessResponse{Products: []Product{}}) // TODO: replace this
}
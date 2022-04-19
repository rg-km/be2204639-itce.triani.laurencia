package repository

import (
	"strconv"

	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/db"
)

type ProductRepository struct {
	db db.DB
}

func NewProductRepository(db db.DB) ProductRepository {
	return ProductRepository{db}
}

func (u *ProductRepository) LoadOrCreate() ([]Product, error) {
<<<<<<< HEAD
<<<<<<< HEAD
	//return []Product{}, nil

	// buka file datanya
	// kita lakukan readall (akses membaca file)
	// ubah json => menjadi struct

	data, err := u.db.Load("products")
	if err != nil {
		var head = [][]string{
			{"category", "product_name", "price"},
		}

		if err = u.db.Save("products", head); err != nil {
			return nil, err
		}
	}

	var products []Product

	for i := 1; i < len(data); i++ {
		price, err := strconv.Atoi(data[i][2])
		if err != nil {
			return nil, err
		}

		products = append(products, Product{
			Category:    data[i][0],
			ProductName: data[i][1],
			Price:       price,
		})
	}

	return products, nil
}

func (u *ProductRepository) SelectAll() ([]Product, error) {
	return u.LoadOrCreate()
}
=======
=======
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
	return []Product{}, nil // TODO: replace this
}

func (u *ProductRepository) SelectAll() ([]Product, error) {
	return []Product{}, nil // TODO: replace this
}
<<<<<<< HEAD
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
=======
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015

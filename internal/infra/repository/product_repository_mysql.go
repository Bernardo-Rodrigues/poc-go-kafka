package repository

import (
	"database/sql"
	"test/internal/entity"
)

type ProductRepositoryMySql struct {
	DB *sql.DB
}

func NewProductRepositoryMySql(db *sql.DB) *ProductRepositoryMySql {
	return &ProductRepositoryMySql{DB: db}
}

func (repository *ProductRepositoryMySql) Create(product *entity.Product) error {
	_, err := repository.DB.Exec("Insert into products (id, name, price) values(?,?,?)",
		product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func (repository *ProductRepositoryMySql) FindAll() ([]*entity.Product, error) {
	rows, err := repository.DB.Query("Select * from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product

	for rows.Next() {
		var product entity.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}

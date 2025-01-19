package products

import (
	rep "web-store-go/infra/config/repository"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetAllProducts() []Product {
	db := rep.Conectar()
	result, error := db.Query("select * from produtos")
	if error != nil {
		panic(error.Error())
	}

	products := []Product{}
	product := Product{}
	for result.Next() {
		var id int
		var nome string
		var description string
		var price float64
		var quantity int

		err := result.Scan(&id, &nome, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = nome
		product.Description = description
		product.Price = price
		product.Quantity = quantity
		products = append(products, product)
	}

	defer db.Close()
	return products
}

func CreateNew(name string, description string, quantity int, price float64) {
	insert := "INSERT INTO produtos (nome, descricao, quantidade, preco) VALUES ($1, $2, $3, $4)"

	db := rep.Conectar()

	statement, error := db.Prepare(insert)
	if error != nil {
		panic(error.Error())
	}

	statement.Exec(name, description, quantity, price)
	defer db.Close()
}

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
		var descricao string
		var preco float64
		var quantidade int

		err := result.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = nome
		product.Description = descricao
		product.Price = preco
		product.Quantity = quantidade
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

func DeleteProduct(id string) {
	delete := "DELETE FROM produtos WHERE id = $1"
	db := rep.Conectar()

	statement, error := db.Prepare(delete)
	if error != nil {
		panic(error.Error())
	}

	statement.Exec(id)
	defer db.Close()
}

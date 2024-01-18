package domain

import "SouzaBernardo/first-go-api/database"

type Product struct {
	Id 			int
	Name        string
	Description string
	Price       float64
	Quantity    int
}


func GetProducts() []Product {
	db := database.ConnectDB()

	allProduct, err := db.Query("select * from product")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for allProduct.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = allProduct.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)

	}

	defer db.Close()
	return products
}
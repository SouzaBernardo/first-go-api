package entity

import (
	"github.com/SouzaBernardo/first-go-api/model"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func Get(productId string) Product {
	db := model.ConnectDB()

	get, err := db.Query("select * from product where id=$1", productId)

	if err != nil {
		panic(err.Error())
	}

	var id, quantity int
	var name, description string
	var price float64

	for get.Next() {
		err := get.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}
	}

	product := Product{
		Id:          id,
		Name:        name,
		Description: description,
		Price:       price,
		Quantity:    quantity,
	}

	defer db.Close()
	return product
}

func GetAll() []Product {
	db := model.ConnectDB()

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

func Create(name string, description string, price float64, quantity int) {
	db := model.ConnectDB()

	insert, err := db.Prepare("insert into product(name, description, price, quantity) values ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insert.Exec(name, description, price, quantity)
	defer db.Close()
}

func Delete(id string) {
	db := model.ConnectDB()
	remove, err := db.Prepare("remove from product where id=$1")

	if err != nil {
		panic(err.Error())
	}

	remove.Exec(id)
	db.Close()
}

func Edit(id string, name, description string, price float64, quantity int) {
	db := model.ConnectDB()

	update, err := db.Prepare("update product set name=$2, description=$3, price=$4, quantity=$5 where id=$1")

	if err != nil {
		panic(err.Error())
	}

	update.Exec(id, name, description, price, quantity)
	db.Close()
}

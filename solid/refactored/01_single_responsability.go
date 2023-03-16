package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Product struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type Order struct {
	Id       int        `json:"id"`
	Products []*Product `json:"products"`
}

func (o *Order) AddProduct(code int, description string) {
	o.Products = append(o.Products, &Product{
		Code:        code,
		Description: description,
	})
}

func (o *Order) Count() int {
	return len(o.Products)
}

// WRONG
func (o *Order) MustSave() {
	fileName := fmt.Sprintf(
		"order_%d.txt", o.Id,
	)
	f, err := os.Create(fileName)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	content, err := json.Marshal(o)
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.WriteString(string(content))

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	order := Order{Id: 1}
	order.AddProduct(1, "Pepperoni Pizza")
	order.AddProduct(2, "Marguerita Pizza")
	order.MustSave()

}

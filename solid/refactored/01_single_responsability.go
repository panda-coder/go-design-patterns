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

func (o *Order) ToJSON() ([]byte, error) {
	return json.Marshal(o)
}
func (o *Order) GetId() int {
	return o.Id
}

func (o *Order) Count() int {
	return len(o.Products)
}

type DataModel interface {
	GetId() int
	ToJSON() ([]byte, error)
}

type DataSaver interface {
	Save(d DataModel) error
}

type FileDataSaver struct{}

func (fos FileDataSaver) Save(prefix string, d DataModel) error {
	fileName := fmt.Sprintf("%s_%d.txt", prefix, d.GetId())
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	content, err := d.ToJSON()
	if err != nil {
		return err
	}

	_, err = f.WriteString(string(content))
	if err != nil {
		return err
	}
	return nil
}

func main() {
	order := Order{Id: 1}
	order.AddProduct(1, "Pepperoni Pizza")
	order.AddProduct(2, "Marguerita Pizza")
	fileSaver := FileDataSaver{}
	err := fileSaver.Save("order", &order)
	if err != nil {
		log.Fatal(err)
	}
}

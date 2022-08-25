package main

import (
	"fmt"
	"log"

	"github.com/hromov/query/evaluator"
	"github.com/hromov/query/products"
)

type EvaluatorService interface {
	Eval(string) (float64, error)
}

type Logic struct {
	Constrain string
	Rule      string
}

var productList = []products.Product{
	{Name: "Parcel 1", Height: 50, Price: 100},
	{Name: "Parcel 2", Height: 75, Price: 200},
	{Name: "Parcel 3", Height: 250, Price: 300},
}

var logics = []Logic{
	{Constrain: averageHight, Rule: "%d * 1.1 + 10"},
	{Constrain: increasedHight, Rule: "%d * 2"},
}

func main() {
	es := evaluator.NewService()
	for _, p := range productList {
		newPrice, err := getNewPrice(es, p, logics)
		if err != nil {
			log.Printf("Error: %s\n", err)
		} else {
			log.Printf("\tNew Price for Product %s is %d. Old price is %d\n", p.Name, newPrice, p.Price)
		}
	}

}

func getNewPrice(es EvaluatorService, p products.Product, logics []Logic) (int, error) {
	price := p.Price
	for _, logic := range logics {
		if p.IsValid(logic.Constrain) {
			if newPrice, err := es.Eval(fmt.Sprintf(logic.Rule, price)); err != nil {
				return 0, err
			} else {
				price = int(newPrice)
			}
		}
	}
	return price, nil
}

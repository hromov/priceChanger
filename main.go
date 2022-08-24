package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/diegoholiveira/jsonlogic/v3"
	"github.com/hromov/query/evaluator"
)

type Product struct {
	Name   string `json:"name"`
	Height int    `json:"height"`
	Price  int    `json:"price"`
}

func (p *Product) isOk(logic string) bool {
	productBytes, _ := json.Marshal(p)
	data := strings.NewReader(string(productBytes))
	logicReader := strings.NewReader(logic)

	var result bytes.Buffer

	jsonlogic.Apply(logicReader, data, &result)
	return strings.Contains(result.String(), "true")
}

var products = []Product{
	{Name: "Parcel 1", Height: 50, Price: 100},
	{Name: "Parcel 2", Height: 75, Price: 200},
	{Height: 250, Price: 300},
}
var logics []string = []string{averageHight, increasedHight}
var rules []string = []string{`%d * 1.1 + 10`, `%d * 2`}

func main() {
	es := evaluator.NewService()
	for _, p := range products {
		for i, logic := range logics {
			if p.isOk(logic) {
				// fmt.Printf("Product Height %d is ok for logick = %s\n", p.Height, logic)
				res, err := es.Eval(fmt.Sprintf(rules[i], p.Price))
				if err != nil {
					log.Println(err)
				}
				log.Printf("Product Height %d, price = %s\n", p.Height, res)
			}
		}
	}
}

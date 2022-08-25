package products

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/diegoholiveira/jsonlogic"
)

type Product struct {
	Name   string `json:"name"`
	Height int    `json:"height"`
	Price  int    `json:"price"`
}

func (p *Product) IsValid(logic string) bool {
	productBytes, _ := json.Marshal(p)
	data := strings.NewReader(string(productBytes))
	logicReader := strings.NewReader(logic)

	var result bytes.Buffer

	jsonlogic.Apply(logicReader, data, &result)
	return strings.Contains(result.String(), "true")
}

package candlestick

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	c := Candlestick{}
	file, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Printf("File error %v\n", err)
		os.Exit(1)
	}
	var data interface{}

	json.Unmarshal(file, &data)

	var candlesF Data
	c.Load(data, &candlesF)
	if len(candlesF.Date) > 0 {
		fmt.Println("Exito")
	}
}

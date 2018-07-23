package candlestick

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
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

	candles := c.Load(data)
	fmt.Println(reflect.TypeOf(candles).Kind())

}

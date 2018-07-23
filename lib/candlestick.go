package lib

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Candlestick struct {
	data interface{}
}

type Data struct {
	Date   []time.Time `json:"date"`
	Open   []float64   `json:"open"`
	High   []float64   `json:"high"`
	Low    []float64   `json:"low"`
	Close  []float64   `json:"close"`
	Volume []float64   `json:"volume"`
}
type Candle struct {
	Date   time.Time
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume float64
}

type Result struct {
	Candlesticks []Candlestick
}

func (data Candlestick) Load(d interface{}) {
	//var candlesticks Data
	data.data = d
	var candles Data
	switch reflect.TypeOf(data.data).Kind() {
	case reflect.Slice:
		dataParsed := reflect.ValueOf(data.data)
		numRows := dataParsed.Len()
		//fmt.Println(dataParsed)
		c := NewCandlestick(numRows)

		for i := 0; i < numRows; i++ {
			/*c.Date[i], _ = time.Parse("2006-01-02T15:04:05", dataParsed.Index(i).Time) //"2017-11-28T16:50:00"
			c.Open[i] = dataParsed.Index(i)["Open"]
			c.High[i] = dataParsed.Index(i).High
			c.Low[i] = dataParsed.Index(i).Low
			c.Close[i] = dataParsed.Index(i).Close
			c.Volume[i] = dataParsed.Index(i).Volume*/
			//fmt.Println(dataParsed.Index(i))

			dd := dataParsed.Index(i)
			original := dd.Elem()
			//fmt.Println(dd.Kind())
			//fmt.Println(original.MapIndex(original.MapKeys()[1]))
			for j := 0; j < 5; j++ {
				key := original.MapKeys()[j]
				val := original.MapIndex(key).Elem().String()
				if strings.Compare(key.String(), "low") == 0 {
					num, err := strconv.ParseFloat(val, 64)
					if err != nil {
						fmt.Println(err)
					}
					//fmt.Println(num)
					c.Low[i] = num
				} else if strings.Compare(key.String(), "high") == 0 {
					num, err := strconv.ParseFloat(val, 64)
					if err != nil {
						fmt.Println(err)
					}
					c.High[i] = num
				} else if strings.Compare(key.String(), "open") == 0 {
					num, err := strconv.ParseFloat(val, 64)
					if err != nil {
						fmt.Println(err)
					}
					//fmt.Println(num)
					c.Open[i] = num
				} else if strings.Compare(key.String(), "close") == 0 {
					num, err := strconv.ParseFloat(val, 64)
					if err != nil {
						fmt.Println(err)
					}
					//fmt.Println(num)
					c.Close[i] = num
				} else if strings.Compare(key.String(), "volume") == 0 {
					num, err := strconv.ParseFloat(val, 64)
					if err != nil {
						fmt.Println(err)
					}
					//fmt.Println("Volume: ", num)
					c.Volume[i] = num
				} else if strings.Compare(key.String(), "date") == 0 {
					c.Date[i], _ = time.Parse("2006-01-02T15:04:05", val)
				}
			}
		}
		candles.Date = append(candles.Date, c.Date...)
		candles.Open = append(candles.Open, c.Open...)
		candles.High = append(candles.High, c.High...)
		candles.Low = append(candles.Low, c.Low...)
		candles.Close = append(candles.Close, c.Close...)
		candles.Volume = append(candles.Volume, c.Volume...)
		fmt.Println(candles.Volume)
	}
}
func NewCandlestick(bars int) Data {
	return Data{
		Date:   make([]time.Time, bars),
		Open:   make([]float64, bars),
		High:   make([]float64, bars),
		Low:    make([]float64, bars),
		Close:  make([]float64, bars),
		Volume: make([]float64, bars),
	}
}

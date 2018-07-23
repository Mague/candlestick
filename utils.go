package candlestick

import (
	"math"
)

func CandlesLen(candle *Candle) float64 {
	return math.Abs(candle.Open - candle.Close)
}

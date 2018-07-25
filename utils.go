package candlestick

import (
	"math"
)

func BodyLen(candle *Candle) float64 {
	return math.Abs(candle.Open - candle.Close)
}
func WickLen(candle *Candle) float64 {
	return candle.High - math.Max(candle.Open, candle.Close)
}
func TailLen(candle *Candle) float64 {
	return math.Min(candle.Open, candle.Close) - candle.Low
}
func IsBullish(candle *Candle) bool {
	return candle.Open < candle.Close
}
func IsBearish(candle *Candle) bool {
	return candle.Open > candle.Close
}
func IsHammerLike(candle *Candle) bool {
	return WickLen(candle) > (BodyLen(candle)*2) && TailLen(candle) < BodyLen(candle)
}
func IsInvertedHammerLike(candle *Candle) bool {
	return WickLen(candle) > (BodyLen(candle)*2) && TailLen(candle) < BodyLen(candle)
}

func IsEngulfed(candleShorTest *Candle, candleLongTest *Candle) bool {
	return BodyLen(candleShorTest) < BodyLen(candleLongTest)
}

func IsGap(lowest *Candle, upMost *Candle) bool {
	return math.Max(lowest.Open, lowest.Close) < math.Min(upMost.Open, upMost.Close)
}
func IsGapUp(previous *Candle, current *Candle) bool {
	return IsGap(previous, current)
}

func IsGapDown(previous *Candle, current *Candle) bool {
	return IsGap(current, previous)
}

//Pattern detection
func IsHammer(candle *Candle) bool {
	return IsBullish(candle) && IsInvertedHammerLike(candle)
}
func IsInvertedHammer(candle *Candle) bool {
	return IsBearish(candle) && IsInvertedHammerLike(candle)
}
func IsHangingMan(previous *Candle, current *Candle) bool {
	return IsBullish(previous) && IsBearish(current) && IsGapUp(previous, current) && IsHammerLike(current)
}
func IsShootingStar(previous *Candle, current *Candle) bool {
	return IsBullish(previous) && IsBearish(current) && IsGapUp(previous, current) && IsInvertedHammerLike(current)
}

func IsBullishEngulfing(previous *Candle, current *Candle) bool {
	return IsBearish(previous) && IsBullish(current) && IsEngulfed(previous, current)
}

func IsBearishEngulfing(previous *Candle, current *Candle) bool {
	return IsBullish(previous) && IsBearish(current) && IsEngulfed(previous, current)
}

func IsBullishHarami(previous *Candle, current *Candle) bool {
	return IsBearish(previous) && IsBullish(current) && IsEngulfed(current, previous)
}
func IsBearishHarami(previous *Candle, current *Candle) bool {
	return IsBullish(previous) && IsBearish(current) && IsEngulfed(current, previous)
}

func IsBullishKicker(previous *Candle, current *Candle) bool {
	return IsBearish(previous) && IsBullish(current) && IsGapUp(previous, current)
}

func IsBearishKicker(previous *Candle, current *Candle) bool {
	return IsBullish(previous) &&
		IsBearish(current) &&
		IsGapDown(previous, current)
}

// Dynamic array search for callback arguments.
/*func FindPattern(data *Data, cb func(*Candle) bool) int {
	upperBound := len(data.Date)
	//matches := make([]bool, 0, 1)
	cont := 0
	for i := 0; i < upperBound; i++ {
		candle := Candle{
			Date:   data.Date[i],
			Open:   data.Open[i],
			Close:  data.Close[i],
			High:   data.High[i],
			Low:    data.Low[i],
			Volume: data.Volume[i],
		}
		if cb(&candle) {
			//append(matches, args[1])
			fmt.Println("Patron detectado")
			cont++
		}

	}

	return cont
}*/

package candlestick

import (
	"strconv"
	"testing"
	"time"
)

func TestBodyLen(t *testing.T) {

	BodyLen(getCandle())
}
func TestWickLen(t *testing.T) {
	WickLen(getCandle())
}
func TestTailLen(t *testing.T) {
	TailLen(getCandle())
}
func TestIsBullish(t *testing.T) {
	IsBullish(getCandle())
}
func TestIsBearish(t *testing.T) {
	IsBearish(getCandle())
}
func TestIsHammerLike(t *testing.T) {
	IsHammerLike(getCandle())
}
func TestIsInvertedHammerLike(t *testing.T) {
	IsInvertedHammerLike(getCandle())
}
func TestIsEngulfed(t *testing.T) {
	IsEngulfed(getCandle(), getCandle())
}
func TestIsGap(t *testing.T) {
	IsGap(getCandle(), getCandle())
}
func TestIsGapUp(t *testing.T) {
	IsGapUp(getCandle(), getCandle())
}
func TestIsGapDown(t *testing.T) {
	IsGapDown(getCandle(), getCandle())
}
func TestIsHammer(t *testing.T) {
	IsHammer(getCandle())
}
func TestIsInvertedHammer(t *testing.T) {
	IsInvertedHammer(getCandle())
}
func TestIsHangingMan(t *testing.T) {
	IsHangingMan(getCandle(), getCandle())
}
func TestIsShootingStar(t *testing.T) {
	IsShootingStar(getCandle(), getCandle())
}
func TestIsBullishEngulfing(t *testing.T) {
	IsBullishEngulfing(getCandle(), getCandle())
}
func TestIsBearishEngulfing(t *testing.T) {
	IsBearishEngulfing(getCandle(), getCandle())
}
func TestIsBullishHarami(t *testing.T) {
	IsBullishHarami(getCandle(), getCandle())
}
func TestIsBearishHarami(t *testing.T) {
	IsBearishHarami(getCandle(), getCandle())
}
func TestIsBullishKicker(t *testing.T) {
	IsBullishKicker(getCandle(), getCandle())
}
func TestIsBearishKicker(t *testing.T) {
	IsBearishKicker(getCandle(), getCandle())
}
func getCandle() *Candle {
	tt, _ := time.Parse("2006-01-02 15:04:05", "2018-07-25 15:16:00")
	open, _ := strconv.ParseFloat("0.05751399", 64)
	close, _ := strconv.ParseFloat("0.05750398", 64)
	high, _ := strconv.ParseFloat("0.05751399", 64)
	low, _ := strconv.ParseFloat("0.05750398", 64)
	volume, _ := strconv.ParseFloat("0.43611832", 64)
	candle := Candle{
		Date:   tt,
		Open:   open,
		Close:  close,
		High:   high,
		Low:    low,
		Volume: volume,
	}
	return &candle
}

package strategy

import (
	"fmt"
	"quant/base/bar"
	"quant/base/indicator"
	"quant/base/strategy"
)

const (
	debug = false

	test_bar_series_and_sma = true
)

func init() {
	if debug {
		fmt.Println("quanta.strategy package init")
	}
}

type MyStrategy struct {
	strategy.Strategy

	sma   *indicator.SMA
	ema16 *indicator.EMA
	ema26 *indicator.EMA
}

func (this *MyStrategy) Init(symbol string, barSeries *bar.BarSeries) {
	if debug {
		fmt.Println("MyStrategy.Init()")
	}
	this.Strategy.Init(symbol, barSeries)

	// ToDo: init your strategy below:
	this.Name = "MyStrategy"
}

func (this *MyStrategy) OnStrategyStart() {
	fmt.Println("MyStrategy", "OnStrategyStart", "Name", this.Name, "Symbol:", this.Symbol)

	if test_bar_series_and_sma {
		this.sma = indicator.NewSMA(this.BarSeries, 6)
		this.ema16 = indicator.NewEMA(this.BarSeries, 16)
		this.ema26 = indicator.NewEMA(this.BarSeries, 26)
	}
}

func (this *MyStrategy) OnStrategyStop() {
	fmt.Println("MyStrategy", "OnStrategyStop", "Name", this.Name, "Symbol:", this.Symbol)

	if test_bar_series_and_sma {
		printslice("base:", this.BarSeries.Values())
		printslice("sma:", this.sma.Values())
		printslice("ema16:", this.ema16.Values())
		printslice("ema26:", this.ema26.Values())

	}
}

func (this *MyStrategy) OnBar(bar_ bar.Bar) {
	if debug {
		fmt.Println("MyStrategy", this.Symbol, "OnBar")
	}

	// now := this.BarSeries.Now()
	now := bar_.DateTime

	if !this.sma.IsFake(&now) {
		if this.sma.ValueAtTime(&now) > this.BarSeries.ValueAtTime(&now) {
			fmt.Print("sma > base: Buy\n")

			printslice("base:", this.BarSeries.Values())
			printslice("sma:", this.sma.Values())
		}
	}
}

func printslice(tag string, values []float64) {
	fmt.Print(tag, " ")
	for _, one := range values {
		fmt.Printf("%0.3f ", one)
	}
	fmt.Print("\n")

}

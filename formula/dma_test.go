package formula

import (
	"fmt"
	"gitee.com/quant1x/pandas"
	"gitee.com/quant1x/pandas/stat"
	"testing"
)

func TestDMA(t *testing.T) {
	f0 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	s0 := pandas.NewSeriesWithoutType("f0", f0)
	fmt.Println(DMA(s0, 5))
	s2 := []float64{1, 2, 3, 4, 3, 3, 2, 1, stat.DTypeNaN, stat.DTypeNaN, stat.DTypeNaN, stat.DTypeNaN}
	fmt.Println(s2)
}

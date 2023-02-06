package formula

import (
	"fmt"
	"gitee.com/quant1x/pandas"
	"testing"
)

func TestIF(t *testing.T) {
	S := pandas.NewSeries(pandas.SERIES_TYPE_FLOAT32, "", []float32{1, 1, 1})
	A := pandas.NewSeries(pandas.SERIES_TYPE_FLOAT32, "", []float32{11, 12, 13})
	B := pandas.NewSeries(pandas.SERIES_TYPE_FLOAT32, "", []float32{21, 22, 23})
	fmt.Println(IF(S, A, B))
}

func TestIFF(t *testing.T) {
	S := pandas.NewSeries(pandas.SERIES_TYPE_FLOAT32, "", []float32{1, 1, 1})
	A := pandas.NewSeries(pandas.SERIES_TYPE_FLOAT32, "", []float32{11, 12, 13})
	B := pandas.NewSeries(pandas.SERIES_TYPE_FLOAT32, "", []float32{21, 22, 23})
	fmt.Println(IFF(S, A, B))
}

func TestIFN(t *testing.T) {
	S := pandas.NewSeries(pandas.SERIES_TYPE_FLOAT32, "", []float32{1, 0, 1})
	A := pandas.NewSeries(pandas.SERIES_TYPE_FLOAT32, "", []float32{11, 12, 13})
	B := pandas.NewSeries(pandas.SERIES_TYPE_FLOAT32, "", []float32{21, 22, 23})
	fmt.Println(IFN(S, A, B))
}

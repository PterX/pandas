package formula

import (
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/pandas/stat"
)

// ABS 计算S的绝对值
func ABS(S stat.Series) stat.Series {
	s := S.DTypes()
	d := num.Abs(s)
	//fmt.Printf("%p\n", d)
	//return pandas.NewNDArray(stat.SERIES_TYPE_DTYPE, "", d)
	//return stat.NewNDArray(d...)
	return stat.ToSeries(d...)
}

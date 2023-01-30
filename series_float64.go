package pandas

import (
	"fmt"
	"gitee.com/quant1x/pandas/algorithms/winpooh32/math"
	"github.com/huandu/go-clone"
	"github.com/viterin/vek"
	"strconv"
)

type SeriesFloat64 struct {
	SeriesFrame
	Data []float64
}

func NewSeriesFloat64(name string, vals ...interface{}) *SeriesFloat64 {
	series := SeriesFloat64{
		SeriesFrame: SeriesFrame{
			name:         name,
			nilCount:     0,
			valFormatter: DefaultValueFormatter,
		},
		Data: []float64{},
	}

	series.Data = make([]float64, 0) // Warning: filled with 0.0 (not NaN)
	size := len(series.Data)
	for idx, v := range vals {
		// Special case
		if idx == 0 {
			if fs, ok := vals[0].([]float64); ok {
				for idx, v := range fs {
					val := series.valToPointer(v)
					if isNaN(val) {
						series.nilCount++
					}
					if idx < size {
						series.Data[idx] = val
					} else {
						series.Data = append(series.Data, val)
					}
				}
				break
			}
		}

		val := series.valToPointer(v)
		if isNaN(val) {
			series.nilCount++
		}

		if idx < size {
			series.Data[idx] = val
		} else {
			series.Data = append(series.Data, val)
		}
	}

	var lVals int
	if len(vals) > 0 {
		if fs, ok := vals[0].([]float64); ok {
			lVals = len(fs)
		} else {
			lVals = len(vals)
		}
	}

	if lVals < size {
		series.nilCount = series.nilCount + size - lVals
		// Fill with NaN
		for i := lVals; i < size; i++ {
			series.Data[i] = nan()
		}
	}

	return &series
}

func (s *SeriesFloat64) valToPointer(v interface{}) float64 {
	switch val := v.(type) {
	case nil:
		return nan()
	case bool:
		if val == true {
			return float64(1)
		}
		return float64(0)
	case int:
		return float64(val)
	case int32:
		return float64(val)
	case int64:
		return float64(val)
	case float64:
		return val
	case string:
		f, err := strconv.ParseFloat(val, 64)
		if err != nil {
			_ = v.(float64) // Intentionally panic
		}
		return f
	default:
		f, err := strconv.ParseFloat(fmt.Sprintf("%v", v), 64)
		if err != nil {
			_ = v.(float64) // Intentionally panic
		}
		return f
	}
}

// Type returns the type of data the series holds.
func (s *SeriesFloat64) Type() string {
	return SERIES_TYPE_FLOAT
}

func (s *SeriesFloat64) NRows() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.Data)
}

func (s *SeriesFloat64) Shift(periods int) *Series {
	var d Series
	d = clone.Clone(s).(Series)
	if periods == 0 {
		return &d
	}

	values := d.Values().([]float64)

	var (
		naVals []float64
		dst    []float64
		src    []float64
	)

	if shlen := int(math.Abs(float64(periods))); shlen < len(values) {
		if periods > 0 {
			naVals = values[:shlen]
			dst = values[shlen:]
			src = values
		} else {
			naVals = values[len(values)-shlen:]
			dst = values[:len(values)-shlen]
			src = values[shlen:]
		}

		copy(dst, src)
	} else {
		naVals = values
	}

	for i := range naVals {
		naVals[i] = math.NaN()
	}

	return &d
}

func (s *SeriesFloat64) Values() any {
	return s.Data
}

func (s *SeriesFloat64) Repeat(x any, repeats int) *Series {
	a := s.valToPointer(x)

	//switch val := x.(type) {
	//case int:
	//	a = float64(val)
	//case int32:
	//	a = float64(val)
	//case int64:
	//	a = float64(val)
	//default:
	//	a = float64(val)
	//
	//}
	//
	//switch reflect.TypeOf(x).Kind() {
	//case reflect.Int, reflect.Int32, reflect.Int64:
	//	a = x.(float64)
	//case reflect.Float32, reflect.Float64:
	//	a = x.(float64)
	//}
	//if f, ok := x.(float64); ok {
	//	a = f
	//} else {
	//	a = nan()
	//}
	data := vek.Repeat_Into(s.Data, a, repeats)
	var d Series
	d = NewSeriesFloat64(s.name, data)
	return &d
}

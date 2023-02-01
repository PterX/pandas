package pandas

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"math"
	"reflect"
)

// Type is a convenience alias that can be used for a more type safe way of
// reason and use Series types.
type Type = string

// Supported Series Types
const (
	SERIES_TYPE_INVAILD = "unknown" // 未知类型
	SERIES_TYPE_BOOL    = "bool"    // 布尔类型
	SERIES_TYPE_INT     = "int"     // int64
	SERIES_TYPE_FLOAT   = "float"   // float64
	SERIES_TYPE_STRING  = "string"  // string
)

// StringFormatter is used to convert a value
// into a string. Val can be nil or the concrete
// type stored by the series.
type StringFormatter func(val interface{}) string

type Series interface {
	// Name 取得series名称
	Name() string
	// Rename renames the series.
	Rename(n string)
	// Type returns the type of data the series holds.
	// 返回类型的字符串
	Type() Type
	// NRows 获得行数
	Len() int
	// Values 获得全部数据集
	Values() any
	// Empty returns an empty Series of the same type
	Empty() Series
	// Records returns the elements of a Series as a []string
	Records() []string
	// Copy 复制
	Copy() Series
	// Subset 获取子集
	Subset(start, end int, opt ...any) Series
	// Repeat elements of an array.
	Repeat(x any, repeats int) Series
	// Shift index by desired number of periods with an optional time freq.
	// 使用可选的时间频率按所需的周期数移动索引.
	Shift(periods int) Series
	// Rolling creates new RollingWindow
	Rolling(window int) RollingWindow
	// Mean calculates the average value of a series
	Mean() float64
	// StdDev calculates the standard deviation of a series
	StdDev() float64
}

// NewSeries 指定类型创建序列
func NewSeries(t Type, name string, vals ...interface{}) *Series {
	var series Series
	if t == SERIES_TYPE_BOOL {
		series = NewSeriesBool(name, vals...)
	} else if t == SERIES_TYPE_INT {
		series = NewSeriesInt64(name, vals...)
	} else if t == SERIES_TYPE_STRING {
		series = NewSeriesString(name, vals...)
	} else {
		// 默认全部强制转换成float64
		series = NewSeriesFloat64(name, vals...)
	}
	return &series
}

// 泛型方法, 构造序列, 比其它方式对类型的统一性要求更严格
func GenericSeries[T GenericType](name string, values ...T) *Series {
	// 第一遍, 确定类型, 找到第一个非nil的值
	var _type Type = SERIES_TYPE_STRING
	for _, v := range values {
		// 泛型处理这里会出现一个错误, invalid operation: v == nil (mismatched types T and untyped nil)
		//if v == nil {
		//	continue
		//}
		vv := reflect.ValueOf(v)
		vk := vv.Kind()
		switch vk {
		//case reflect.Invalid: // {interface} nil
		//	series.assign(idx, size, Nil2Float)
		//case reflect.Slice: // 切片, 不定长
		//	for i := 0; i < vv.Len(); i++ {
		//		tv := vv.Index(i).Interface()
		//		str := AnyToFloat64(tv)
		//		series.assign(idx, size, str)
		//	}
		//case reflect.Array: // 数组, 定长
		//	for i := 0; i < vv.Len(); i++ {
		//		tv := vv.Index(i).Interface()
		//		av := AnyToFloat64(tv)
		//		series.assign(idx, size, av)
		//	}
		//case reflect.Struct: // 忽略结构体
		//	continue
		//default:
		//	vv := AnyToFloat64(val)
		//	series.assign(idx, size, vv)
		case reflect.Bool:
			_type = SERIES_TYPE_BOOL
		case reflect.Int64:
			_type = SERIES_TYPE_INT
		case reflect.Float64:
			_type = SERIES_TYPE_FLOAT
		case reflect.String:
			_type = SERIES_TYPE_STRING
		default:
			panic(fmt.Errorf("unknown type, %+v", v))
		}
		break
	}
	return NewSeries(_type, name, values)
}

// DefaultIsEqualFunc is the default comparitor to determine if
// two values in the series are the same.
func DefaultIsEqualFunc(a, b interface{}) bool {
	return cmp.Equal(a, b)
}

// DefaultFormatter will return a string representation
// of the data in a particular row.
func DefaultFormatter(v interface{}) string {
	if v == nil {
		return StringNaN
	}
	return fmt.Sprintf("%v", v)
}

func detectTypes[T GenericType](v T) (Type, any) {
	var _type = SERIES_TYPE_STRING
	vv := reflect.ValueOf(v)
	vk := vv.Kind()
	switch vk {
	case reflect.Invalid:
		_type = SERIES_TYPE_INVAILD
	case reflect.Bool:
		_type = SERIES_TYPE_BOOL
	case reflect.Int64:
		_type = SERIES_TYPE_INT
	case reflect.Float64:
		_type = SERIES_TYPE_FLOAT
	case reflect.String:
		_type = SERIES_TYPE_STRING
	default:
		panic(fmt.Errorf("unknown type, %+v", v))
	}
	return _type, vv.Interface()
}

// Shift series切片, 使用可选的时间频率按所需的周期数移动索引
func Shift[T GenericType](s *Series, periods int, cbNan func() T) Series {
	var d Series
	d = clone(*s).(Series)
	if periods == 0 {
		return d
	}

	values := d.Values().([]T)

	var (
		naVals []T
		dst    []T
		src    []T
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
		naVals[i] = cbNan()
	}
	_ = naVals
	return d
}

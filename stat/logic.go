package stat

//// Gt 比较 v > x
//func Gt(v []DType, x any) []bool {
//	return __compare(v, x, vek.Gt)
//}
//
//// Gte 比较 v >= x
//func Gte(v []DType, x any) []bool {
//	return __compare(v, x, vek.Gte)
//}
//
//// Lt 比较 v < x
//func Lt(v []DType, x any) []bool {
//	return __compare(v, x, vek.Lt)
//}
//
//// Lte 比较 v <= x
//func Lte(v []DType, x any) []bool {
//	return __compare(v, x, vek.Lte)
//}
//
//// __compare 比较 v 和 x
//func __compare(v []DType, x any, comparator func(x, y []float64) []bool) []bool {
//	vlen := len(v)
//
//	// 处理默认值
//	defaultValue := DType(0)
//	var X []DType
//	switch vx := x.(type) {
//	case int:
//		X = Repeat[DType](DType(vx), vlen)
//	case []DType:
//		xlen := len(vx)
//		if vlen < xlen {
//			vlen = xlen
//		}
//		X = Align[DType](vx, defaultValue, vlen)
//	case Series:
//		vs := vx.DTypes()
//		xlen := len(vs)
//		if vlen < xlen {
//			vlen = xlen
//		}
//		X = Align(vs, defaultValue, vlen)
//	default:
//		panic(exception.New(1, "error window"))
//	}
//	return comparator(v, X)
//}

func __compare[T ~[]E, E any](x T, y any, comparator func(f1, f2 DType) bool) []bool {
	var d = []bool{}
	switch Y := y.(type) {
	case nil, int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, float32, float64, bool, string:
		f2 := Any2DType(Y)
		d = __compare_dtype(x, f2, comparator)
	case []float32:
		d = __compare_slice(x, Y, comparator)
	case []float64:
		d = __compare_slice(x, Y, comparator)
	case []int:
		d = __compare_slice(x, Y, comparator)
	case []int8:
		d = __compare_slice(x, Y, comparator)
	case []int16:
		d = __compare_slice(x, Y, comparator)
	case []int32:
		d = __compare_slice(x, Y, comparator)
	case []int64:
		d = __compare_slice(x, Y, comparator)
	case []uint:
		d = __compare_slice(x, Y, comparator)
	case []uint8:
		d = __compare_slice(x, Y, comparator)
	case []uint16:
		d = __compare_slice(x, Y, comparator)
	case []uint32:
		d = __compare_slice(x, Y, comparator)
	case []uint64:
		d = __compare_slice(x, Y, comparator)
	case []uintptr:
		d = __compare_slice(x, Y, comparator)
	case []string:
		d = __compare_slice(x, Y, comparator)
	case []bool:
		d = __compare_slice(x, Y, comparator)
	default:
		// 其它类型原样返回
		panic(Throw(any(x)))
	}
	return d
}

func __compare_dtype[T ~[]E, E any](x T, y DType, comparator func(f1, f2 DType) bool) []bool {
	var bs = []bool{}
	xLen := len(x)
	// b不是切片
	bs = make([]bool, xLen)
	for i := 0; i < xLen; i++ {
		A := Any2DType(x[i])
		bs[i] = comparator(A, y)
	}
	return bs
}

func __compare_slice[T ~[]E, E any, T2 ~[]E2, E2 any](x T, y T2, comparator func(f1, f2 DType) bool) []bool {
	var bs = []bool{}
	xLen := len(x)
	// b不是切片
	bs = make([]bool, xLen)
	yLen := len(y)
	if xLen >= yLen {
		bs = make([]bool, xLen)
		for i := 0; i < yLen; i++ {
			f1 := Any2DType(x[i])
			f2 := Any2DType(y[i])
			bs[i] = comparator(f1, f2)
		}
		for i := yLen; i < xLen; i++ {
			f1 := Any2DType(x[i])
			f2 := DType(0)
			bs[i] = comparator(f1, f2)
		}
	} else {
		bs = make([]bool, yLen)
		for i := 0; i < xLen; i++ {
			f1 := Any2DType(x[i])
			f2 := Any2DType(y[i])
			bs[i] = comparator(f1, f2)
		}
		for i := xLen; i < yLen; i++ {
			f1 := DType(0)
			f2 := Any2DType(y[i])
			bs[i] = comparator(f1, f2)
		}
	}
	return bs
}

var (
	__logic_gt = func(f1, f2 DType) bool {
		return f1 > f2
	}

	__logic_gte = func(f1, f2 DType) bool {
		return f1 >= f2
	}

	__logic_lt = func(f1, f2 DType) bool {
		return f1 < f2
	}
	__logic_lte = func(f1, f2 DType) bool {
		return f1 <= f2
	}
)

// Gt 比较 v > x
func Gt[S ~[]E, E any](v S, x any) []bool {
	return __compare(v, x, __logic_gt)
}

// Gte 比较 v >= x
func Gte[S ~[]E, E any](v S, x any) []bool {
	return __compare(v, x, __logic_gte)
}

// Lt 比较 v < x
func Lt[S ~[]E, E any](v S, x any) []bool {
	return __compare(v, x, __logic_lt)
}

// Lte 比较 v <= x
func Lte[S ~[]E, E any](v S, x any) []bool {
	return __compare(v, x, __logic_lte)
}

package pandas

import (
	"fmt"
	"testing"
)

func TestDataFrameT0(t *testing.T) {
	var s1 Series
	s1 = NewSeriesFloat64("sales", nil, 50.3, 23.4, 56.2)
	fmt.Println(s1)
	expected := 4

	if s1.Len() != expected {
		t.Errorf("wrong val: expected: %v actual: %v", expected, s1.Len())
	}
	s2 := s1.Shift(2)
	df := NewDataFrame(s1, *s2)
	fmt.Println(df)

	_ = s2
}

func TestLoadStructs(t *testing.T) {
	type testStruct struct {
		A string
		B int
		C bool
		D float64
	}
	type testStructTags struct {
		A string  `dataframe:"a,string"`
		B int     `dataframe:"b,string"`
		C bool    `dataframe:"c,string"`
		D float64 `dataframe:"d,string"`
		E int     `dataframe:"-"` // ignored
		f int     // ignored
	}
	data := []testStruct{
		{"a", 1, true, 0.0},
		{"b", 2, true, 0.5},
	}
	dataTags := []testStructTags{
		{"a", 1, true, 0.0, 0, 0},
		{"NA", 2, true, 0.5, 0, 0},
	}

	df1 := LoadStructs(data)
	fmt.Println(df1)
	df2 := LoadStructs(dataTags)
	fmt.Println(df2)
}

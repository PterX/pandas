package tests

import (
	"encoding/json"
	"fmt"
	"gitee.com/quant1x/pandas/dataframe"
	"gitee.com/quant1x/pandas/series"
	"strings"
	"testing"
)

func TestCsv(t *testing.T) {
	csvStr := `
Country,Date,Age,Amount,Id
"United States",2012-02-01,50,112.1,01234
"United States",2012-02-01,32,321.31,54320
"United Kingdom",2012-02-01,17,18.2,12345
"United States",2012-02-01,32,321.31,54320
"United Kingdom",2012-02-01,NA,18.2,12345
"United States",2012-02-01,32,321.31,54320
"United States",2012-02-01,32,321.31,54320
Spain,2012-02-01,66,555.42,00241
`
	df := dataframe.ReadCSV(strings.NewReader(csvStr))
	fmt.Println(df)
	df.SetNames("a", "b", "c", "d", "e")
	s1 := df.Col("d")
	fmt.Println(s1)

	closes := df.Col("d")
	ma5 := closes.Rolling(5).Mean()
	series.New(closes, series.Float, "")
	fmt.Println(ma5)
}

type T1 struct {
	X []int64 `json:"x"`
}

func TestEwm(t *testing.T) {
	//a := make(map[string][]int, 8)
	t01 := map[string]int64{
		"x": 1,
	}
	fmt.Println(t01)
	t02 := map[string][]int64{
		"x": {1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	fmt.Println(t02)
	text := `{"x":[1,2,3,4,5,6,7,8,9]}`
	reader := strings.NewReader(text)
	parser := json.NewDecoder(reader)
	var t1 T1
	a1 := parser.Decode(&t1)
	fmt.Println(a1, t1)
	var t2 map[string][]int
	a2 := parser.Decode(&t2)
	fmt.Println(a2, t2)
	df := dataframe.ReadJSON(reader)
	fmt.Println(df)
	values := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := series.New(values, series.Int, "x")
	df = dataframe.New(s1)
	fmt.Println(df)
	xs := df.Col("x")
	r1 := xs.Rolling(5).Mean()
	fmt.Println(r1)

	e1 := xs.EWM(series.Alpha{Span: 5, At: series.AlphaSpan}, false, false).Mean()
	fmt.Println(e1)

	df1 := dataframe.New(e1)
	fmt.Println(df1)

	e2 := xs.EWM(series.Alpha{Span: 5, At: series.AlphaSpan}, true, false).Mean()
	fmt.Println(e2)

	df2 := dataframe.New(e1, e2)
	fmt.Println(df2)
}

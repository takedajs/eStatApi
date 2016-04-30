package main

import (
	"strconv"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
	"github.com/gonum/plot/vg"

	"./estatapi"
)

func main() {

	resp := estatapi.Get()
	values := resp.GET_STATS_DATA.STATISTICAL_DATA.DATA_INF.VALUE
	m, f := format(values)

	male := plotter.Values{m["0"], m["20"], m["40"], m["60"], m["80"], m["100"]}
	female := plotter.Values{f["0"], f["20"], f["40"], f["60"], f["80"], f["100"]}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "male&female"
	p.X.Label.Text = "old"
	p.Y.Label.Text = "number(thousand)"

	w := vg.Points(20)

	maleBars, err := plotter.NewBarChart(male, w)
	if err != nil {
		panic(err)
	}
	maleBars.LineStyle.Width = vg.Length(0)
	maleBars.Color = plotutil.Color(0)
	maleBars.Offset = -w

	femaleBars, err := plotter.NewBarChart(female, w)
	if err != nil {
		panic(err)
	}
	femaleBars.LineStyle.Width = vg.Length(0)
	femaleBars.Color = plotutil.Color(1)

	p.Add(maleBars, femaleBars)
	p.Legend.Add("male", maleBars)
	p.Legend.Add("female", femaleBars)
	p.Legend.Top = true
	p.Legend.Left = true
	p.NominalX("0", "20", "40", "60", "80", "100over")

	if err := p.Save(5*vg.Inch, 3*vg.Inch, "graf.png"); err != nil {
		panic(err)
	}
}

func format(values []estatapi.Value) (map[string]float64, map[string]float64) {
	male := make(map[string]float64)
	female := make(map[string]float64)
	for _, v := range values {
		if v.Cat02 == "001" {
			if v.Cat01 == "002" {
				switch v.Cat03 {
				case "01001":
					male["0"], _ = strconv.ParseFloat(v.Dollar, 64)
				case "01021":
					male["20"], _ = strconv.ParseFloat(v.Dollar, 64)
				case "01041":
					male["40"], _ = strconv.ParseFloat(v.Dollar, 64)
				case "01061":
					male["60"], _ = strconv.ParseFloat(v.Dollar, 64)
				case "01081":
					male["80"], _ = strconv.ParseFloat(v.Dollar, 64)
				case "01101":
					male["100"], _ = strconv.ParseFloat(v.Dollar, 64)
				}
			} else if v.Cat01 == "003" {
				switch v.Cat03 {
				case "01001":
					female["0"], _ = strconv.ParseFloat(v.Dollar, 64)
				case "01021":
					female["20"], _ = strconv.ParseFloat(v.Dollar, 64)
				case "01041":
					female["40"], _ = strconv.ParseFloat(v.Dollar, 64)
				case "01061":
					female["60"], _ = strconv.ParseFloat(v.Dollar, 64)
				case "01081":
					female["80"], _ = strconv.ParseFloat(v.Dollar, 64)
				case "01101":
					female["100"], _ = strconv.ParseFloat(v.Dollar, 64)
				}
			}
		}
	}
	return male, female
}

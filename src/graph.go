package src

import (
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"os"
	"time"
)

// GraphableProbability is the primary interface for graphing a
// set of probability data. It requires a graph title, a list of
// time.Time s for the X-axis and a list of one or more lists of
// YAxisItem s for the Y-axis
type GraphableProbability interface {
	Title() string
	XAxis() []time.Time
	YAxis() [][]YAxisItem
}

// YAxisItem is an interface that allows us to graph a
// named time-series in the graph. It requires a label
// and a `0.00` - `1.00` value.
type YAxisItem interface {
	Label() string
	Value() float64
}

// Graph takes a GraphableProbability and graphs it, saving the resulting file in
// the outputFilename arg. This produces a line graph of probabilities vs. dates
func Graph(probability GraphableProbability, outputFilename string) {
	// get the x and y-axis data
	yAxisData := lineItems(probability)

	// get the dates a little nicer format
	xAxisData := XAxisNice(probability.XAxis())

	// create a line chart with the probability's title
	bar := charts.NewLine()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: probability.Title(),
	}))

	// set the x axis using the nice dates
	bar.SetXAxis(xAxisData)

	// add all of the series info from the y-axis data
	for _, yAxisItem := range yAxisData {
		bar.AddSeries(yAxisItem.label, yAxisItem.lineData)
	}

	// create the output file
	f, err := os.Create(outputFilename)
	if err != nil {
		panic(err)
	}

	// render the graph
	err = bar.Render(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Wrote '%s' graph to file '%s'\n", probability.Title(), outputFilename)
}

// XAxisNice converts a list of times into the YYYY-MM-DD format
func XAxisNice(times []time.Time) []string {
	layout := "2006-01-02"
	output := make([]string, 0)
	for _, v := range times {
		// format the time.Time as YYYY-MM-DD
		output = append(output, v.Format(layout))
	}
	return output
}

// yAxis is a struct that converts a list of YAxisItem into the format
// that go-echarts wants for a data series
type yAxis struct {
	label    string
	lineData []opts.LineData
}

// lineItems converts the Y-axis items from the GraphableProbability
// into a list of yAxis items that can be graphed via go-echarts
func lineItems(probability GraphableProbability) []yAxis {
	output := make([]yAxis, 0)
	// for each item in the
	for _, x := range probability.YAxis() {

		// write the label (all values are the same in this YAxisItem list)
		data := yAxis{label: x[0].Label()}

		// for each item in the Y-axis data, convert it into a opts.LineData value
		lineData := make([]opts.LineData, 0)
		for _, y := range x {
			lineData = append(lineData, opts.LineData{Value: y.Value()})
		}

		// save the line data to the series
		data.lineData = lineData

		// save the series to the list
		output = append(output, data)
	}

	return output
}

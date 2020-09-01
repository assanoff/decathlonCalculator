package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type place struct {
	name string
}

type byPoints []*Athlete

//NewPlace ...
func NewPlace() *place {
	return &place{}
}

func (x byPoints) Len() int           { return len(x) }
func (x byPoints) Less(i, j int) bool { return x[j].total < x[i].total }
func (x byPoints) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func printResult(athlets []*Athlete) {
	const format = "%v\t%v\t%v\t\n"
	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 3, ' ', tabwriter.TabIndent)
	fmt.Fprintf(tw, format, "Athlete", "Result", "Place")
	fmt.Fprintf(tw, format, "-------", "-------", "-------")
	for _, t := range athlets {
		fmt.Fprintf(tw, format, t.name, t.total, t.place.name)
	}
	tw.Flush()
}

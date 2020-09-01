package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

func main() {
	sportEvents = initEvents()
	athletes := []*Athlete{}
	f, _ := os.Open("result.csv")
	defer f.Close()
	r := csv.NewReader(bufio.NewReader(f))
	r.Comma = ';'
	r.Comment = '#'
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		a, err := NewAthlete(line[0], line[1:])
		if err != nil {
			fmt.Println("Could not create Athlete", err)
			return
		}
		athletes = append(athletes, a)
	}
	sort.Sort(byPoints(athletes))

	var pl = NewPlace()
	for i, v := range athletes {
		pl.name += fmt.Sprintf("%d", i+1)
		v.place = pl
		if (i+1) > len(athletes)-1 || athletes[i].total != athletes[i+1].total {
			pl = NewPlace()
		}
	}

	printResult(athletes)
}

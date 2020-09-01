package main

import (
	"strconv"
	"strings"
)

type Athlete struct {
	name   string
	total  int
	points []Point
	place  *place
}

type Point struct {
	id    int
	point float64
}

// NewPoint ...
func NewPoint(id int, p float64) Point {
	return Point{
		id:    id,
		point: p,
	}
}

// NewAthlete ...
func NewAthlete(n string, scores []string) (*Athlete, error) {
	points := []Point{}
	result := 0
	for id, p := range scores {
		e := sportEvents[id]
		athletePoint, err := Convert(p, e.k)
		if err != nil {
			return nil, err
		}
		point := NewPoint(id, athletePoint)
		result += e.calc(e, point.point)
		points = append(points, point)
	}
	return &Athlete{
		name:   n,
		points: points,
		total:  result,
	}, nil
}

// Convert ...
func Convert(s string, k int) (float64, error) {
	var min float64
	var sec float64
	var err error
	sl := strings.Split(s, ":")
	if len(sl) == 2 {
		min, err = strconv.ParseFloat(sl[0], 64)
		if err != nil {
			return 0.0, err
		}
	}
	if sec, err = strconv.ParseFloat(sl[len(sl)-1], 64); err != nil {
		return 0.0, err
	}
	return (min*60.0 + sec) * float64(k), nil
}

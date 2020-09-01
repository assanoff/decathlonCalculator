package main

import "math"

const (
	m100 int = iota
	longJump
	shotPut
	highJump
	m400
	m110
	discus
	poleVault
	javelin
	m1500
)

var sportEvents []Event

// Event ...
type Event struct {
	name int
	calc func(Event, float64) int
	A    float64
	B    float64
	C    float64
	k    int
}

func initEvents() []Event {
	return []Event{
		NewEvent(m100, calcTrack, 25.4347, 18, 1.81, 1),
		NewEvent(longJump, calcDistance, 0.14354, 220, 1.4, 100),
		NewEvent(shotPut, calcDistance, 51.39, 1.5, 1.05, 1),
		NewEvent(highJump, calcDistance, 0.8465, 75, 1.42, 100),
		NewEvent(m400, calcTrack, 1.53775, 82, 1.81, 1),
		NewEvent(m110, calcTrack, 5.74352, 28.5, 1.92, 1),
		NewEvent(discus, calcDistance, 12.91, 4, 1.1, 1),
		NewEvent(poleVault, calcDistance, 0.2797, 100, 1.35, 100),
		NewEvent(javelin, calcDistance, 10.14, 7, 1.08, 1),
		NewEvent(m1500, calcTrack, 0.03768, 480, 1.85, 1),
	}
}

// NewEvent ...
func NewEvent(name int, calcFunc func(Event, float64) int, a, b, c float64, k int) Event {
	return Event{
		name: name,
		calc: calcFunc,
		A:    a,
		B:    b,
		C:    c,
		k:    k,
	}
}

// // Points = INT(A(B — P)C) for track events (faster time produces a better score)
func calcTrack(e Event, p float64) int {
	return int(e.A * (math.Pow((e.B - p), e.C)))
}

// Points = INT(A(P — B)C) for field events (greater distance or height produces a better score)
func calcDistance(e Event, p float64) int {
	return int(e.A * (math.Pow((p - e.B), e.C)))
}

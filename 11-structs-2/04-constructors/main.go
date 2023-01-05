package main

import (
	"errors"
	"fmt"
	"time"
)

type DateRange struct {
	start time.Time
	end   time.Time
}

func (d DateRange) Hours() float64 {
	return d.end.Sub(d.start).Hours()
}

func NewDateRange(start time.Time, end time.Time) (DateRange, error) {
	if start.IsZero() {
		return DateRange{}, errors.New("empty start")
	}
	if end.IsZero() {
		return DateRange{}, errors.New("empty end")
	}
	if end.Before(start) {
		return DateRange{}, errors.New("end before start")
	}

	return DateRange{
		start: start,
		end:   end,
	}, nil
}

func main() {
	lifetime, err := NewDateRange(
		time.Date(1815, 12, 10, 0, 0, 0, 0, time.UTC),
		time.Date(1852, 11, 27, 0, 0, 0, 0, time.UTC),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(lifetime.Hours())

	travelInTime, err := NewDateRange(
		time.Date(1852, 11, 27, 0, 0, 0, 0, time.UTC),
		time.Date(1815, 12, 10, 0, 0, 0, 0, time.UTC),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(travelInTime.Hours())
}

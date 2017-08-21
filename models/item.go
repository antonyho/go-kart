package models

import "time"

type Item struct {
	Name         string
	Description  string
	Price        float64
	Availability uint
	Begin        time.Time
	End          time.Time
	//Listed       bool
}

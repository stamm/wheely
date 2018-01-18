package types

import "time"

type CalculateRequest struct {
	StartLat  float64 `json:"start_lat"`
	StartLong float64 `json:"start_long"`
	EndLat    float64 `json:"end_lat"`
	EndLong   float64 `json:"end_long"`
}

type CalculateResponse struct {
	Distance int64  `json:"distance"`
	Duration int64  `json:"duration"`
	Err      string `json:"err,omitempty"`
}

type Result struct {
	Distance int64
	Duration time.Duration
}

type Point struct {
	Lat  float64
	Long float64
}

func NewPoint(lat, long float64) Point {
	return Point{Lat: lat, Long: long}
}

type Travel struct {
	Start Point
	End   Point
}

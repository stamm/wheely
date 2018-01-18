package distance

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

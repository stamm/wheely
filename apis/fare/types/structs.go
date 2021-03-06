package types

type CalculateRequest struct {
	StartLat  float64 `json:"start_lat"`
	StartLong float64 `json:"start_long"`
	EndLat    float64 `json:"end_lat"`
	EndLong   float64 `json:"end_long"`
}

type CalculateResponse struct {
	Amount int64  `json:"amount"`
	Err    string `json:"err,omitempty"`
}

type Result struct {
	Amount int64
}

type Tariff struct {
	Delivery  int64 `json:"delivery"`
	ForMinute int64 `json:"for_minute"`
	ForKm     int64 `json:"for_km"`
	Minimum   int64 `json:"minimum"`
}

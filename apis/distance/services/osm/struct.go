package osm

type response struct {
	Routes []route `json:"routes"`
}

type route struct {
	Summary summary `json:"summary"`
}

type summary struct {
	// Distance indicates the total distance covered by this leg.
	Distance float64 `json:"distance"`
	// Duration indicates total time required for this leg.
	Duration float64 `json:"duration"`
}

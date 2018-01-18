package google

import "time"

type response struct {
	Routes []Route `json:"routes"`
	Status string  `json:"status"`
}

type Route struct {

	// Legs contains information about a leg of the route, between two locations within the
	// given route. A separate leg will be present for each waypoint or destination specified.
	// (A route with no waypoints will contain exactly one leg within the legs array.)
	Legs []leg `json:"legs"`
}

type leg struct {

	// Distance indicates the total distance covered by this leg.
	Distance distance `json:"distance"`

	// Duration indicates total time required for this leg.
	Duration duration `json:"duration"`
}

type distance struct {
	// Meters is the numeric distance, always in meters. This is intended to be used only in
	// algorithmic situations, e.g. sorting results by some user specified metric.
	Meters int64 `json:"value"`
}

type duration struct {
	// Seconds indicates the duration, in seconds.
	Seconds int64 `json:"value"`
}

// Duration returns the time.Duration for this internal Duration.
func (d *duration) Duration() time.Duration {
	if d == nil {
		return time.Duration(0)
	}
	return time.Duration(d.Seconds) * time.Second
}

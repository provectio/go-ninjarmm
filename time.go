package ninjarmm

import (
	"encoding/json"
	"time"
)

// Better implementation of `double“ provided by NinjaAPI for time.Time.
type Time time.Time

// UnmarshalJSON implements the json.Unmarshaler interface for ninjarmm.Time.
func (t *Time) UnmarshalJSON(b []byte) (err error) {
	var double float64
	if err = json.Unmarshal(b, &double); err != nil {
		return
	}

	var parsed time.Time
	if double > 0 {
		parsed = time.Unix(int64(double), int64(double*1e9)%1e9)
	} else {
		parsed = time.Time{}
	}

	*t = Time(parsed)
	return
}

// MarshalJSON implements the json.Marshaler interface for ninjarmm.Time.
func (j Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j))
}

// Format implements the fmt.Formatter interface for ninjarmm.Time.
func (j Time) Format(s string) string {
	return time.Time(j).Format(s)
}

// String implements the fmt.Stringer interface for ninjarmm.Time.
func (j Time) String() string {
	return time.Time(j).String()
}

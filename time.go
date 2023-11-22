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
	var double float64
	if !time.Time(j).IsZero() {
		double = float64(time.Time(j).UnixNano()) / 1e9
	}
	return json.Marshal(double)
}

// Format implements the fmt.Formatter interface for ninjarmm.Time.
func (j Time) Format(s string) string {
	return time.Time(j).Format(s)
}

// String implements the fmt.Stringer interface for ninjarmm.Time.
func (j Time) String() string {
	return time.Time(j).String()
}

// Add implements the time.Time.Add method for ninjarmm.Time.
func (j Time) Add(d time.Duration) Time {
	return Time(time.Time(j).Add(d))
}

// After implements the time.Time.After method for ninjarmm.Time.
func (j Time) After(u time.Time) bool {
	return time.Time(j).After(u)
}

// AppendFormat implements the time.Time.AppendFormat method for ninjarmm.Time.
func (j Time) AppendFormat(b []byte, f string) []byte {
	return time.Time(j).AppendFormat(b, f)
}

// Before implements the time.Time.Before method for ninjarmm.Time.
func (j Time) Before(u time.Time) bool {
	return time.Time(j).Before(u)
}

// Clock implements the time.Time.Clock method for ninjarmm.Time.
func (j Time) Clock() (int, int, int) {
	return time.Time(j).Clock()
}

// Date implements the time.Time.Date method for ninjarmm.Time.
func (j Time) Date() (int, time.Month, int) {
	return time.Time(j).Date()
}

// Sub implements the time.Time.Sub method for ninjarmm.Time.
func (j Time) Sub(u time.Time) time.Duration {
	return time.Time(j).Sub(u)
}

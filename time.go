package ninjarmm

import (
	"encoding/json"
	"time"
)

type Time time.Time

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

func (j Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j))
}

func (j Time) Format(s string) string {
	return time.Time(j).Format(s)
}

func (j Time) String() string {
	return time.Time(j).String()
}

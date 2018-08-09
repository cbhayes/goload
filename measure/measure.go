package measure

import (
	"fmt"
	"time"
)

type Measure struct {
	Elapsed time.Duration
	Status	int
	Body	[]byte
}

func Create(start time.Time, end time.Time, status int, body []byte) Measure {
	return Measure{Elapsed: end.Sub(start), Status: status, Body: body}
}

func (measure Measure) Format() string {
	d := measure.Elapsed
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	d -= s * time.Second
	return fmt.Sprint("Duration: %d:%d:%d:%d", h, m, s, d)
}
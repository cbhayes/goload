package measure

import (
	"fmt"
	"time"

	"github.com/jamiealquiza/tachymeter"
)

type Measure struct {
	Elapsed time.Duration
	Status  int
	Body    []byte
}

func Create(start time.Time, end time.Time, status int, body []byte) Measure {
	return Measure{Elapsed: end.Sub(start), Status: status, Body: body}
}

func Stats(measures *[]Measure) {
	t := tachymeter.New(&tachymeter.Config{Size: len(*measures)})

	for _, measure := range *measures {
		t.AddTime(measure.Elapsed)
	}
	
	fmt.Println(t.Calc().String())
	fmt.Println(t.Calc().Histogram.String(25))
}

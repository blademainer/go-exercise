package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

const format = "2006-01-02 15:04:05.999"

type S struct {
	CreatedTime JsonTime
}

type JsonTime time.Time

// MarshalJSON implements the json.Marshaler interface.
// The time is a quoted string in RFC 3339 format, with sub-second precision added if present.
func (f JsonTime) MarshalJSON() ([]byte, error) {
	t := time.Time(f)

	if y := t.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(format)+2)
	b = append(b, '"')
	b = t.AppendFormat(b, format)
	b = append(b, '"')
	return b, nil
}

func main() {
	s := &S{
		CreatedTime: JsonTime(time.Now()),
	}
	marshal, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshal))
}

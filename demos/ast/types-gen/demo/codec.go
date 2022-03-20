package codec

import (
	"reflect"
)

// Codec defines the interface services uses to encode and decode messages.  Note
// that implementations of this interface must be thread safe; a Codec's
// methods can be called from concurrent goroutines.
type Codec interface {
	// Marshal returns the wire format of v.
	Marshal(v interface{}) (string, error)
	// Unmarshal parses the wire format into v.
	Unmarshal(data string, v interface{}) error
}

type Coders struct {
}

func NewCodec() Coders {
	return Coders{}
}

func (c Coders) Funcs() []string {
	tc := reflect.TypeOf(c)
	rs := make([]string, 0, tc.NumMethod()-1)
	for i := 0; i < tc.NumMethod(); i++ {
		m := tc.Method(i)
		if m.Name == "Funcs" {
			continue
		}
		rs = append(rs, m.Name)
	}
	return rs
}

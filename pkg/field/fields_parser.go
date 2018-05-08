package field

type (
	// Parser is the default implementation of the Binder interface.
	Parser struct{
		Tag string
		EscapeHTML bool
		Quoted bool
	}

	// Unmarshaler is the interface used to wrap the UnmarshalParam method.
	Unmarshaler interface {
		// UnmarshalParam decodes and assigns a value from an form or query param.
		UnmarshalParam(param string) error
	}

	Marshaler interface {
		Marshal() (string, error)
	}
)

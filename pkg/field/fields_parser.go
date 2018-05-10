package field

type (
	// Parser is the default implementation of the Binder interface.
	Parser struct{
		Tag    string
		Escape bool
		Quoted bool
		// Delimiter between groups. For example: we should convert struct to http form, so this GroupDelimiter is '&' and PairDelimiter is '='
		GroupDelimiter byte
		// Delimiter between key and value.
		PairDelimiter byte
		// Field should sort by field name
		Sort bool
		// Ignore fields that value is nil
		IgnoreNilValueField bool
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

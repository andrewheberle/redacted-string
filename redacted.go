package redacted

// Redacted can be used to print a redacted string using fmt.Println(Redacted("this will be printed as all *"))
type Redacted string

const RedactedCharacter = '*'

func (r Redacted) String() string {
	b := make([]byte, 0)
	for range r {
		b = append(b, RedactedCharacter)
	}

	return string(b)
}

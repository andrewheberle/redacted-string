package redacted

import (
	"strings"
)

// Censor allows you to hide sensitive values. This should be created using New rather than directly.
type Censor struct {
	s                 string
	redactedCharacter rune
	redactedLength    uint
}

const (
	// The default character to use in place of real characters in a redacted string for a Censor created via New.
	RedactedCharacter rune = '*'
	// The default maximum length of the redacted string for a Censor created via New. A value of zero (0) means the redacted string will be the same length as the original.
	RedactedLength uint = 8
)

func (c Censor) String() string {
	// use defaults if not initialised
	return redact(c.s, c.redactedCharacter, c.redactedLength)
}

// Set changes (or sets) the string to be redacted
func (c *Censor) Set(s string) {
	c.s = s
}

func redact(s string, c rune, size uint) string {
	// redact string
	var b strings.Builder

	max := int(size)
	if max == 0 {
		max = len(s)
	}
	for n := 0; n < max; n++ {
		b.WriteRune(c)
	}

	return b.String()
}

// New creates a new Censor that can be used to redact the provided string
func New(s string, opts ...Option) *Censor {
	censor := new(Censor)

	censor.s = s
	censor.redactedCharacter = RedactedCharacter
	censor.redactedLength = RedactedLength

	for _, o := range opts {
		o(censor)
	}

	return censor
}

// Redact can be used to print a redacted string (ie printed as all '********').
//
// It is also possible to pass options to customise the behaviour of the redaction in the same way as New accepts options.
func Redact(s string, opts ...Option) string {
	return New(s, opts...).String()
}

// The Redacted function is to provide backwards compatibility to the previously documented API
func Redacted(s string) string {
	return redact(s, RedactedCharacter, RedactedLength)
}

// Option type to change the behaviour of a Censor
type Option func(*Censor)

// WithCharacter sets the character
func WithCharacter(char rune) Option {
	return func(c *Censor) {
		c.redactedCharacter = char
	}
}

// WithLength sets the length of the redacted string. A value of zero means the redacted version is the same length as the original.
func WithLength(length uint) Option {
	return func(c *Censor) {
		c.redactedLength = length
	}
}

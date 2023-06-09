package str

import "github.com/zxcfer/xtring/parsers"

// ParseFunc is a function that attempts to parse the
// string into a real type and returns whether that was
// successful or not.
type ParseFunc func(s string) (interface{}, bool)

// Parser is a slice of ParseFunc functions that will be
// called in order.
type Parser []ParseFunc

// New creates a new Parser with the specified ParseFunc
// functions. Each function will be tried in order until
// one is able to parse the strings.
// If none can, the original string is returned untouched.
func New(funcs ...ParseFunc) Parser {
	return Parser(funcs)
}

// Parse parses s with each ParseFunc in order returning
// the result.
// If no parsers are successful, returns the string untouched.
func (p Parser) Parse(s string) interface{} {
	v, _ := p.parseWith(s)
	return v
}

// ParseWith parses s with each ParseFunc in order returning
// the result, and returns the ParseFunc that was successful.
// If no parsers are successful, returns the string untouched
// and the second argument will be nil.
func (p Parser) parseWith(s string) (interface{}, ParseFunc) {
	for _, try := range p {
		if val, ok := try(s); ok {
			return val, try
		}
	}
	return s, nil
}

// Parse parses s with the DefaultParser.
// For more information, see Parser.Parse.
func Parse(s string) interface{} {
	return DefaultParser.Parse(s)
}

// DefaultParser is the default Parser that includes
// all built-in ParseFunc functions.
var DefaultParser = Parser([]ParseFunc{
	parsers.Quoted,
	parsers.Nil,
	parsers.Null,
	parsers.Bool,
	parsers.Int,
	parsers.Int64,
	parsers.UInt,
	parsers.Uint64,
	parsers.Float64,
})

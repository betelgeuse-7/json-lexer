package main

type Token int

const (
	EOF Token = iota
	UNKNOWN
	WHITESPACE
	NUMBER

	IDENT

	// KEYWORDS
	TRUE  // true
	FALSE // false
	NULL  // null

	// SPECIAL SYMBOLS
	BEGIN_ARRAY     // [
	BEGIN_OBJECT    // {
	END_ARRAY       // ]
	END_OBJECT      // }
	NAME_SEPARATOR  // :
	VALUE_SEPARATOR // ,
	ESCAPE          // \
	QUOTATION_MARK  // "
)

func (t Token) String() string {
	return map[Token]string{
		EOF:             "EOF",
		UNKNOWN:         "UNKNOWN",
		WHITESPACE:      "WHITESPACE",
		NUMBER:          "NUMBER",
		IDENT:           "IDENT",
		TRUE:            "TRUE",
		FALSE:           "FALSE",
		NULL:            "NULL",
		BEGIN_ARRAY:     "BEGIN_ARRAY",
		BEGIN_OBJECT:    "BEGIN_OBJECT",
		END_ARRAY:       "END_ARRAY",
		END_OBJECT:      "END_OBJECT",
		NAME_SEPARATOR:  "NAME_SEPARATOR",
		VALUE_SEPARATOR: "VALUE_SEPARATOR",
		ESCAPE:          "ESCAPE",
		QUOTATION_MARK:  "QUOTATION_MARK",
	}[t]
}

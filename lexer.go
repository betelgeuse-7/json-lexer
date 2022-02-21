package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"unicode"
)

const EOF_RUNE = rune(0)

type Position struct {
	line   int
	column int
}

type Lexer struct {
	pos Position
	r   *bufio.Reader
}

// reset column and increment line by one
func (l *Lexer) resetPos() {
	l.pos.column = 0
	l.pos.line++
}

// TOKEN (used for JSON method)
// ????
type T struct {
	Type    string `json:"type"`
	Literal string `json:"literal"`
	Line    int    `json:"line"`
	Column  int    `json:"column"`
}

func NewLexer(r io.Reader) *Lexer {
	return &Lexer{
		r: bufio.NewReader(r),
		pos: Position{
			line:   1,
			column: 0,
		},
	}
}

// read one rune from l.r and return it
func (l *Lexer) read() rune {
	r, _, err := l.r.ReadRune()
	if err != nil {
		return EOF_RUNE
	}
	return r
}

// put the rune back in l.r (*bufio.Reader)
func (l *Lexer) unread() {
	_ = l.r.UnreadRune()
}

func (l *Lexer) lexWhitespace() (pos Position, tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(l.read())

	for {
		ch := l.read()

		if ch == EOF_RUNE {
			break
		}
		if !isWhitespace(ch) {
			l.unread()
			break
		}
		buf.WriteRune(ch)
	}

	return l.pos, WHITESPACE, buf.String()
}

func (l *Lexer) lexIdent() (pos Position, tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(l.read())

	for {
		ch := l.read()

		if ch == EOF_RUNE {
			break
		}
		if !unicode.IsNumber(ch) && !isLetter(ch) && ch != '_' {
			l.unread()
			break
		}
		_, err := buf.WriteRune(ch)
		if err != nil {
			log.Printf("lexIdent err: %s\n", err.Error())
		}
	}

	bufStr := buf.String()

	switch bufStr {
	case "true":
		return l.pos, TRUE, "true"
	case "false":
		return l.pos, FALSE, "false"
	case "null":
		return l.pos, NULL, "null"
	default:
		return l.pos, IDENT, bufStr
	}
}

func (l *Lexer) lexNumber() (pos Position, tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(l.read())

	for {
		ch := l.read()

		if ch == EOF_RUNE {
			break
		} else if !unicode.IsNumber(ch) {
			l.unread()
			break
		}
		buf.WriteRune(ch)
	}

	return l.pos, NUMBER, buf.String()
}

// scan l.r's buffer and lex it
func (l *Lexer) lex() (pos Position, tok Token, lit string) {
	ch := l.read()
	l.pos.column++

	if isWhitespace(ch) {
		if ch == '\n' {
			l.resetPos()
		}
		l.unread()
		return l.lexWhitespace()
	} else if isLetter(ch) {
		l.unread()
		return l.lexIdent()
	} else if unicode.IsNumber(ch) {
		l.unread()
		return l.lexNumber()
	}

	// use naked returns ?
	// i don't think, i could.
	// anyways...

	switch ch {
	case EOF_RUNE:
		return l.pos, EOF, ""
	case ' ':
		return l.pos, WHITESPACE, " "
	case '[':
		return l.pos, BEGIN_ARRAY, "["
	case ']':
		return l.pos, END_ARRAY, "]"
	case '{':
		return l.pos, BEGIN_OBJECT, "{"
	case '}':
		return l.pos, END_OBJECT, "}"
	case ':':
		return l.pos, NAME_SEPARATOR, ":"
	case ',':
		return l.pos, VALUE_SEPARATOR, ","
	case '\\':
		return l.pos, ESCAPE, "\\"
	case '"':
		return l.pos, QUOTATION_MARK, "\""
	}

	return l.pos, UNKNOWN, string(ch)
}

// return JSON representation of lexed tokens
func (l *Lexer) JSON() ([]byte, error) {
	tt := []T{}

	for {
		t := T{}
		pos, tok, lit := l.lex()
		if tok == EOF {
			break
		}
		t.Type = tok.String()
		t.Literal = lit
		t.Column = pos.column
		t.Line = pos.line

		tt = append(tt, t)
	}

	bx, err := json.MarshalIndent(tt, "", "   ")
	if err != nil {
		return []byte{}, err
	}
	return bx, nil
}

func isWhitespace(ch rune) bool {
	return unicode.IsSpace(ch)
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || unicode.IsSymbol(ch)
}

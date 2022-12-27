package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current character)
	readPosition int  // position after current position (allows peeking at next character while evaluating position)
	ch           byte // current character being evaluated
}

// New returns a new lexer initialized with the provided input
func New(input string) *Lexer {
	l := &Lexer{input: input}
	// initialize the lexer state by reading in the first character of the input
	l.ReadChar()
	return l
}

// ReadChar reads the next character from the input into the lexer state
func (l *Lexer) ReadChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// NextToken returns the next token from the lexer's set of parsed tokens
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	}

	l.ReadChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

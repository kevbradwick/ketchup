package ketchup

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(ASSIGN, l.ch)
	case ';':
		tok = newToken(SEMICOLON, l.ch)
	case '(':
		tok = newToken(LPAREN, l.ch)
	case ')':
		tok = newToken(RPAREN, l.ch)
	case ',':
		tok = newToken(COMMA, l.ch)
	case '+':
		tok = newToken(PLUS, l.ch)
	case '{':
		tok = newToken(LBRACE, l.ch)
	case '}':
		tok = newToken(RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = INT
			return tok
		} else {
			tok = newToken(ILLEGAL, l.ch)
		}
	}

	// advance character
	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		l.readChar()
	}
}

func newToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

// returns the identifier which can be a string of chars including a-zA-Z and _
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	if ch >= '0' && ch <= '9' {
		return true
	}
	return false
}

// / loads the current character into position
func (l *Lexer) readChar() {
	// check if we are at the end of the string and if so, set the char to zero
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		// otherwise get the current character from the input string
		l.ch = l.input[l.readPosition]
	}

	// advance the current position
	l.position = l.readPosition
	l.readPosition += 1
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

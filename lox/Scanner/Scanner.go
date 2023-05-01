package Scanner

import (
	"Golox/lox/Token"
	"Golox/lox/errors"
	"unicode/utf8"
)

type Scanner struct {
	source  string
	tokens  []token.Token
	start   int
	current int
	line    int
}

func New(source string) Scanner {
	s := Scanner{source: source, line: 1}
	return s
}

func (s *Scanner) ScanTokens() ([]token.Token, error) {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.start = s.current
	s.addToken(token.EOF)

	return s.tokens, nil
}

func (s *Scanner) scanToken() {
	r := s.advance()
	switch r {
	case rune('('):
		s.addToken(token.LEFT_PAREN)
	case rune(')'):
		s.addToken(token.RIGHT_PAREN)
	case rune('{'):
		s.addToken(token.LEFT_BRACE)
	case rune('}'):
		s.addToken(token.RIGHT_BRACE)
	case rune(','):
		s.addToken(token.COMMA)
	case rune('.'):
		s.addToken(token.DOT)
	case rune('-'):
		s.addToken(token.MINUS)
	case rune('+'):
		s.addToken(token.PLUS)
	case rune(';'):
		s.addToken(token.SEMICOLON)
	case rune('*'):
		s.addToken(token.STAR)
	case rune('!'):
		if s.match('=') {
			s.addToken(token.BANG_EQUAL)
		} else {
			s.addToken(token.BANG)
		}
	case rune('='):
		if s.match('=') {
			s.addToken(token.EQUAL_EQUAL)
		} else {
			s.addToken(token.EQUAL)
		}
	case rune('<'):
		if s.match('=') {
			s.addToken(token.LESS_EQUAL)
		} else {
			s.addToken(token.LESS)
		}
	case rune('>'):
		if s.match('=') {
			s.addToken(token.GREATER_EQUAL)
		} else {
			s.addToken(token.GREATER)
		}
	case rune('/'):
		if s.match('/') {
			for !s.isAtEnd() && s.peek() != '\n' {
				s.advance()
			}
		} else {
			s.addToken(token.SLASH)
		}
	case rune(' '):
	case rune('\r'):
	case rune('\t'):
	case rune('\n'):
		s.line++
	case rune('"'):
		s.string()
	default:
		if isDigit(r) {
			s.number()
		} else if isAlpha(r) {
			s.identifier()
		} else {
			errors.Report(s.line, "", "Unexpected character.")
		}
	}
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}
func (s *Scanner) advance() rune {
	r, len := utf8.DecodeRuneInString(s.source[s.current:])
	s.current += len
	return r
}

func (s *Scanner) addToken(t token.TokenType) {
	s.addTokenLiteral(t, nil)
}

func (s *Scanner) addTokenLiteral(tokenType token.TokenType, literal any) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, token.Token{Type: tokenType, Lexeme: text, Literal: literal, Line: s.line})
}

func (s *Scanner) match(r rune) bool {
	if s.isAtEnd() {
		return false
	}

	ex, _ := utf8.DecodeRuneInString(s.source[s.current:])
	if ex != r {
		return false
	}

	s.advance()
	return true
}

func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return rune(0)
	}
	r, _ := utf8.DecodeRuneInString(s.source[s.current:])
	return r
}
func (s *Scanner) peekNext() rune {
	if s.current+1 >= len(s.source) {
		return rune(0)
	}

	r, _ := utf8.DecodeRuneInString(s.source[s.current+1:])
	return r
}

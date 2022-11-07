package main

import (
	"unicode"
	"unicode/utf8"
)

type lexer struct {
	sql string
	pos int
}

func (l *lexer) NextToken() token {
	return l.nextToken()
}

func (l *lexer) NextTokenAfterSpace() token {
	for {
		next := l.nextToken()
		if next.kind != tokenSpace {
			return next
		}
	}
}

func (l *lexer) PeekTokens(lookAhead uint) []token {
	var start = l.pos
	var toks []token
	for i := uint(0); i < lookAhead; i++ {
		toks = append(toks, l.nextToken())
	}
	l.pos = start
	return toks
}

func (l *lexer) PeekTokenAfterSpace() token {
	var start = l.pos
	for {
		tk := l.nextToken()
		if tk.kind != tokenSpace {
			l.pos = start
			return tk
		}
	}
}

func (l *lexer) nextRune() (rune, int) {
	if l.pos >= len(l.sql) {
		return -1, 0
	}
	r, size := utf8.DecodeRuneInString(l.sql[l.pos:])
	l.pos += size
	return r, size
}

func (l *lexer) peekRune() (rune, int) {
	if l.pos >= len(l.sql) {
		return -1, 0
	}
	r, size := utf8.DecodeRuneInString(l.sql[l.pos:])
	return r, size
}

func (l *lexer) nextToken() token {
	r, _ := l.peekRune()
	if r == -1 {
		return token{kind: tokenEnd}
	}

	switch {
	case unicode.IsDigit(r):
		return l.lexNumToken()
	case isAlphaNumeric(r):
		return l.lexIdentToken()
	case unicode.IsSpace(r):
		return l.lexSpaceToken()
	case r == '(' || r == ')':
		return l.lexParensToken()
	case r == '"':
		return l.lexStrToken()
	case r == ',':
		return l.lexDelToken()
	case r == '=':
		return l.lexEqToken()
	case r == '*':
		return l.lexStarToken()
	default:
		panic("not handled")
	}
}

func (l *lexer) lexNumToken() token {
	var num = make([]rune, 0, 16)
	for {
		r, size := l.nextRune()
		if r == -1 {
			break
		}
		if unicode.IsDigit(r) {
			num = append(num, r)
			continue
		}
		l.pos -= size
		break
	}
	return token{
		kind: tokenInt,
		val:  string(num),
	}
}

func (l *lexer) lexStrToken() token {
	r, _ := l.nextRune()
	assert(r, '"')

	var str = make([]rune, 0, 16)
	for {
		r, _ := l.nextRune()
		if r == -1 {
			panic(r)
		}
		if r != '"' {
			str = append(str, r)
			continue
		}

		if str[len(str)-1] == '\\' {
			// escaped
			str = append(str, r)
			continue
		}
		break
	}

	return token{
		kind: tokenStr,
		val:  string(str),
	}
}

func (l *lexer) lexIdentToken() token {
	var ident = make([]rune, 0, 16)
	for {
		r, size := l.nextRune()
		if r == -1 {
			break
		}
		if isAlphaNumeric(r) {
			ident = append(ident, r)
			continue
		}
		l.pos -= size
		break
	}

	val := string(ident)
	keywordType, ok := keywordsToToken[val]
	if ok {
		return token{kind: keywordType, val: val}
	}
	return token{kind: tokenIdent, val: val}
}

func (l *lexer) lexSpaceToken() token {
	for {
		r, size := l.nextRune()
		if r == -1 {
			break
		}
		if unicode.IsSpace(r) {
			continue
		}
		l.pos -= size
		break
	}
	return token{
		kind: tokenSpace,
		val:  "",
	}
}

func (l *lexer) lexParensToken() token {
	r, _ := l.nextRune()
	switch {
	case r == '(':
		return token{kind: tokenOpenPar, val: "("}
	case r == ')':
		return token{kind: tokenClosePar, val: ")"}
	default:
		panic("bug")
	}
}

func (l *lexer) lexDelToken() token {
	r, _ := l.nextRune()
	switch {
	case r == ',':
		return token{kind: tokenDel, val: ","}
	default:
		panic("bug")
	}
}

func (l *lexer) lexEqToken() token {
	r, _ := l.nextRune()
	switch {
	case r == '=':
		return token{kind: tokenEq, val: "="}
	default:
		panic("bug")
	}
}

func (l *lexer) lexStarToken() token {
	r, _ := l.nextRune()
	switch {
	case r == '*':
		return token{kind: tokenStar, val: "*"}
	default:
		panic("bug")
	}
}

func isAlphaNumeric(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}

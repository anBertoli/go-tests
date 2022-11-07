package main

import (
	"errors"
	"fmt"
)

type parser struct {
	lx *lexer
}

func (p *parser) parseStmt() (selectStmt, error) {
	next := p.lx.PeekTokenAfterSpace()

	switch next.kind {
	case tokenSelect:
		return p.parseSelectStmt(), nil
	case tokenEnd:
		return selectStmt{}, errEnd
	default:
		panic(fmt.Sprintf("%+v", next))
	}
}

// Expressions

type expr struct {
	kind  exprKind
	raw   string
	left  *expr
	op    exprOp
	right *expr
}

type exprKind uint

const (
	exprBinComp exprKind = iota
	exprBinBool
	exprIdent
	exprInt
	exprStr
)

func (ek exprKind) String() string {
	switch ek {
	case exprBinComp:
		return "comp"
	case exprBinBool:
		return "bool"
	case exprIdent:
		return "ident"
	case exprInt:
		return "int"
	case exprStr:
		return "str"
	default:
		panic("boh")
	}
}

// Expressions operators

type exprOp uint

const (
	opEq exprOp = iota
	opAnd
	opOr
)

func (ep exprOp) String() string {
	s, found := opToString[ep]
	if !found {
		s = "<invalid>"
	}
	return s
}

var opToString = map[exprOp]string{
	opEq:  "=",
	opAnd: "AND",
	opOr:  "OR",
}

func (p *parser) parseExpr() expr {
	comp := p.parseComparison()

	for {
		var op exprOp
		next := p.lx.PeekTokenAfterSpace()
		switch next.kind {
		case tokenAnd:
			op = opAnd
		case tokenOr:
			op = opOr
		default:
			// ended expression
			return comp
		}
		_ = p.lx.NextTokenAfterSpace()

		nextComp := p.parseComparison()
		clone := comp
		comp = expr{
			kind:  exprBinBool,
			left:  &clone,
			op:    op,
			right: &nextComp,
		}
	}
}

// a == b
func (p *parser) parseComparison() expr {
	left, wToken := p.parseExprVal()
	if wToken != nil {
		panic(wToken)
	}

	// TODO
	var op exprOp
	next := p.lx.PeekTokenAfterSpace()
	switch next.kind {
	case tokenEq:
		op = opEq
	default:
		// ended expression
		return left
	}
	_ = p.lx.NextTokenAfterSpace()

	right, wToken := p.parseExprVal()
	if wToken != nil {
		panic(wToken)
	}

	return expr{
		kind:  exprBinComp,
		left:  &left,
		op:    op,
		right: &right,
	}
}

func (p *parser) parseExprVal() (expr, *token) {
	next := p.lx.NextTokenAfterSpace()
	switch next.kind {
	case tokenIdent:
		return expr{raw: next.val, kind: exprIdent}, nil
	case tokenStr:
		return expr{raw: next.val, kind: exprStr}, nil
	case tokenInt:
		return expr{raw: next.val, kind: exprInt}, nil
	default:
		return expr{}, &next
	}
}

func assert[T comparable](got, want T) {
	if got != want {
		panic(fmt.Sprintf(
			"assert { got: %v, want: %v }",
			got, want,
		))
	}
}

var (
	errEnd = errors.New("end")
)

package main

import "strconv"

type selectStmt struct {
	table   string
	columns []string
	where   *expr
	limit   uint
}

// SELECT <col1, col2, col3> FROM <table> WHERE <expr> LIMIT <int>

func (p *parser) parseSelectStmt() selectStmt {
	next := p.lx.NextTokenAfterSpace()
	assert(next.kind, tokenSelect)

	// columns
	var columns []string
	var delPrev bool
	for {
		next = p.lx.NextTokenAfterSpace()
		switch {
		case next.kind == tokenIdent:
			columns = append(columns, next.val)
			delPrev = false
			continue
		case next.kind == tokenDel && !delPrev:
			delPrev = true
			continue
		case next.kind == tokenDel && delPrev:
			panic("multiple delimiters")
		}
		if next.kind != tokenFrom {
			panic(next.String())
		}
		break
	}

	// table
	var table string
	next = p.lx.NextTokenAfterSpace()
	if next.kind != tokenIdent {
		panic("table name expected")
	}
	table = next.val

	// where
	var where expr
	var limit uint64
	next = p.lx.NextTokenAfterSpace()
	if next.kind == tokenWhere {
		where = p.parseExpr()
		next = p.lx.NextTokenAfterSpace()
	}
	if next.kind == tokenLimit {
		limitEx, _ := p.parseExprVal()
		if limitEx.kind == exprInt {
			limit, _ = strconv.ParseUint(limitEx.raw, 10, 64)
		} else {
			panic("LIMIT <int> expected")
		}
		next = p.lx.NextTokenAfterSpace()
	}
	if next.kind != tokenEnd {
		panic("END expected")
	}

	return selectStmt{
		table:   table,
		columns: columns,
		where:   &where,
		limit:   uint(limit),
	}
}

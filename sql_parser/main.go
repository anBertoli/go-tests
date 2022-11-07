package main

import (
	"fmt"
	"strings"
)

func main() {
	lx := lexer{
		sql: "SELECT col1," +
			"col2 FROM users WHERE col2=786 AND col4=1243 OR col5=\"ciao\" LIMIT 23",
		pos: 0,
	}
	parser := parser{
		lx: &lx,
	}

	for {
		stmt, err := parser.parseStmt()
		if err != nil {
			fmt.Println(err)
			break
		}
		printStmt(0, stmt)
	}
}

func printStmt(depth int, stmt selectStmt) {
	fmt.Printf("%s{SELECT}\n", indent(depth))
	fmt.Printf("%s{cols: '%+v'}\n", indent(depth+1), stmt.columns)
	fmt.Printf("%s{FROM}\n", indent(depth))
	fmt.Printf("%s{table: '%s'}\n", indent(depth+1), stmt.table)
	fmt.Printf("%s{WHERE}\n", indent(depth))
	if stmt.where != nil {
		printExpr(depth+1, *stmt.where)
	}
	fmt.Printf("%s{LIMIT}\n", indent(depth))
	fmt.Printf("%s{limit: %d}\n", indent(depth), stmt.limit)
}

func printExpr(depth int, exp expr) {
	switch exp.kind {
	case exprIdent:
		fmt.Printf("%s{expr, kind: '%s', ident: '%s'}\n", indent(depth), exp.kind,
			exp.raw)
	case exprInt:
		fmt.Printf("%s{expr, kind: '%s', num: '%s'}\n", indent(depth), exp.kind, exp.raw)
	case exprStr:
		fmt.Printf("%s{expr, kind: '%s', str: '%s'}\n", indent(depth), exp.kind, exp.raw)
	case exprBinComp, exprBinBool:
		fmt.Printf("%s{op: '%s' (%d)}\n", indent(depth), exp.op.String(), exp.op)
		fmt.Printf("%s{left}\n", indent(depth))
		printExpr(depth+1, *exp.left)
		fmt.Printf("%s{right}\n", indent(depth))
		printExpr(depth+1, *exp.right)
	}
}

func indent(depth int) string {
	return strings.Repeat(" ", depth*3)
}

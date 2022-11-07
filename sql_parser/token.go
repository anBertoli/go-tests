package main

import "fmt"

type token struct {
	kind tokenKind
	val  string
}

func (t *token) String() string {
	s, found := tokenToStr[t.kind]
	if !found {
		s = "<invalid>"
	}
	return fmt.Sprintf(
		"token {kind: '%s' (%d), val: '%s'}",
		s, t.kind, t.val,
	)
}

type tokenKind uint

const (
	tokenSpace tokenKind = iota // space/tabs/newlines
	tokenEnd                    // end of input

	tokenIdent // identifier
	tokenStr   // quoted string
	tokenInt   // integer number
	tokenFloat // float number

	tokenEq   // =
	tokenGt   // >
	tokenGtEq // >=
	tokenLs   // <
	tokenLsEq // <=

	tokenDel      // ,
	tokenOpenPar  // (
	tokenClosePar // )
	tokenStar     // *

	tokenAnd    // AND keyword
	tokenOr     // OR keyword
	tokenInsert // INSERT keyword
	tokenInto   // INTO keyword
	tokenSelect // SELECT keyword
	tokenDelete // DELETE keyword
	tokenFrom   // FROM keyword
	tokenWhere  // WHERE keyword
	tokenLimit  // LIMIT keyword
	tokenValues // VALUES keyword
	tokenUpdate // UPDATE keyword
	tokenSet    // SET keyword
	tokenCreate // CREATE keyword
	tokenDrop   // DROP keyword
)

var tokenToStr = map[tokenKind]string{
	tokenSpace: "tokenSpace",
	tokenEnd:   "tokenEnd",

	tokenIdent: "tokenIdent",
	tokenStr:   "tokenStr",
	tokenInt:   "tokenInt",
	tokenFloat: "tokenFloat",

	tokenEq:   "tokenEq",
	tokenGt:   "tokenGt",
	tokenGtEq: "tokenGtEq",
	tokenLs:   "tokenLs",
	tokenLsEq: "tokenLsEq",

	tokenDel:      "tokenDel",
	tokenOpenPar:  "tokenOpenPar",
	tokenClosePar: "tokenClosePar",
	tokenStar:     "tokenStar",

	tokenAnd:    "tokenAnd",
	tokenOr:     "tokenOr",
	tokenInsert: "tokenInsert",
	tokenInto:   "tokenInto",
	tokenSelect: "tokenSelect",
	tokenDelete: "tokenDelete",
	tokenFrom:   "tokenFrom",
	tokenWhere:  "tokenWhere",
	tokenLimit:  "tokenLimit",
	tokenValues: "tokenValues",
	tokenUpdate: "tokenUpdate",
	tokenSet:    "tokenSet",
	tokenCreate: "tokenCreate",
	tokenDrop:   "tokenDrop",
}

var keywordsToToken = map[string]tokenKind{
	"AND":    tokenAnd,
	"OR":     tokenOr,
	"INSERT": tokenInsert,
	"INTO":   tokenInto,
	"SELECT": tokenSelect,
	"DELETE": tokenDelete,
	"FROM":   tokenFrom,
	"WHERE":  tokenWhere,
	"LIMIT":  tokenLimit,
	"VALUES": tokenValues,
	"UPDATE": tokenUpdate,
	"SET":    tokenSet,
	"CREATE": tokenCreate,
	"DROP":   tokenDrop,
}

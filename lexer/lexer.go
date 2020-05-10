package lexer

import (
	"strings"
	"unicode"
)

type Token struct {
	Type map[int] bool
	Lexem  string
}

var T map[string]int = map[string]int {
	"IDENTIFER": 1,
	"LB": 2,
	"RB": 3,
	"RELATION": 4,
	"SIGN": 5,
	"ADDITION": 6,
	"MULTI": 7,
	"NOT": 8,
	"CONSTANT": 9,
	"LP": 10,
	"RP": 11,
	"SEMICOLON": 12,
	"ASSIGNMENT": 13,
}

var multi map[string]bool = map[string]bool {
	"*": true,
	"/": true,
	"div": true,
	"mod": true,
	"and": true,
}

func GetTokenType(token string) map[int] bool {
	tokenType := make(map[int] bool)
	if token == "("{
		tokenType[T["LP"]] = true
	} else if token == ")" {
		tokenType[T["RP"]] = true
	} else if token == "{" {
		tokenType[T["LB"]] = true
	} else if token == "}" {
		tokenType[T["RB"]] = true
	} else if token == ";" {
		tokenType[T["SEMICOLON"]] = true
	}
	if token == "=" {
		tokenType[T["ASSIGNMENT"]] = true
	}
	if token == "not" {
		tokenType[T["NOT"]] = true
	}
	if strings.Index("=<><<=>>=", token) != -1 {
		tokenType[T["RELATION"]] = true
	}
	if strings.Index("+-or", token) != -1 {
		tokenType[T["ADDITION"]] = true
	}
	if multi[token] {
		tokenType[T["MULTI"]] = true
	}
	if strings.Index("+-", token) != -1 {
		tokenType[T["SIGN"]] = true
	}

	if len(tokenType) == 0 {
		if unicode.IsLetter(rune(token[0])) {
			tokenType[T["IDENTIFER"]] = true
		} else {
			tokenType[T["CONSTANT"]] = true
		}
	}
	return tokenType
}

func Analyze(input string) []Token {
	strs := strings.Split(input, "\n")
	var lexems []string
	for _, str := range strs {
		splt := strings.Split(str, " ")
		if len(splt) > 1 {
			lexems = append(lexems, splt...)
		} else {
			lexems = append(lexems, str)
		}
		 
	}

	tokens := make([]Token, len(lexems))
	for i, lexem := range lexems {
		tokens[i] = Token{
			Type: GetTokenType(lexem),
			Lexem: lexem,
		}
	}
	return tokens
}
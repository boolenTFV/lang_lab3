package parser

import (
	// "strings"
	"lab3/lexer"
	"fmt"
)

/*
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
*/
type NonTerm struct {
	Name string
	Value interface{}
}

func Factor(tokens []lexer.Token, state *int) *NonTerm {
	Name := "Фактор"
	*state++;
	if tokens[*state].Type[lexer.T["IDENTIFER"]] {
		return &NonTerm{Name, tokens[*state].Lexem}
	}

	if tokens[*state].Type[lexer.T["CONSTANT"]] {
		return &NonTerm{Name, tokens[*state].Lexem}
	}

	lp := tokens[*state]
	if lp.Type[lexer.T["LP"]] {
		simple := SimpleExpression(tokens, state)
		if simple != nil {
			*state++
			rp := tokens[*state]
			if rp.Type[lexer.T["RP"]] {
				return &NonTerm{Name, []interface{}{lp.Lexem, simple, rp.Lexem}}
			}
		}
		return nil
	}

	not := tokens[*state]
	if not.Type[lexer.T["NOT"]] {
		factor := Factor(tokens, state);
		if factor != nil {
			return &NonTerm{Name, []interface{}{not.Lexem, factor}}
		}
	}

	return nil
}

func Term(tokens []lexer.Token, state *int) *NonTerm {
	Name := "Терм"
	factor := Factor(tokens, state);
	if factor != nil {
		stateCpy := *state
		tail := TermTail(tokens, state);
		if tail != nil {
			return &NonTerm{Name, []interface{}{factor, tail}}
		}
		*state = stateCpy
		return &NonTerm{Name, factor}
	}
	return nil
}

func TermTail(tokens []lexer.Token, state *int) *NonTerm {
	Name := "Хвост терма"
	*state++
	multi := tokens[*state]
	if multi.Type[lexer.T["MULTI"]] {
		factor := Factor(tokens, state);
		if factor != nil {
			stateCpy := *state
			tail := TermTail(tokens, state);
			if tail != nil {
				return &NonTerm{Name, []interface{}{multi.Lexem, factor, tail}}
			}
			*state = stateCpy
			return &NonTerm{Name, []interface{}{multi.Lexem, factor}}
		}
	}
	return nil
}

func SimpleExpression(tokens []lexer.Token, state *int) *NonTerm {
	Name := "Простое выражение"
	*state++;
	sign := tokens[*state]
	if !sign.Type[lexer.T["SIGN"]] {
		*state--;
	}

	term := Term(tokens, state)
	if term != nil {
		stateCpy := *state
		tail := SimpleExpressionTail(tokens, state)
		if tail != nil {
			if sign.Type[lexer.T["SIGN"]] {
				return &NonTerm{Name, []interface{}{sign.Lexem, term, tail}}
			}
			return &NonTerm{Name, []interface{}{term, tail}}
		}
		*state = stateCpy

		if sign.Type[lexer.T["SIGN"]] {
			return &NonTerm{Name, []interface{}{sign.Lexem, term}}
		}
		return &NonTerm{Name, term}
	}
	return nil
}

func SimpleExpressionTail(tokens []lexer.Token, state *int) *NonTerm {
	Name := "Хвост простого выражения"
	*state++;
	add := tokens[*state]
	if add.Type[lexer.T["ADDITION"]] {
		term := Term(tokens, state)
		if term != nil {
			stateCpy := *state
			tail := SimpleExpressionTail(tokens, state)
			if tail != nil {
				return &NonTerm{Name, []interface{}{add.Lexem, term, tail}}
			}
			*state = stateCpy

			return &NonTerm{Name, []interface{}{add.Lexem, term}}
		}
	}

	return nil
}

func Expression(tokens []lexer.Token, state *int) *NonTerm {
	Name := "Выражение"
	simple := SimpleExpression(tokens, state)
	if simple != nil {
		stateCpy := *state
		*state++
		relation := tokens[*state]
		if relation.Type[lexer.T["RELATION"]] {
			simple2 := SimpleExpression(tokens, state)
			if simple2 != nil {
				return &NonTerm{Name, []interface{}{simple, relation.Lexem, simple2}}
			}
		}
		*state = stateCpy
		return &NonTerm{Name, simple}
	}

	return nil
}



func Assignment(tokens []lexer.Token, state *int) *NonTerm {
	Name := "Оператор"
	*state++
	id := tokens[*state]
	if id.Type[lexer.T["IDENTIFER"]] {
		*state++
		assignment := tokens[*state]
		if assignment.Type[lexer.T["ASSIGNMENT"]] {
			expression := Expression(tokens, state)
			if expression != nil {
				return &NonTerm{Name, []interface{}{id.Lexem, assignment.Lexem, expression}}
			}
		}
		return nil
	}
	
	lb := tokens[*state]
	if lb.Type[lexer.T["LB"]] {
		list := AssignmentList(tokens, state)
		if list != nil {
			*state++
			rb := tokens[*state]
			if rb.Type[lexer.T["RB"]] {
				return &NonTerm{Name, []interface{}{lb.Lexem, list, rb.Lexem}}
			}
		}
	}

	return nil
}

func Tail(tokens []lexer.Token, state *int) *NonTerm {
	Name := "Хвост"
	*state++;
	semicolon := tokens[*state]
	if semicolon.Type[lexer.T["SEMICOLON"]] {
		assignment := Assignment(tokens, state)
		if assignment != nil {
			stateCpy := *state
			tail := Tail(tokens, state)
			if tail != nil {
				return &NonTerm{Name, []interface{}{semicolon.Lexem, assignment, tail}}
			}
			*state = stateCpy
			return &NonTerm{Name, []interface{}{semicolon.Lexem, assignment}}
		}
	}
	return nil
}

func AssignmentList(tokens []lexer.Token, state *int) *NonTerm {
		Name := "Список операторов"
		assignment := Assignment(tokens, state)
		if assignment != nil {
			stateCpy := *state
			tail := Tail(tokens, state)
			if tail != nil {
				return &NonTerm{Name, []interface{}{assignment, tail}}
			}
			*state = stateCpy
			return &NonTerm{Name, assignment}
		}

	return nil
}

func Program(tokens []lexer.Token) *NonTerm {
	Name := "Программа"
	state := 0
	lb := tokens[state]
	if lb.Type[lexer.T["LB"]] {
		list := AssignmentList(tokens, &state)
		if list != nil {
			state++
			rb := tokens[state]
			if rb.Type[lexer.T["RB"]] {
				return &NonTerm{Name, []interface{}{lb.Lexem, list, rb.Lexem}}
			}
		}
	}
	if state  >= len(tokens) {
		state = len(tokens) - 1
	}
	fmt.Printf("Ошибка в позиции %v, возле токена '%v' \n", state, tokens[state].Lexem)
	return nil
}
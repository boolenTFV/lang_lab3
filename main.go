package main

import (
	"fmt"
	"lab3/lexer"
	"lab3/parser"
	"encoding/json"
    "io/ioutil"
)

	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func writeFile(val []byte) {
    err := ioutil.WriteFile("res.json", val, 0644)
    check(err)
}

func readFile(file string) string{
	dat, err := ioutil.ReadFile(file)
	check(err)
	return string(dat)
}

func main() {
	fmt.Println("start prog")
	program:= readFile("prog.gr")
	tokens := lexer.Analyze(program)
	programTree := parser.Program(tokens)

	fmt.Println("start prog1")
	program = readFile("prog1.gr")
	tokens = lexer.Analyze(program)
	programTree = parser.Program(tokens)
	
	fmt.Println("start prog2")
	program= readFile("prog2.gr")
	tokens = lexer.Analyze(program)
	programTree = parser.Program(tokens)
	res, _ := json.MarshalIndent(programTree, "", "  ")
	writeFile(res)

	fmt.Println("start err")
	program = readFile("err.gr")
	tokens = lexer.Analyze(program)
	programTree = parser.Program(tokens)

	fmt.Println("start err1")
	program = readFile("err1.gr")
	tokens = lexer.Analyze(program)
	programTree = parser.Program(tokens)

	fmt.Println("start err2")
	program = readFile("err2.gr")
	tokens = lexer.Analyze(program)
	programTree = parser.Program(tokens)
}


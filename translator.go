package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type CommandType int

const (
	CArithmetic CommandType = iota
	CPush
	CPop
	CLabel
	CGoto
	CIf
	CFunction
	CReturn
	CCall
)

// Handles the parsing of a single .vm file,
// and encapsulates access to the input code.
//
// It reads VM commands, parses them, and provides
// convenient access to their compoenents.
// In addition, it removes all white space and comments.
type Parser struct {
	InputFile      string
	CurrentCommand string
	Scanner        *bufio.Scanner
}

// Opens the input fiel and gets ready to parse it.
func NewParser(filename string) *Parser {
	p := new(Parser)
	p.InputFile = filename
	file, err := os.Open(p.InputFile)
	if err != nil {
		log.Fatal(err)
	}
	// defer file.Close()
	p.Scanner = bufio.NewScanner(file)
	return p
}

// Are there more commmands in the input
func (p *Parser) HasMoreCommands() bool {
	return p.Scanner.Scan()
}

// Reads the next cmd from the input and makes it the current cmd.
//
// Should be called only if `HasMoreCommands()` is true.
// Initially, there is no current cmdS
func (p *Parser) Advance() {
	for {
		token := strings.TrimSpace(p.Scanner.Text())
		if len(token) > 0 {
			p.CurrentCommand = token
			return
		}
		p.Scanner.Scan() // Advance to next token
	}
}

// Returns the type of the current cmd.
func (p *Parser) CommandType() CommandType {
	tokens := strings.Fields(p.CurrentCommand)
	switch tokens[0] {
	case "add", "sub", "neg", "eq", "gt", "lt", "and", "or", "not":
		fmt.Println("COMMAND TYPE C_ARITHMETIC: ", p.CurrentCommand)
		return CArithmetic
	case "push":
		fmt.Println("COMMAND TYPE C_PUSH: ", p.CurrentCommand)
		return CPush
	case "pop":
		fmt.Println("COMMAND TYPE C_POP: ", p.CurrentCommand)
		return CPop
	case "label":
		fmt.Println("COMMAND TYPE C_LABEL: ", p.CurrentCommand)
		return CLabel
	case "goto":
		fmt.Println("COMMAND TYPE C_GOTO: ", p.CurrentCommand)
		return CGoto
	case "if":
		fmt.Println("COMMAND TYPE C_IF: ", p.CurrentCommand)
		return CIf
	case "function":
		fmt.Println("COMMAND TYPE C_Function: ", p.CurrentCommand)
		return CFunction
	case "call":
		fmt.Println("COMMAND TYPE C_POP: ", p.CurrentCommand)
		return CCall
	case "return":
		fmt.Println("COMMAND TYPE C_POP: ", p.CurrentCommand)
		return CReturn
	}
	return -1
}

// Returns the first argument of the current cmd.
func (p *Parser) Arg1() string {
	return ""
}

// Returns the second argument of the current cmd.
func (p *Parser) Arg2() int {
	return 0
}

func ReadFile() {
	parser := new(Parser)
	file, err := os.Open(parser.InputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	filename := "BasicTest.vm"
	parser := NewParser(filename)
	for parser.HasMoreCommands() {
		parser.Advance()
		parser.CommandType()
		parser.Arg1()
		parser.Arg2()
		fmt.Println(parser.CurrentCommand)
		fmt.Println("")
	}
}

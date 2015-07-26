package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	p.CurrentCommand = p.Scanner.Text()

}

// Returns the type of the current cmd.
func (p *Parser) CommandType() CommandType {
	return CArithmetic
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
		fmt.Println(parser.CurrentCommand)
	}
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
	InputFile          string
	CurrentCommand     string
	CurrentCommandType CommandType
	Scanner            *bufio.Scanner
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
		if isComment(p.Scanner.Text()) == false {
			token := strings.TrimSpace(p.Scanner.Text())
			if len(token) > 0 {
				p.CurrentCommand = token
				return
			}
		}
		p.Scanner.Scan() // Advance to next token
	}
}

func isComment(str string) bool {
	return strings.Index(str, "//") == 0
}

// Returns the type of the current cmd.
func (p *Parser) CommandType() CommandType {
	tokens := strings.Fields(p.CurrentCommand)
	switch tokens[0] {
	case "add", "sub", "neg", "eq", "gt", "lt", "and", "or", "not":
		fmt.Println("COMMAND TYPE : C_ARITHMETIC")
		return CArithmetic
	case "push":
		fmt.Println("COMMAND TYPE : C_PUSH")
		return CPush
	case "pop":
		fmt.Println("COMMAND TYPE : C_POP")
		return CPop
	case "label":
		fmt.Println("COMMAND TYPE : C_LABEL")
		return CLabel
	case "goto":
		fmt.Println("COMMAND TYPE : C_GOTO")
		return CGoto
	case "if":
		fmt.Println("COMMAND TYPE : C_IF")
		return CIf
	case "function":
		fmt.Println("COMMAND TYPE : C_Function")
		return CFunction
	case "call":
		fmt.Println("COMMAND TYPE : C_POP")
		return CCall
	case "return":
		fmt.Println("COMMAND TYPE : C_POP")
		return CReturn
	}
	return -1
}

// Returns the first argument of the current cmd.
func (p *Parser) Arg1() string {
	tokens := strings.Fields(p.CurrentCommand)
	switch p.CurrentCommandType {
	case CArithmetic:
		return tokens[0]
	default:
		return tokens[1]
	}
}

// Returns the second argument of the current cmd.
func (p *Parser) Arg2() int {
	tokens := strings.Fields(p.CurrentCommand)
	d, _ := strconv.ParseInt(tokens[2], 0, 0)
	return int(d)
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

type CodeWriter struct {
	outputFilename string
}

func NewCodeWriter(outputFilename string) *CodeWriter {
	c := new(CodeWriter)
	c.outputFilename = outputFilename
	return c
}

func (c *CodeWriter) SetFileName(filename string) {

}

func (c *CodeWriter) WriteArithmetic(command string) {

}

func (c *CodeWriter) WritePushPop(command CommandType, segment string, index int) {

}

func (c *CodeWriter) Close() {

}

func main() {
	inputFilename := "BasicTest.vm"
	outputFilename := "Prog.asm"

	parser := NewParser(inputFilename)
	codeWriter := NewCodeWriter(outputFilename)

	for parser.HasMoreCommands() {
		parser.Advance()
		parser.CurrentCommandType = parser.CommandType()

		if parser.CurrentCommandType != CReturn {
			args1 := parser.Arg1()
			fmt.Println("ARGS1 : ", args1)
		}
		switch parser.CurrentCommandType {
		case CPush, CPop, CFunction, CCall:
			args2 := parser.Arg2()
			fmt.Println("ARGS2 : ", args2)
		}

		fmt.Println(parser.CurrentCommand)
		fmt.Println("")
	}
}

package repl

import (
	// "bufio"
	"bufio"
	"fmt"
	"io"
	"monkey/compiler"
	"monkey/lexer"
	"monkey/object"

	// "monkey/object"
	"monkey/parser"
	"monkey/vm"
	"os"
	"strings"
	// "monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	constants := []object.Object{}
	globals := make([]object.Object,vm.GlobalsSize)
	symbolTable := compiler.NewSymbolTable()
	// env := object.NewEnvironment()
	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := strings.TrimSpace(scanner.Text())

		if line == "exit" {
			io.WriteString(out, "Good Bye!\n")
			os.Exit(0)
		}
		if line == ""{
			continue
		}

		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}
		comp := compiler.NewWithState(symbolTable,constants)
		err := comp.Compile(program)
		
		
		if err != nil {
			fmt.Fprintf(out,"Woops! Compilation failed:\n %s\n",err)
			continue
		}
		code := comp.Bytecode()
		machine:=vm.NewWithGlobalsStore(code,globals)
		err = machine.Run()
		if err != nil {
			fmt.Fprintf(out,"Woops! Executing bytecode failed:\n %s\n",err)
			continue
		}
		lastPopped := machine.LastPoppedStackElem()
		io.WriteString(out,lastPopped.Inspect())
		io.WriteString(out,"\n")
	}
}
func RunScript(script string,out io.Writer){

		l := lexer.New(script)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
		}
		comp := compiler.New()
		err := comp.Compile(program)
		
		
		
		if err != nil {
			fmt.Fprintf(out,"Woops! Compilation failed:\n %s\n",err)
		}
		code := comp.Bytecode()
		
		// fmt.Println("Constant count: ",len(code.Constants))
		// for _,o := range code.Constants{
		// 	fmt.Println(o.Inspect())
		// }
		globals := make([]object.Object,vm.GlobalsSize)
		machine:=vm.NewWithGlobalsStore(code,globals)
		err = machine.Run()
		if err != nil {
			fmt.Fprintf(out,"Woops! Executing bytecode failed:\n %s\n",err)
		}
		lastPopped := machine.LastPoppedStackElem()
		io.WriteString(out,lastPopped.Inspect())
		io.WriteString(out,"\n")
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

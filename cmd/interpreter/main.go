package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/WillKirkmanM/bytecode/pkg/bytecode"
	"github.com/WillKirkmanM/bytecode/pkg/examples"
)

func main() {
    if len(os.Args) > 1 {
        filename := os.Args[1]
        if err := runFile(filename); err != nil {
            fmt.Fprintf(os.Stderr, "Error: %v\n", err)
            os.Exit(1)
        }
    } else {
        runInteractive()
    }
}

func runFile(filename string) error {
    content, err := os.ReadFile(filename)
    if err != nil {
        return fmt.Errorf("failed to read file: %w", err)
    }

    assembler := bytecode.NewAssembler()
    instructions, err := assembler.AssembleFromString(string(content))
    if err != nil {
        return fmt.Errorf("assembly error: %w", err)
    }

    vm := bytecode.NewVM()
    vm.LoadProgram(instructions)
    
    if err := vm.Execute(); err != nil {
        return fmt.Errorf("execution error: %w", err)
    }

    return nil
}

func runInteractive() {
    fmt.Println("Bytecode Interpreter v1.0")
    fmt.Println("Type 'help' for commands, 'exit' to quit")
    
    vm := bytecode.NewVM()
    scanner := bufio.NewScanner(os.Stdin)
    
    for {
        fmt.Print("> ")
        if !scanner.Scan() {
            break
        }
        
        input := strings.TrimSpace(scanner.Text())
        if input == "" {
            continue
        }
        
        switch input {
        case "exit", "quit":
            fmt.Println("Goodbye!")
            return
            
        case "help":
            printHelp()
            
        case "examples":
            runExamples(vm)
            
        case "debug on":
            vm.SetDebug(true)
            fmt.Println("Debug mode enabled")
            
        case "debug off":
            vm.SetDebug(false)
            fmt.Println("Debug mode disabled")
            
        default:
            if strings.HasPrefix(input, "run ") {
                code := strings.TrimPrefix(input, "run ")
                runCode(vm, code)
            } else {
                fmt.Println("Unknown command. Type 'help' for available commands.")
            }
        }
    }
}

func printHelp() {
    fmt.Println("Available commands:")
    fmt.Println("  help          - Show this help message")
    fmt.Println("  examples      - Run example programs")
    fmt.Println("  debug on/off  - Enable/disable debug mode")
    fmt.Println("  run <code>    - Run assembly code")
    fmt.Println("  exit          - Exit the interpreter")
    fmt.Println("\nExample assembly code:")
    fmt.Println("  run LOAD_CONST 42; PRINT; HALT")
}

func runExamples(vm *bytecode.VM) {
    fmt.Println("\n=== Running Example Programs ===")
    
    fmt.Println("\n1. Simple Calculator (10 + 20):")
    vm.LoadProgram(examples.SimpleCalculator())
    if err := vm.Execute(); err != nil {
        fmt.Printf("Error: %v\n", err)
    }
    
    fmt.Println("\n2. Countdown from 5:")
    vm.LoadProgram(examples.CountDown(5))
    if err := vm.Execute(); err != nil {
        fmt.Printf("Error: %v\n", err)
    }
    
    fmt.Println("\n3. Fibonacci sequence (first 10 numbers):")
    vm.LoadProgram(examples.FibonacciSequence(10))
    if err := vm.Execute(); err != nil {
        fmt.Printf("Error: %v\n", err)
    }
}

func runCode(vm *bytecode.VM, code string) {
    code = strings.ReplaceAll(code, ";", "\n")
    
    assembler := bytecode.NewAssembler()
    instructions, err := assembler.AssembleFromString(code)
    if err != nil {
        fmt.Printf("Assembly error: %v\n", err)
        return
    }
    
    vm.LoadProgram(instructions)
    if err := vm.Execute(); err != nil {
        fmt.Printf("Execution error: %v\n", err)
    }
}
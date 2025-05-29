package bytecode

import "fmt"

type Opcode uint8

const (
    // Core instructions
    HALT       Opcode = 0x00 // Stop execution
    LOAD_CONST Opcode = 0x01 // Load constant to stack
    LOAD_VAR   Opcode = 0x02 // Load variable to stack
    STORE_VAR  Opcode = 0x03 // Store top of stack to variable
    
    // Arithmetic operations
    ADD Opcode = 0x04 // Add two values from stack
    SUB Opcode = 0x05 // Subtract two values from stack
    MUL Opcode = 0x06 // Multiply two values from stack
    DIV Opcode = 0x07 // Divide two values from stack
    
    // Control flow
    JMP Opcode = 0x08 // Jump to address
    JZ  Opcode = 0x09 // Jump if zero
    JNZ Opcode = 0x0A // Jump if not zero
    
    // I/O
    PRINT Opcode = 0x0B // Print top of stack
)

type Instruction struct {
    Opcode   Opcode
    Operand1 int32
    Operand2 int32
}

func (i Instruction) String() string {
    switch i.Opcode {
    case HALT:
        return "HALT"
    case LOAD_CONST:
        return fmt.Sprintf("LOAD_CONST %d", i.Operand1)
    case LOAD_VAR:
        return fmt.Sprintf("LOAD_VAR %d", i.Operand1)
    case STORE_VAR:
        return fmt.Sprintf("STORE_VAR %d", i.Operand1)
    case ADD:
        return "ADD"
    case SUB:
        return "SUB"
    case MUL:
        return "MUL"
    case DIV:
        return "DIV"
    case JMP:
        return fmt.Sprintf("JMP %d", i.Operand1)
    case JZ:
        return fmt.Sprintf("JZ %d", i.Operand1)
    case JNZ:
        return fmt.Sprintf("JNZ %d", i.Operand1)
    case PRINT:
        return "PRINT"
    default:
        return fmt.Sprintf("UNKNOWN(%d)", i.Opcode)
    }
}
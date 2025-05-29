package bytecode

import (
	"fmt"
	"strconv"
	"strings"
)

type Assembler struct {
    labels map[string]int32
}

func NewAssembler() *Assembler {
    return &Assembler{
        labels: make(map[string]int32),
    }
}

func (a *Assembler) AssembleFromString(source string) ([]Instruction, error) {
    lines := strings.Split(source, "\n")
    return a.assembleLines(lines)
}

func (a *Assembler) cleanLine(line string) string {
    if idx := strings.Index(line, ";"); idx != -1 {
        line = line[:idx]
    }
    return strings.TrimSpace(line)
}

func (a *Assembler) assembleLines(lines []string) ([]Instruction, error) {
    var instructions []Instruction
    
    pc := int32(0)
    for _, line := range lines {
        line = a.cleanLine(line)
        if line == "" {
            continue
        }
        
        if strings.HasSuffix(line, ":") {
            label := strings.TrimSuffix(line, ":")
            a.labels[label] = pc
        } else {
            pc++
        }
    }
    
    for _, line := range lines {
        line = a.cleanLine(line)
        if line == "" || strings.HasSuffix(line, ":") {
            continue
        }
        
        instruction, err := a.parseLine(line)
        if err != nil {
            return nil, fmt.Errorf("error parsing line '%s': %w", line, err)
        }
        
        instructions = append(instructions, instruction)
    }
    
    return instructions, nil
}

func (a *Assembler) parseLine(line string) (Instruction, error) {
    parts := strings.Fields(line)
    if len(parts) == 0 {
        return Instruction{}, fmt.Errorf("empty instruction")
    }
    
    opcode := strings.ToUpper(parts[0])
    
    switch opcode {
    case "HALT":
        return Instruction{Opcode: HALT}, nil
        
    case "LOAD_CONST":
        if len(parts) != 2 {
            return Instruction{}, fmt.Errorf("LOAD_CONST requires 1 operand")
        }
        value, err := a.parseOperand(parts[1])
        if err != nil {
            return Instruction{}, err
        }
        return Instruction{Opcode: LOAD_CONST, Operand1: value}, nil
        
    case "LOAD_VAR":
        if len(parts) != 2 {
            return Instruction{}, fmt.Errorf("LOAD_VAR requires 1 operand")
        }
        value, err := a.parseOperand(parts[1])
        if err != nil {
            return Instruction{}, err
        }
        return Instruction{Opcode: LOAD_VAR, Operand1: value}, nil
        
    case "STORE_VAR":
        if len(parts) != 2 {
            return Instruction{}, fmt.Errorf("STORE_VAR requires 1 operand")
        }
        value, err := a.parseOperand(parts[1])
        if err != nil {
            return Instruction{}, err
        }
        return Instruction{Opcode: STORE_VAR, Operand1: value}, nil
        
    case "ADD":
        return Instruction{Opcode: ADD}, nil
    case "SUB":
        return Instruction{Opcode: SUB}, nil
    case "MUL":
        return Instruction{Opcode: MUL}, nil
    case "DIV":
        return Instruction{Opcode: DIV}, nil
        
    case "JMP":
        if len(parts) != 2 {
            return Instruction{}, fmt.Errorf("JMP requires 1 operand")
        }
        value, err := a.parseOperand(parts[1])
        if err != nil {
            return Instruction{}, err
        }
        return Instruction{Opcode: JMP, Operand1: value}, nil
        
    case "JZ":
        if len(parts) != 2 {
            return Instruction{}, fmt.Errorf("JZ requires 1 operand")
        }
        value, err := a.parseOperand(parts[1])
        if err != nil {
            return Instruction{}, err
        }
        return Instruction{Opcode: JZ, Operand1: value}, nil
        
    case "JNZ":
        if len(parts) != 2 {
            return Instruction{}, fmt.Errorf("JNZ requires 1 operand")
        }
        value, err := a.parseOperand(parts[1])
        if err != nil {
            return Instruction{}, err
        }
        return Instruction{Opcode: JNZ, Operand1: value}, nil
        
    case "PRINT":
        return Instruction{Opcode: PRINT}, nil
        
    default:
        return Instruction{}, fmt.Errorf("unknown instruction: %s", opcode)
    }
}

func (a *Assembler) parseOperand(operand string) (int32, error) {
    if value, exists := a.labels[operand]; exists {
        return value, nil
    }
    
    value, err := strconv.ParseInt(operand, 10, 32)
    if err != nil {
        return 0, fmt.Errorf("invalid operand: %s", operand)
    }
    
    return int32(value), nil
}
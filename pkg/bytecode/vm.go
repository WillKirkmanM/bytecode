package bytecode

import (
	"errors"
	"fmt"
)

type VM struct {
    instructions []Instruction
    stack        []int32
    variables    map[int32]int32
    pc           int32
    running      bool
    debug        bool
}

func NewVM() *VM {
    return &VM{
        instructions: make([]Instruction, 0),
        stack:        make([]int32, 0),
        variables:    make(map[int32]int32),
        pc:           0,
        running:      false,
        debug:        false,
    }
}

func (vm *VM) LoadProgram(instructions []Instruction) {
    vm.instructions = instructions
    vm.pc = 0
}

func (vm *VM) SetDebug(debug bool) {
    vm.debug = debug
}

func (vm *VM) push(value int32) {
    vm.stack = append(vm.stack, value)
}

func (vm *VM) pop() (int32, error) {
    if len(vm.stack) == 0 {
        return 0, errors.New("stack underflow")
    }
    value := vm.stack[len(vm.stack)-1]
    vm.stack = vm.stack[:len(vm.stack)-1]
    return value, nil
}

func (vm *VM) peek() (int32, error) {
    if len(vm.stack) == 0 {
        return 0, errors.New("stack is empty")
    }
    return vm.stack[len(vm.stack)-1], nil
}

func (vm *VM) Execute() error {
    vm.running = true
    vm.pc = 0
    
    for vm.running && int(vm.pc) < len(vm.instructions) {
        if err := vm.executeInstruction(); err != nil {
            return fmt.Errorf("execution error at PC %d: %w", vm.pc, err)
        }
    }
    
    return nil
}

func (vm *VM) executeInstruction() error {
    if int(vm.pc) >= len(vm.instructions) {
        return errors.New("program counter out of bounds")
    }
    
    instruction := vm.instructions[vm.pc]
    
    if vm.debug {
        fmt.Printf("PC: %d, Instruction: %s, Stack: %v\n", vm.pc, instruction.String(), vm.stack)
    }
    
    switch instruction.Opcode {
    case HALT:
        vm.running = false
        
    case LOAD_CONST:
        vm.push(instruction.Operand1)
        vm.pc++
        
    case LOAD_VAR:
        value, exists := vm.variables[instruction.Operand1]
        if !exists {
            return fmt.Errorf("undefined variable %d", instruction.Operand1)
        }
        vm.push(value)
        vm.pc++
        
    case STORE_VAR:
        value, err := vm.pop()
        if err != nil {
            return err
        }
        vm.variables[instruction.Operand1] = value
        vm.pc++
        
    case ADD:
        b, err := vm.pop()
        if err != nil {
            return err
        }
        a, err := vm.pop()
        if err != nil {
            return err
        }
        vm.push(a + b)
        vm.pc++
        
    case SUB:
        b, err := vm.pop()
        if err != nil {
            return err
        }
        a, err := vm.pop()
        if err != nil {
            return err
        }
        vm.push(a - b)
        vm.pc++
        
    case MUL:
        b, err := vm.pop()
        if err != nil {
            return err
        }
        a, err := vm.pop()
        if err != nil {
            return err
        }
        vm.push(a * b)
        vm.pc++
        
    case DIV:
        b, err := vm.pop()
        if err != nil {
            return err
        }
        if b == 0 {
            return errors.New("division by zero")
        }
        a, err := vm.pop()
        if err != nil {
            return err
        }
        vm.push(a / b)
        vm.pc++
        
    case JMP:
        vm.pc = instruction.Operand1
        
    case JZ:
        value, err := vm.pop()
        if err != nil {
            return err
        }
        if value == 0 {
            vm.pc = instruction.Operand1
        } else {
            vm.pc++
        }
        
    case JNZ:
        value, err := vm.pop()
        if err != nil {
            return err
        }
        if value != 0 {
            vm.pc = instruction.Operand1
        } else {
            vm.pc++
        }
        
    case PRINT:
        value, err := vm.peek()
        if err != nil {
            return err
        }
        fmt.Println(value)
        vm.pc++
        
    default:
        return fmt.Errorf("unknown opcode: %d", instruction.Opcode)
    }
    
    return nil
}

func (vm *VM) GetStack() []int32 {
    return append([]int32(nil), vm.stack...)
}

func (vm *VM) GetVariables() map[int32]int32 {
    vars := make(map[int32]int32)
    for k, v := range vm.variables {
        vars[k] = v
    }
    return vars
}
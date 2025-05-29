package examples

import "github.com/WillKirkmanM/bytecode/pkg/bytecode"

func SimpleCalculator() []bytecode.Instruction {
    return []bytecode.Instruction{
        {Opcode: bytecode.LOAD_CONST, Operand1: 10},
        {Opcode: bytecode.LOAD_CONST, Operand1: 20},
        {Opcode: bytecode.ADD},
        {Opcode: bytecode.PRINT},
        {Opcode: bytecode.HALT},
    }
}

func FibonacciSequence(n int32) []bytecode.Instruction {
    return []bytecode.Instruction{
        // Initialize variables: fib(0) = 0, fib(1) = 1
        {Opcode: bytecode.LOAD_CONST, Operand1: 0},
        {Opcode: bytecode.STORE_VAR, Operand1: 0}, // var[0] = a = 0
        {Opcode: bytecode.LOAD_CONST, Operand1: 1},
        {Opcode: bytecode.STORE_VAR, Operand1: 1}, // var[1] = b = 1
        {Opcode: bytecode.LOAD_CONST, Operand1: n},
        {Opcode: bytecode.STORE_VAR, Operand1: 2}, // var[2] = counter = n
        
        {Opcode: bytecode.LOAD_VAR, Operand1: 0},
        {Opcode: bytecode.PRINT},
        
        // Loop: while counter > 0
        // loop_start (PC = 8):
        {Opcode: bytecode.LOAD_VAR, Operand1: 2}, // Load counter
        {Opcode: bytecode.JZ, Operand1: 20},      // If counter == 0, jump to end
        
        {Opcode: bytecode.LOAD_VAR, Operand1: 1},
        {Opcode: bytecode.PRINT},
        
        // Calculate next Fibonacci: temp = a + b
        {Opcode: bytecode.LOAD_VAR, Operand1: 0}, // Load a
        {Opcode: bytecode.LOAD_VAR, Operand1: 1}, // Load b
        {Opcode: bytecode.ADD},                   // a + b
        {Opcode: bytecode.STORE_VAR, Operand1: 3}, // var[3] = temp = a + b
        
        // Update: a = b, b = temp
        {Opcode: bytecode.LOAD_VAR, Operand1: 1}, // Load b
        {Opcode: bytecode.STORE_VAR, Operand1: 0}, // a = b
        {Opcode: bytecode.LOAD_VAR, Operand1: 3}, // Load temp
        {Opcode: bytecode.STORE_VAR, Operand1: 1}, // b = temp
        
        // Decrement counter
        {Opcode: bytecode.LOAD_VAR, Operand1: 2}, // Load counter
        {Opcode: bytecode.LOAD_CONST, Operand1: 1},
        {Opcode: bytecode.SUB},                   // counter - 1
        {Opcode: bytecode.STORE_VAR, Operand1: 2}, // Store back to counter
        
        {Opcode: bytecode.JMP, Operand1: 8}, // Jump back to loop_start
        
        // end (PC = 20):
        {Opcode: bytecode.HALT},
    }
}

func CountDown(start int32) []bytecode.Instruction {
    return []bytecode.Instruction{
        {Opcode: bytecode.LOAD_CONST, Operand1: start},
        {Opcode: bytecode.STORE_VAR, Operand1: 0}, // counter = start
        
        // loop_start (PC = 2):
        {Opcode: bytecode.LOAD_VAR, Operand1: 0}, // Load counter
        {Opcode: bytecode.JZ, Operand1: 9},       // If counter == 0, jump to end
        {Opcode: bytecode.LOAD_VAR, Operand1: 0}, // Load counter
        {Opcode: bytecode.PRINT},                 // Print counter
        {Opcode: bytecode.LOAD_VAR, Operand1: 0}, // Load counter
        {Opcode: bytecode.LOAD_CONST, Operand1: 1},
        {Opcode: bytecode.SUB},                   // counter - 1
        {Opcode: bytecode.STORE_VAR, Operand1: 0}, // Store back to counter
        {Opcode: bytecode.JMP, Operand1: 2},      // Jump back to loop_start
        
        // end (PC = 9):
        {Opcode: bytecode.HALT},
    }
}